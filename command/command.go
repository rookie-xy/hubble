package command

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/types/value"
)

type Command struct {
    flag     string
    key      string
    object   types.Object
    details  string
}

func New(flag string, key string, object types.Object, details string) *Command {
    return &Command{ flag, key, object, details }
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
    if o := r.object; o != nil {
        return value.New(o)
    }

    return nil
}

func (r *Command) GetObject() types.Object {
    return r.object
}

func (r *Command) SetValue(o types.Object) error {
    return SetObject(nil, r, o)
}

func (r *Command) Clear() {
    r.object = nil
}
