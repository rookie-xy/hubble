package register

import (
    "fmt"
    "github.com/rookie-xy/hubble/proxy"
)

func Server(name string, f proxy.Server) {
    if name == "" {
        return
    }

    if _, exists := proxy.Reverses[name]; exists {
        panic(fmt.Sprintf("client '%v' already registered ", name))
    }

    proxy.Reverses[name] = f
}
