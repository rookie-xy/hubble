package register

import (
    "fmt"
    "github.com/rookie-xy/hubble/proxy"
)

func Server(name string, s proxy.Server) {
    if name == "" {
        return
    }

    if _, exists := proxy.Reverses[name]; exists {
        fmt.Printf("This server '%v' already registered\n", name)
        return
    }

    proxy.Reverses[name] = s
}
