package adapter

import (
    "os"
    "time"
    "github.com/rookie-xy/hubble/state"
    "fmt"
    "sync"
)

type FileState struct {
    state.State

    Id         string         `json:"-"`
    Fileinfo   os.FileInfo    `json:"-"`
    Source     string         `json:"file"`
    Lno        uint64         `json:"lno"`
    Offset     int64          `json:"offset"`
    Timestamp  time.Time      `json:"timestamp"`
    TTL        time.Duration  `json:"ttl"`
    Type       string         `json:"type"`
}

func New() FileState {
    return FileState{
        State:      state.New(),
        Timestamp:  time.Now(),
        TTL:        -1,
        Type:       "file",
    }
}

func (fs *FileState) Init(fi os.FileInfo, path, Type string) error {
    fs.Fileinfo = fi
    fs.Source = path
    fs.Type = Type

    return nil
}

// ID returns a unique id for the state as a string
func (fs *FileState) ID() string {
    // Generate id on first request. This is needed as id is
    // not set when converting back from json
    if fs.Id == "" {
        fs.Id = fs.Type
    }

    return fs.Id
}

// IsEqual compares the state to an other state supporing
// stringer based on the unique string
func (fs *FileState) IsEqual(new *FileState) bool {
    return fs.ID() == new.ID()
}

// IsEmpty returns true if the state is empty
func (fs *FileState) IsEmpty() bool {
    return *fs == FileState{}
}

type FileStates struct {
    States []FileState `json:"states"`
    sync.RWMutex
}

func News() *FileStates {
    return &FileStates{
        States: []FileState{},
    }
}

func (r *FileStates) Update(state FileState) {
    r.Lock()
    defer r.Unlock()

    index, _ := r.findPrevious(state)
    if index >= 0 {
        r.States[index] = state
    } else {
        r.States = append(r.States, state)
        fmt.Println("finder", "New state added for %s", state.Source)
    }
}

func (r *FileStates) FindPrevious(newState FileState) FileState {
    r.RLock()
    defer r.RUnlock()

    _, state := r.findPrevious(newState)
    return state
}

func (s *FileStates) findPrevious(newState FileState) (int, FileState) {
    for index, oldState := range s.States {
        if oldState.IsEqual(&newState) {
            return index, oldState
        }
    }

    return -1, FileState{}
}

func (r *FileStates) Cleanup() int {
    r.Lock()
    defer r.Unlock()

    statesBefore := len(r.States)
    currentTime := time.Now()
    states := r.States[:0]

    for _, state := range r.States {
        expired := (state.TTL > 0 && currentTime.Sub(state.Timestamp) > state.TTL)

        if state.TTL == 0 || expired {
            if state.On {
                fmt.Println("state", "State removed for %v because of older: %v", state.Source, state.TTL)
                continue // drop state
            } else {
                fmt.Println("State for %s should have been dropped, but couldn't as state is not finished.", state.Source)
            }
        }

        states = append(states, state) // in-place copy old state
    }

    r.States = states

    return statesBefore - len(r.States)
}

// Count returns number of states
func (r *FileStates) Count() int {
    r.RLock()
    defer r.RUnlock()

    return len(r.States)
}

// Returns a copy of the file states
func (r *FileStates) GetStates() []FileState {
    r.RLock()
    defer r.RUnlock()

    newStates := make([]FileState, len(r.States))
    copy(newStates, r.States)

    return newStates
}

func (r *FileStates) SetStates(states []FileState) {
    r.Lock()
    defer r.Unlock()

    r.States = states
}

func (r *FileStates) Copy() *FileStates {
//    states := NewStates()
//    states.states = r.GetStates()

//    return states
    return nil
}
