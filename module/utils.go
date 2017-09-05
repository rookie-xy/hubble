package module

import (
    "fmt"

    "github.com/rookie-xy/hubble/log"
)

func Setup(key string, log log.Log) Template {
    if key == "" {
        goto J_RET
    }

    if this, exist := Pool[key]; exist {
        if new := *this; new != nil {
            return new(log)
        } else {
            fmt.Println("New func is nil")
        }

    } else {
        fmt.Println("Not found key: ", key)
    }

J_RET:
    return nil
}
