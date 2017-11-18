package command

import (
    "strings"

    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/plugin"
    "fmt"
)

func Setup(flag, value string) error {
    if flag == "" || value == "" {
        return fmt.Errorf("[command] flag or value is nil")
    }

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

func SetObject(_ *Item, c *Command, o types.Object) error {
    if c == nil || o == nil {
        return fmt.Errorf("command or object is nil")
    }

    c.object = o
    return nil
}
