package factory

import (
    "fmt"
    "github.com/rookie-xy/worker/src/codec"
    "github.com/rookie-xy/worker/src/observer"
//    "github.com/rookie-xy/worker/src/channel"
    "github.com/rookie-xy/worker/src/client"
    "errors"
)

// factory method
func Codec(cfg *codec.Config) (codec.Codec, error) {
    key := "json"
    if name := cfg.Name; name != "" {
        key = name
    }

    method := codec.Codecs[key]
    if method == nil {
        return nil, fmt.Errorf("'%v' output codec is not available", key)
    }

    return method(cfg)
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

func Client(name string) (client.Client, error) {
    key := ""
    if name != key {
        key = name
    }

    factory := client.Clients[key]
    if factory == nil {
        fmt.Errorf("'%v' sender is not available", key)
        return nil, errors.New("client error")
    }

    return factory(name)
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
