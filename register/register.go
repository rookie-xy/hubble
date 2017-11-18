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
        commands(key, items)
    }

    if f != nil {
        if err := modules(key, f); err != nil {
            fmt.Println(err)
        }
    }
}

func commands(_ string, items []command.Item) {
    for _, item := range items {
        command.Pool = append(command.Pool, item)
    }
}

func modules(key string, f module.Factory) error {
    if _, exist := module.Pool[key]; !exist {
        module.Pool[key] = &f
        return nil
    }

    return fmt.Errorf("The %s is exist", key)
}
