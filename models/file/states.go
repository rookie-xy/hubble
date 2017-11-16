package file

import (
	"sync"
	"fmt"
    "time"
)

type States struct {
    States []State `json:"states"`
    sync.RWMutex
}

func News() *States {
    return &States{
        States: []State{},
    }
}

func (s *States) Update(new State) {
    s.Lock()
    defer s.Unlock()

    index, _ := s.findPrevious(new)
    new.Timestamp = time.Now()

    if index >= 0 {
        s.States[index] = new
    } else {
        s.States = append(s.States, new)
        fmt.Printf("New file added for %s\n", new.Source)
    }
}

func (s *States) FindPrevious(new State) State {
    s.RLock()
    defer s.RUnlock()

    _, state := s.findPrevious(new)
    return state
}

func (s *States) findPrevious(new State) (int, State) {
    for index, old := range s.States {
        if old.IsEqual(&new) {
            return index, old
        }
    }

    return -1, State{}
}

func (s *States) Cleanup() int {
    s.Lock()
    defer s.Unlock()

    statesBefore := len(s.States)
    currentTime := time.Now()
    states := s.States[:0]

    for _, state := range s.States {
        expired := (state.TTL > 0 && currentTime.Sub(state.Timestamp) > state.TTL)

        if state.TTL == 0 || expired {
            if state.Finished {
                fmt.Println("models", "State removed for %v because of older: %v", state.Source, state.TTL)
                continue // drop models
            } else {
                fmt.Println("State for %s should have been dropped, but couldn't as models is not finished.", state.Source)
            }
        }

        states = append(states, state) // in-place copy old models
    }

    s.States = states

    return statesBefore - len(s.States)
}

// Count returns number of states
func (s *States) Count() int {
    s.RLock()
    defer s.RUnlock()

    return len(s.States)
}

// Returns a copy of the file states
func (s *States) Get() []State {
    s.RLock()
    defer s.RUnlock()

    new := make([]State, len(s.States))
    copy(new, s.States)

    return new
}

func (s *States) Set(states []State) {
    s.Lock()
    defer s.Unlock()

    s.States = states
}

func (s *States) Copy() *States {
    states := News()
    states.States = s.Get()

    return states
}
