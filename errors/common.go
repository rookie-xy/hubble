package errors

import "errors"

var (
    ErrType      = errors.New("type is not equal")
    ErrConfigure = errors.New("not found agents configure")
)
