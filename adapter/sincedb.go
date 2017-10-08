package adapter

import (
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/types"
)

type SinceDB interface {
    proxy.Forward
    Get() []types.Value
}

func ToSinceDB(client proxy.Forward) SinceDB {
    return client.(SinceDB)
}
