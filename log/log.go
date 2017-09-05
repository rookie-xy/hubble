package log

import "github.com/rookie-xy/hubble/types"

//Decorator
type Log interface {
    Print(a types.Object)
}

type log struct {
    name string
}

func New() *log {
    return &log{}
}

func (r *log) Print(a types.Object) {
    return
}
/*
type Component interface {
	Operation() string
}

type ConcreteComponent struct {
}

func (self *ConcreteComponent) Operation() string {
	return "I am component!"
}

type ConcreteDecorator struct {
	component Component
}

func (self *ConcreteDecorator) Operation() string {
	return "<strong>" + self.component.Operation() + "</strong>"
}
*/
