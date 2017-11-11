package source

import (
	"errors"
	"bufio"
)

var (
    ErrFileTruncate    = errors.New("detected file being truncated")
    ErrRenamed         = errors.New("file was renamed")
    ErrRemoved         = errors.New("file was removed")
    ErrInactive        = errors.New("file inactive")
    ErrClosed          = errors.New("reader closed")

    ErrTooLong         = bufio.ErrTooLong
    ErrFinalToken      = bufio.ErrFinalToken
    ErrAdvanceTooFar   = bufio.ErrAdvanceTooFar
    ErrNegativeAdvance = bufio.ErrNegativeAdvance
)
