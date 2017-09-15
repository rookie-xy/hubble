package adapter

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/proxy"
)

type Sincedb interface {
    proxy.Forward
    Search()
}

func AdapteeSincedb(fw proxy.Forward) Sincedb {
    return fw
}
