package log

import (
	"strings"
	"fmt"
)

func Parse(level string, verbose bool) (Level, error) {
	lvl := INFO

	switch strings.ToLower(level) {
	case "debug":
		lvl = DEBUG
	case "info":
		lvl = INFO
	case "warn":
		lvl = WARN
	case "error":
		lvl = ERROR
	case "fatal":
		lvl = FATAL
	default:
		return lvl, fmt.Errorf("invalid log-level '%s'", level)
	}
	if verbose {
		lvl = DEBUG
	}
	return lvl, nil
}

func Logf(logger Logger, cfg Level, msg Level, f string, args ...interface{}) {
	if cfg > msg {
		return
	}
	logger.Output(3, fmt.Sprintf(msg.String()+": "+f, args...))
}
