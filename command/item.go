package command

import (
    "github.com/rookie-xy/hubble/types"
)

type Set func(cmd *Item, meta *Command, val types.Object) error

type Item struct {
    Command  *Command
    Type      int
    Scope     string
    Name      string
    Set       Set
    Load      types.Object
}

var Pool []Item
