package adapter

import (
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/output"
    "fmt"
)

type BatchForward interface {
    proxy.Forward
    Senders(events []event.Event) error
}

func ToBatchForward(f proxy.Forward) BatchForward {
    return f.(BatchForward)
}

func ToOutput(f proxy.Forward) (proxy.Forward, error) {
	fmt.Println("wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww")
    return f.(output.Output).New()
}
