package command

import (
    "fmt"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/state"
)

func Display(_ *Item, meta *Command, _ types.Object) int {
    if meta != nil {
        fmt.Println(meta.details)
    }

    return state.Done
}

