package module

import (
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/factory"
    "github.com/rookie-xy/hubble/memento"
    "fmt"
)

// facade
type module struct {
    log.Log

    configure  Template
    modules    []Template

    mains      []func()
    exits      []func(int)

    done       chan struct{}
}

func New(log log.Log) *module {
    return &module{
        Log: log,
    }
}

func (r *module) Init() {
    if length := len(r.modules); length > 0 {
        for i, module := range r.modules {
            if module != nil {
                module.Init()

                r.mains = append(r.mains, module.Main)
                r.exits = append(r.exits, module.Exit)
            } else {
                fmt.Println("module is nil")
                if i > 0 {
                    r.Exit(0)
                }
                return
            }
        }
    }

    fmt.Println("Not found module")
}

func (r *module) Main() {
    if length := len(r.mains); length > 0 {
        for _, main := range r.mains {
            go main()
        }

    } else {
	    fmt.Println("Not found module")
	    return
    }

    for {
        select {
        // 只跟操作系统打交到
        case <- r.done:
            return
        }
    }
}

func (r *module) Exit(code int) {
    if n := len(r.exits); n > 0 {
        for _, exit := range r.exits {
            exit(code)
        }
    }

    r.configure.Exit(0)
    close(r.done)
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