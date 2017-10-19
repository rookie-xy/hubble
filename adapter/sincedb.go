package adapter

import (
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/event"
)

type SinceDB interface {
	BatchForward
    Get() []types.Value
}

func FileSinceDB(client proxy.Forward) SinceDB {
    return client.(SinceDB)
}
