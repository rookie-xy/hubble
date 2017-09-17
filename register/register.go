package register

import (
    "sync"
    "github.com/rookie-xy/hubble/module"
    "github.com/rookie-xy/hubble/command"
)

// singleton
type singleton struct {
}

var instance *singleton
var once sync.Once

func getInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })

    return instance
}

// flyweight
func Module(scope, name string, items []command.Item, f module.Factory) {
    merge := getInstance()

    key := ""
    if scope != key && name != key {
        key = scope + "." + name

    } else {
        return
    }

    if l := len(items); l <= 0 {
        return
    } else {
        merge.Command(key, items)
    }

    if f != nil {
        merge.Module(key, f)
    }
}

func (r *singleton) Command(key string, value []command.Item) {
    for _, e := range value {
        command.Pool = append(command.Pool, e)
    }
}

func (r *singleton) Module(key string, f module.Factory) {

    if _, exist := module.Pool[key]; !exist {
        //fmt.Println("register: ", key)
        module.Pool[key] = &f
    }
}
