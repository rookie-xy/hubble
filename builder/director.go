package builder

import (
    "errors"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/module"
    "github.com/rookie-xy/hubble/factory"
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

    return errors.New("director failure")
}

func (d *Director) Construct(core []string) error {
    scope := module.Worker
    key   := scope + "." + module.Configure

    configure := module.Setup(key, d.Logger)
    if configure == nil {
        return errors.New("not found configure module")
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
        return errors.New("not found configure subject")
    }

    d.build.Configure(configure)
    return nil
}
