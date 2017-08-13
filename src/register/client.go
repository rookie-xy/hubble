package register

import (
    "strings"
    "fmt"
    "github.com/rookie-xy/worker/src/client"
)

func Client(name string, f client.Factory) {
    name = name[strings.LastIndex(name, ".") + 1:]
    if name == "" {
        return
    }

    if _, exists := client.Clients[name]; exists {
        panic(fmt.Sprintf("client '%v' already registered ", name))
    }

    client.Clients[name] = f
}
