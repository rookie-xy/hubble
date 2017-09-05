package command

import (
    "fmt"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/state"
)

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

