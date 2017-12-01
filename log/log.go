package log

import (
    "log"

  . "github.com/rookie-xy/hubble/log/level"
	"os"
)

type Factory func(l Level, f string, args ...interface{})

type Log interface {
    Output(depth int, s string) error
}

type Logger struct {
   *log.Logger
	level  Level
}

func New() *Logger {
	this := &Logger{
		Logger: log.New(
	        os.Stderr,
	        "",
            log.LstdFlags | log.Lmicroseconds,
        ),
        level: INFO,
    }

    return this
}

func (l *Logger) Set(ll Level) {
	l.level = ll
}

func (l *Logger) Get() Level {
    return l.level
}

func (l *Logger) Copy(logger *Logger) {
    l.Logger = logger.Logger
    l.level = logger.level
}

func (l *Logger) Print(ll Level, f string, args ...interface{}) {
    Print(l.Logger, l.level, ll, f, args...)
}
