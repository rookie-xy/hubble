package register

import (
//    "strings"
    "fmt"
    "github.com/rookie-xy/hubble/proxy"
)

func Client(name string, c proxy.Client) {
//    name = name[strings.LastIndex(name, ".") + 1:]
    if name == "" {
        return
    }

    if _, exists := proxy.Clients[name]; exists {
        fmt.Printf("This client '%v' already registered\n", name)
        return
    }

    proxy.Clients[name] = c
}

func Forword(name string, f proxy.Forward) {
    if name == "" {
        return
    }

    if _, exists := proxy.Forwards[name]; exists {
        fmt.Printf("This forword object '%v' already registered\n", name)
        return
    }

    proxy.Forwards[name] = f
}
