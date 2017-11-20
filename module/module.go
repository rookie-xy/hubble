package module

import (
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/factory"
    "github.com/rookie-xy/hubble/memento"
    "fmt"
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
	if length := len(r.modules); length < 0 {
	    fmt.Println("Not found module")
	    return
    }

    for _, module := range r.modules {
        if module != nil {
            go module.Main()
        }
    }

    for {
        select {
        // 只跟操作系统打交到

        }
    }
}

func (r *module) Exit(code int) {
    if n := len(r.modules); n > 0 {
        for _, module := range r.modules {
            module.Exit(code)
        }
    }

    r.configure.Exit(0)
}

func (r *module) Load(module Template) {
    if module != nil {
        r.modules = append(r.modules, module)
    }
}

func (r *module) Configure(cfg Template) error {
    if cfg != nil {
        r.configure = cfg

    } else {
        return fmt.Errorf("cfg is nil")
    }

    r.configure.Init()

    if subject := factory.Subject(memento.Name); subject != nil {
        if obs := factory.Observer(Configure); obs != nil {
            subject.Attach(obs)
        } else {
            return  fmt.Errorf("Not found configure observer")
        }
    } else {
        return fmt.Errorf("Not found configure subject")
    }

    go r.configure.Main()

    return nil
}