package command

import (
    "fmt"
    "strings"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/types/value"
)

const (
    LINE = 1
    FILE = 2
)

type SetFunc func(cmd *Item, meta *Command, val types.Object) int

type Item struct {
    Command  *Command
    Type      int
    Scope     string
    Set       SetFunc
    State     bool
    Offset    uintptr
    Load      types.Object
}

type Command struct {
    flag     string
    key      string
    value    types.Object
    details  string
}

func New(flag string, key string, value types.Object, details string) *Command {
    return &Command{ flag, key, value, details }
}

func (r *Command) GetFlag() string {
    return r.flag
}

func (r *Command) GetKey() string {
    return r.key
}

func (r *Command) GetDetails() string {
    key := ""
    if v := r.details; v != key {
        return v
    }
    return key
}

func (r *Command) GetValue() types.Value {
    if v := r.value; v != nil {
        return value.New(v)
    }

    return nil
}

func (r *Command) SetValue(o types.Object) int {
    return SetObject(nil, r, o)
}

func (r *Command) Clear() {
    r.value = nil
}

var Pool []Item

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

func List(_ *Item, _ *Command, _ types.Object) int {
    for _, item := range Pool {
        if item.Type != LINE {
            continue
        }

        if Command := item.Command; Command != nil {
            fmt.Printf("%s\t%s\t\t%s\n", Command.flag, Command.key, Command.details)
        }
    }

    return state.Done
}

func Display(_ *Item, meta *Command, _ types.Object) int {
    if meta != nil {
        fmt.Println(meta.details)
    }

    return state.Done
}

func SetObject(_ *Item, c *Command, value types.Object) int {
    if c == nil || value == nil {
        return state.Error
    }

    c.value = value

    return state.Ok
}
