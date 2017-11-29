package register

import (
    "github.com/rookie-xy/hubble/module"
    "github.com/rookie-xy/hubble/command"
    "fmt"
)

// flyweight
func Module(scope, name string, items []command.Item, f module.Factory) {
    key := ""
    if scope != key && name != key {
        key = scope + "." + name

    } else {
        return
    }

    if length := len(items); length <= 0 {
    	fmt.Printf("command items length is %d\n", length)
        return
    } else {
        Commands(key, items)
    }

    if f != nil {
        if err := Modules(key, f); err != nil {
            fmt.Println(err)
        }
    }
}
