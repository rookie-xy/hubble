package log

import (
	"log"
	"os"
	"testing"
)

type options struct {
	Level    string  `flag:"log-level"`
	Verbose  bool    `flag:"verbose"`   // for backwards compatibility
	Logger   Logger

	// private, not really an option
	level Level
}

func newOptions() *options {
	return &options{
		Level: "info",
	}
}

type app struct {
	opts *options
}

func (n *app) logf(level Level, f string, args ...interface{}) {
	Logf(n.opts.Logger, n.opts.level, level, f, args)
}

func newApp(opts *options) *app {
	if opts.Logger == nil {
		opts.Logger = log.New(os.Stderr, "[app] ", log.Ldate|log.Ltime|log.Lmicroseconds)
	}
	n := &app{
		opts: opts,
	}

	var err error
	opts.level, err = Parse(opts.Level, opts.Verbose)
	if err != nil {
		n.logf(FATAL, "%s", err)
		os.Exit(1)
	}

	n.logf(INFO, "app 0.1")
	return n
}

type mockLogger struct {
	Count int
}

func (l *mockLogger) Output(maxdepth int, s string) error {
	l.Count++
	return nil
}

func TestLogging(t *testing.T) {
	logger := &mockLogger{}
	opts := newOptions()
	opts.Logger = logger

	// Test only fatal get through
	opts.Level = "FaTaL"
	hubble := newApp(opts)
	logger.Count = 0
	for i := 1; i <= 5; i++ {
		hubble.logf(Level(i), "Test")
	}
	//test.Equal(t, 1, logger.Count)

	// Test only warnings or higher get through
	opts.Level = "WARN"
	hubble = newApp(opts)
	logger.Count = 0
	for i := 1; i <= 5; i++ {
		hubble.logf(Level(i), "Test")
	}
	//test.Equal(t, 3, logger.Count)

	// Test everything gets through
	opts.Level = "debuG"
	hubble = newApp(opts)
	logger.Count = 0
	for i := 1; i <= 5; i++ {
		hubble.logf(Level(i), "Test")
	}
	//test.Equal(t, 5, logger.Count)

	// Test everything gets through with verbose = true
	opts.Level = "fatal"
	opts.Verbose = true
	hubble = newApp(opts)
	logger.Count = 0
	for i := 1; i <= 5; i++ {
		hubble.logf(Level(i), "Test")
	}
	//test.Equal(t, 5, logger.Count)
}

func TestNoLogger(t *testing.T) {
	opts := newOptions()
	opts.Logger = NilLogger{}
	app := newApp(opts)

	app.logf(ERROR, "should never be logged")
}
