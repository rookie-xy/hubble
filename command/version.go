package command

import (
    "fmt"
    "github.com/rookie-xy/hubble/types"
)

func Display(_ *Item, meta *Command, _ types.Object) error {
    if meta != nil {
        fmt.Println(meta.details)
    }

    return nil
}

