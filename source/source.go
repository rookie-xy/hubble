package source

import (
    "io"
    "errors"
    "bufio"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
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

type Factory func(log.Log, types.Value, Source) (Source, error)

type Source interface {
    io.ReadCloser
}

var Sources = map[string]Factory{}
