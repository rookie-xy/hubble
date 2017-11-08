package file

import (
	"os"
	"time"
)

type State struct {
    Id         string         `json:"-"`
    Finished   bool           `json:"-"`
    Fileinfo   os.FileInfo    `json:"-"`
    Source     string         `json:"file"`
    Lno        uint64         `json:"lno"`
    Offset     int64          `json:"offset"`
    Timestamp  time.Time      `json:"timestamp"`
    TTL        time.Duration  `json:"ttl"`
    Type       string         `json:"type"`
    //File       id.ID
}

func New() State {
    return State{
        Timestamp:  time.Now(),
        TTL:        -1,
        Type:       "file",
    }
}

func (s *State) Init(id string, fi os.FileInfo, path, Type string) error {
	s.Id = id
    s.Fileinfo = fi
    s.Source = path
    s.Type = Type

    return nil
}

// ID returns a unique id for the models as a string
func (fs *State) ID() string {
    // Generate id on first request. This is needed as id is
    // not set when converting back from json
    if fs.Id == "" {
        fs.Id = fs.Type
    }

    return fs.Id
}

// IsEqual compares the models to an other models supporing
// stringer based on the unique string
func (fs *State) IsEqual(new *State) bool {
    return fs.ID() == new.ID()
}

// IsEmpty returns true if the models is empty
func (fs *State) IsEmpty() bool {
    return *fs == State{}
}
