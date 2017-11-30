package builder

import (
    "fmt"

    "github.com/rookie-xy/hubble/module"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/factory"
    "errors"
)

type Director struct {
   *log.Logger
    build Builder
}

func New(l *log.Logger) *Director {
    return &Director{Logger: l}
}

func (d *Director) Director(b Builder) error /**Director*/ {
    if b != nil {
        d.build = b
        return nil
    }

    return errors.New("Director failure")
}

func (d *Director) Construct(core []string) {
    scope := module.Worker
    key   := scope + "." + module.Configure

    configure := module.Setup(key, d.Logger)
    if configure == nil {
        fmt.Println("Not found configure module")
        return
    }

    subject := factory.Subject(module.Configure)
    if subject != nil {
        for _, name := range core {
            key := scope + "." + name
            if module := module.Setup(key, d.Logger); module != nil {
                d.build.Load(module)
            }

            if f := factory.Observer(name); f != nil {
                subject.Attach(f)
            }
        }

    } else {
        fmt.Println("Not found configure subject")
        return
    }

    d.build.Configure(configure)
}

