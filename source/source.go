package source

import (
    "os"
    "io"
)

type Source interface {
    io.ReadCloser

    Name() string
    Stat() (os.FileInfo, error)
    Continuable() bool // can we continue processing after EOF?
}
