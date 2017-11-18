package command

import (
    "fmt"
    "github.com/rookie-xy/hubble/types"
)

func Display(_ *Item, c *Command, _ types.Object) error {
    if c != nil {
        fmt.Println(c.details)
        return nil
    }

    return fmt.Errorf("Display: the command is nil")
}

