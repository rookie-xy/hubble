package command

import (
    "fmt"
    "github.com/rookie-xy/hubble/types"
    "errors"
)

func Version(_ *Item, c *Command, _ types.Object) error {
    if c != nil {
        fmt.Println(c.details)
        return errors.New("version end")
    }

    return fmt.Errorf("Display: the command is nil")
}

