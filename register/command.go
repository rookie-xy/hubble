package register

import "github.com/rookie-xy/hubble/command"

func Commands(_ string, items []command.Item) {
    for _, item := range items {
        command.Pool = append(command.Pool, item)
    }
}
