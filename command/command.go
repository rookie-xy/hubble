package command

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/types/value"
)

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
