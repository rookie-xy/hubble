package builder

import (
    "fmt"

    "github.com/rookie-xy/hubble/module"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/factory"
)

type Director struct {
    log.Log
    build Builder
}

func New(log log.Log) *Director {
    return &Director{Log: log}
}

func (d *Director) Director(b Builder) error /**Director*/ {
    //return &Director{build: b}
    d.build = b
    return nil
}

func (d *Director) Construct(core []string) {
    scope := module.Worker
    key   := scope + "." + module.Configure

    configure := module.Setup(key, d.Log)
    if configure == nil {
        fmt.Println("Not found configure module")
        return
    }

    subject := factory.Subject(module.Configure)
    if subject != nil {
        for _, name := range core {
            key := scope + "." + name
            if module := module.Setup(key, d.Log); module != nil {
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
