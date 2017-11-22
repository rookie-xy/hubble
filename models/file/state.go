package file

import (
	"os"
	"time"
)

type State struct {
    fid        string         `json:"-"`
    Finished   bool           `json:"-"`
    Fileinfo   os.FileInfo    `json:"-"`
    Source     string         `json:"source"`
    Lno        uint64         `json:"lno"`
    Offset     int64          `json:"offset"`
    Timestamp  time.Time      `json:"timestamp"`
    TTL        time.Duration  `json:"ttl"`
    Type       string         `json:"type"`
    ID         ID
}

func New() State {
    return State{
        TTL:        -1,
        Type:       "file",
        Finished:   false,
        Timestamp:  time.Now(),
    }
}

func (s *State) Init(path string, fi os.FileInfo) error {
    s.Source = path
    s.Fileinfo = fi
    s.ID = Id(fi)
    return nil
}

// ID returns a unique id for the models as a string
func (s *State) Fid() string {
    // Generate id on first request. This is needed as id is
    // not set when converting back from json
    if s.fid == "" {
        s.fid = s.ID.String()
    }

    return s.fid
}

// IsEqual compares the models to an other models supporing
// stringer based on the unique string
func (s *State) IsEqual(new *State) bool {
    return s.Fid() == new.Fid()
}

// IsEmpty returns true if the models is empty
func (s *State) IsEmpty() bool {
    return *s == State{}
}
