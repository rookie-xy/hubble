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
        panic(fmt.Sprintf("client '%v' already registered ", name))
    }

    proxy.Clients[name] = c
}

func Forword(name string, F proxy.Forward) {
    if name == "" {
        return
    }

    if _, exists := proxy.Forwards[name]; exists {
        panic(fmt.Sprintf("this pipeline clone name '%v' already registered ", name))
    }

    proxy.Forwards[name] = F
}
