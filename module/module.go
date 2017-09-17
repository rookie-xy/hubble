package module

import (
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/factory"
    "github.com/rookie-xy/hubble/memento"
)

//type NewFunc func(log log.Log) Template

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