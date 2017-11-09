package command

import (
    "strings"

    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/plugin"
    "fmt"
)

func Setup(flag, value string) error {
    for _, item := range Pool {

        if item.Type != LINE || item.Command.flag != flag {
            continue
        }

        return item.Set(&item, item.Command, value)
    }

    return nil
}

func File(scope, name, key string, value types.Object) error {
    for _, item := range Pool {

        if item.Scope != scope || item.Name != name || item.Type != FILE {
            continue
        }

        if item.Command.key != key {
            prefix := item.Command.key
            if n := strings.Index(prefix, "."); n > -1 {
                prefix = prefix[0:n]
            }

            if item.Command.flag == plugin.Flag {
                if strings.HasPrefix(key, prefix) {
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

    return nil
}

func SetObject(_ *Item, c *Command, value types.Object) error {
    if c == nil || value == nil {
        return fmt.Errorf("command or value is nil")
    }

    c.value = value

    return nil
}
