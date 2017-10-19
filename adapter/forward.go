package adapter

import "github.com/rookie-xy/hubble/proxy"

type BatchForward interface {
	proxy.Forward
}

func ToBatchForward(f proxy.Forward) BatchForward {
    return f.(BatchForward)
}
