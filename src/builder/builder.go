package builder

import (
    "fmt"

    "github.com/rookie-xy/worker/src/module"
    "github.com/rookie-xy/worker/src/log"
    "github.com/rookie-xy/worker/src/factory"
)

type Builder interface {
    Configure(m module.Template) int
    module.Module
}

type Director struct {
    log.Log
    build Builder
}

func Directors(b Builder) *Director {
    return &Director{build: b}
}

func (r *Director) Construct(core []string) {
    scope := module.Worker
    key   := scope + "." + module.Configure

    configure := module.Setup(key, r.Log)
    if configure == nil {
        fmt.Println("Not found configure module")
        return
    }

    subject := factory.Subject(module.Configure)
    if subject != nil {
        for _, name := range core {
            key := scope + "." + name
            if module := module.Setup(key, r.Log); module != nil {
                fmt.Println("rrrrrrrrrrrrrrrrr", key)
                r.build.Load(module)
            }

            if f := factory.Observer(name); f != nil {
                fmt.Println(name)
                subject.Attach(f)
            }
        }

    } else {
        fmt.Println("Not found configure subject")
        return
    }

    r.build.Configure(configure)
}
