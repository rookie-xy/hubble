package register

import (
    "fmt"
    "github.com/rookie-xy/worker/src/channel"
    "strings"
)

func Channel(name string, f channel.Factory) {
    name = name[strings.LastIndex(name, "."):]
    if name == "" {
        return
    }

    if _, exists := channel.Channels[name]; exists {
        panic(fmt.Sprintf("this channel '%v' already registered ", name))
    }

    channel.Channels[name] = f
}
