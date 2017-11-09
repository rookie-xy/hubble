package adapter

import (
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/models/file"
)

type SinceDB interface {
    BatchForward
    Load() []file.State
}

func FileSinceDB(client proxy.Forward) SinceDB {
    return client.(SinceDB)
}
