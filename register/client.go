package register

import (
//    "strings"
    "fmt"
    "github.com/rookie-xy/hubble/src/proxy"
)

func Client(name string, f proxy.Client) {
//    name = name[strings.LastIndex(name, ".") + 1:]
    if name == "" {
        return
    }

    if _, exists := proxy.Forwards[name]; exists {
        panic(fmt.Sprintf("client '%v' already registered ", name))
    }

    proxy.Forwards[name] = f
}
