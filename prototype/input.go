package prototype

import "github.com/rookie-xy/hubble/input"

type InputPrototype interface {
	Prototype
	input.Input
}

func Input(this input.Input) input.Input {
    prototype := this.(InputPrototype)
    return prototype.Clone().(input.Input)
}
