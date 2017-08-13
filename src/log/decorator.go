package log

import "github.com/rookie-xy/worker/src/prototype"

type Log interface {
    Print(a prototype.Object)
}

type log struct {
    name string
}

func New() *log {
    return &log{}
}

func (r *log) Print(a prototype.Object) {
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
