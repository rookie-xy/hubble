package command

import (
    "fmt"
    "github.com/rookie-xy/hubble/types"
    "errors"
)

func Help(_ *Item, _ *Command, _ types.Object) error {
	if length := len(Pool); length < 0 {
	    return fmt.Errorf("The pool length is %d \n", length)
    }

    for _, item := range Pool {
        if item.Type != LINE {
            continue
        }

        if Command := item.Command; Command != nil {
            fmt.Printf("%s\t%s\t\t%s\n", Command.flag, Command.key, Command.details)
        }
    }

    return errors.New("help end")
}
