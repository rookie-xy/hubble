package adapter

import (
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/types"
)

type SinceDB interface {
    proxy.Forward
    Add()  int
    Find() types.Object
}

func AdapterSinceDB(client proxy.Forward) SinceDB {
    return client.(SinceDB)
}
