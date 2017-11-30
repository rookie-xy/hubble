package module

import (
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/factory"
    "github.com/rookie-xy/hubble/memento"
    "fmt"
    "github.com/rookie-xy/hubble/log/level"
    "os"
)

// facade
type module struct {
   *log.Logger

    configure  Template
    modules    []Template

    mains      []func()
    exits      []func(int)

    done       chan struct{}
}

func New(l *log.Logger) *module {
    return &module{
        Logger: l,
    }
}

func (r *module) Init() {
    if n := len(r.modules); n > 0 {
        for _, module := range r.modules {
            if module != nil {
                module.Init()

                r.mains = append(r.mains, module.Main)
                r.exits = append(r.exits, module.Exit)

                continue
            }

            r.Print(level.WARN, "module is nil")
        }

    } else {
        r.Print(level.WARN, "Not found module")
    }
}

func (r *module) Main() {
    if n := len(r.mains); n > 0 {
        for _, main := range r.mains {
            go main()
        }

    } else {
	    r.Print(level.WARN,"Not found module")
	    os.Exit(1)
    }

    for {
        select {
        // just hand in with the operating system
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