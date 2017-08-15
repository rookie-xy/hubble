package factory

import (
    "fmt"

    "github.com/rookie-xy/hubble/src/codec"
    "github.com/rookie-xy/hubble/src/observer"
    "github.com/rookie-xy/hubble/src/client"
    "github.com/rookie-xy/hubble/src/command"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/pipeline"
)

// factory method
func Codec(name string, l log.Log, c *command.Command) (codec.Codec, error) {
    key := "json"
    if name != "" {
        key = name
    }

    factory := codec.Codecs[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' codec is not available", key)
    }

    return factory(l, c)
}

func Pipeline(name string, l log.Log, c *command.Command) (pipeline.Pipeline, error) {
    key := "slot"
    if name != "" {
        key = name
    }

    factory := pipeline.Pipelines[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' pipeline is not available", key)
    }

    return factory(l, c)
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

func Client(name string, l log.Log, c *command.Command) (client.Client, error) {
    key := ""
    if name != key {
        key = name
    }

    factory := client.Clients[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' client is not available", key)
    }

    return factory(l, c)
}


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
