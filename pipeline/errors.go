package pipeline

import (
	"errors"
)

var (
    ErrEmpty  = errors.New("pipeline was empty")
    ErrFull   = errors.New("pipeline was full")
    ErrClosed = errors.New("pipeline closed")
)
