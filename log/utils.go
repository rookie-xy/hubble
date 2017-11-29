package log

import (
	"strings"
	"fmt"
  . "github.com/rookie-xy/hubble/log/level"
)

func Parse(level string, verbose bool) (Level, error) {
	this := INFO

	switch strings.ToLower(level) {
	case "debug":
        this = DEBUG
	case "info":
        this = INFO
	case "warn":
        this = WARN
	case "error":
        this = ERROR
	case "fatal":
        this = FATAL
	default:
        return this, fmt.Errorf("invalid log level '%s'", level)
	}

	if verbose {
		this = DEBUG
	}

	return this, nil
}

func Print(log Log, this Level, level Level, f string, args ...interface{}) {
    if this > level {
        return
    }

	log.Output(3, fmt.Sprintf(level.String()+": "+f, args...))
}
