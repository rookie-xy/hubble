package factory

import (
    "fmt"

    "github.com/rookie-xy/hubble/codec"
    "github.com/rookie-xy/hubble/observer"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/pipeline"
    "github.com/rookie-xy/hubble/types"
)

// factory method
func Codec(name string, l log.Log, v types.Value) (codec.Codec, error) {
    key := "json"
    if name != "" {
        key = name
    }

    factory := codec.Codecs[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' codec is not available", key)
    }

    return factory(l, v)
}

func Pipeline(name string, l log.Log, v types.Value) (pipeline.Queue, error) {
    key := "channel"
    if name != "" {
        key = name
    }

    factory := pipeline.Queues[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' pipeline is not available", key)
    }

    return factory(l, v)
}

func Subject(name string) observer.Subject {
    key := ""
    if name != key {
        key = name
    }

    subject := observer.Subjects[key]
    if subject == nil {
        fmt.Println("Not found subject:", key)
        return nil
    }

    return subject
}

func Observer(name string) observer.Observer {
    key := ""
    if name != key {
        key = name
    }

    observer := observer.Observers[key]
    if observer == nil {
        fmt.Errorf("'%v' observer is not available", key)
        return nil
    }

    return observer
}

func Clone(name string) pipeline.Queue {
    key := ""
    if name != key {
        key = name
    }

    clone := pipeline.Clones[key]
    if clone == nil {
        fmt.Errorf("'%v' clone is not available", key)
        return nil
    }

    return clone
}

func Client(name string, l log.Log, v types.Value) (proxy.Forward, error) {
    key := ""
    if name != key {
        key = name
    }

    factory := proxy.Forwards[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' client is not available", key)
    }

    return factory(l, v)
}

func Server(name string, l log.Log, v types.Value) (proxy.Reverse, error) {
    key := ""
    if name != key {
        key = name
    }

    factory := proxy.Reverses[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' server is not available", key)
    }

    return factory(l, v)
}

/*
func Push(name string) channel.Push {
    key := ""
    if name != key {
        key = name
    }

    push := channel.Channels[key]
    if push == nil {
        fmt.Errorf("'%v' channel push is not available", key)
        return nil
    }

    return push.Clone()
}

func Pull(name string) channel.Pull {
    key := ""
    if name != key {
        key = name
    }

    pull := channel.Channels[key]
    if pull == nil {
        fmt.Errorf("'%v' channel pull is not available", key)
        return nil
    }

    return pull
}
*/


/*
type Creater interface {
	CreateProduct(action string) Producter
	registerProduct(Producter)
}

type Producter interface {
	Use() string
}

type ConcreteCreator struct {
	products []*Producter
}

func (self *ConcreteCreator) CreateProduct(action string) Producter {
	var product Producter

	switch action {
	case "A":
		product = &ConcreteProductA{action}
	case "B":
		product = &ConcreteProductB{action}
	case "C":
		product = &ConcreteProductC{action}
	default:
		log.Fatalln("Unknown Action")
	}

	self.registerProduct(product)

	return product
}

func (self *ConcreteCreator) registerProduct(product Producter) {
	self.products = append(self.products, &product)
}

type ConcreteProductA struct {
	action string
}

func (self *ConcreteProductA) Use() string {
	return self.action
}

type ConcreteProductB struct {
	action string
}

func (self *ConcreteProductB) Use() string {
	return self.action
}

type ConcreteProductC struct {
	action string
}

func (self *ConcreteProductC) Use() string {
	return self.action
}
*/
