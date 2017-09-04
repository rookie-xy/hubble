package module

import (
    "github.com/rookie-xy/hubble/src/log"
    "fmt"
    "github.com/rookie-xy/hubble/src/state"
    "github.com/rookie-xy/hubble/src/factory"
    "github.com/rookie-xy/hubble/src/memento"
)

const (
    Flag = "module"
    Worker = "hubble"
    Configure = "configure"
    Plugins = "plugins"
    Agents = "agents"
    Proxy = "proxy"
)

// composite
type Module interface {
    Load(module Template)
    Template
}


type NewFunc  func(log log.Log) Template
type Load func(module Template)

// template
type Template interface {
    Init()
    Main()
    Exit(code int)
}

var Pool = map[string]*NewFunc{}

// facade
type module struct {
    log.Log
    configure Template
    modules   []Template
}

func New(log log.Log) *module {
    return &module{
        Log: log,
    }
}

func (r *module) Init() {
    for _, module := range r.modules {
        if module != nil {
            module.Init()
        }
    }
}

func (r *module) Main() {
    for _, module := range r.modules {
        if module != nil {
            go module.Main()
        }
    }

    for {
        select {

        }
    }
}

func (r *module) Exit(code int) {
        // 重新加载
    /*
    select {

    case <- RELOAD:

    case <- RECONFIGURE:

    case <- EXIT:
        for _, module := range r.children {
            module.Exit()
        }
    }
    */


    for _, module := range r.modules {
        module.Exit(code)
    }
}

func (r *module) Load(module Template) {
    if module != nil {
        r.modules = append(r.modules, module)
    }
}

func (r *module) Configure(cfg Template) int {
    if cfg != nil {
        r.configure = cfg

    } else {
        return state.Error
    }

    r.configure.Init()

    if subject := factory.Subject(memento.Name); subject != nil {
	       if obs := factory.Observer(Configure); obs != nil {
            subject.Attach(obs)
        }
    }

    go r.configure.Main()

    return state.Ok
}

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
