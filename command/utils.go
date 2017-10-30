package command

import (
    "strings"

    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/plugin"
    "fmt"
)

func Setup(flag, value string) int {
    for _, item := range Pool {

        if item.Type != LINE || item.Command.flag != flag {
            continue
        }

        return item.Set(&item, item.Command, value)
    }

    return state.Error
}

func File(nameSpace, key string, value types.Object) int {
    for _, item := range Pool {

        if item.Scope != nameSpace || item.Type != FILE {
            continue
        }

        if item.Command.key != key {
            prefix := item.Command.key
            if n := strings.Index(prefix, "."); n > -1 {
                prefix = prefix[0:n]
            }

            if item.Command.flag == plugin.Flag {
                if strings.HasPrefix(key, prefix) {
                    fmt.Println("pluginnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn: ", key)
                    item.Command.key = key
                } else {
                    continue
                }

            } else {
                continue
            }
        }

        return item.Set(&item, item.Command, value)
    }

    return state.Error
}

func SetObject(_ *Item, c *Command, value types.Object) int {
    if c == nil || value == nil {
        return state.Error
    }

    c.value = value

    return state.Ok
}
