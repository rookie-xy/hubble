package adapter

import (
	"github.com/rookie-xy/hubble/observer"
	"github.com/rookie-xy/hubble/module"
	"github.com/rookie-xy/hubble/types"
)

type ConfigureObserver interface {
	observer.Observer
	Reload(o types.Object) error
}

func ToConfigureObserver(o observer.Observer) ConfigureObserver {
    return o.(ConfigureObserver)
}

type ModuleObserver interface {
	observer.Observer
	module.Template
}

func ToModuleObserver(o observer.Observer) ModuleObserver {
    return o.(ModuleObserver)
}
