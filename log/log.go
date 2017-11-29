package log

import (
    "os"
    "log"

	"github.com/rookie-xy/hubble/types"
  . "github.com/rookie-xy/hubble/log/level"
)

type Factory func(l Level, f string, args ...interface{})

type Log interface {
    Output(depth int, s string) error
}

type logs struct {
   *log.Logger

	Log      Log
	level    Level
}

func New() *logs {
	prefix := prefix.GetValue()
	verbose := verbose.GetValue()
	level := level.GetValue()

	this := &logs{
		Log: log.New(
	        os.Stderr,
	        prefix.GetString(),
            log.LstdFlags | log.Lmicroseconds,
        ),
        level: INFO,
    }

    var err error
    this.level, err = Parse(level.GetString(), verbose.GetBool())
    if err != nil {
        return nil
	}

    return this
}

func (l *logs) Init(prefix, verbose, level  types.Value) error {
	//prefix.GetString()
	return nil
}

func (l *logs) Level() Level {
    return l.level
}
