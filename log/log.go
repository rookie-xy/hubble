package log

type Factory func(l Level, f string, args ...interface{})

type Log interface {
    Output(maxDepth int, s string) error
}

type log struct {
 	Level    string  `flag:"log-level"`
	Verbose  bool    `flag:"verbose"`   // for backwards compatibility
	Log      Log

	// private, not really an option
	level    Level
}

func New() *log {
    return &log{}
}
/*
type Nil struct{}

func (l Nil) Output(maxDepth int, s string) error {
    return nil
}
*/
