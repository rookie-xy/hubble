package prototype

import (
	"github.com/rookie-xy/hubble/proxy"
	"github.com/rookie-xy/hubble/output"
)

type ForwardPrototype interface {
	Prototype
	proxy.Forward
}

func Forward(this proxy.Forward) proxy.Forward {
    prototype := this.(ForwardPrototype)
    return prototype.Clone().(proxy.Forward)
}

type OutputPrototype interface {
	Prototype
	output.Output
}

func Output(this output.Output) output.Output {
    prototype := this.(OutputPrototype)
    return prototype.Clone().(output.Output)
}
