package types

import "github.com/rookie-xy/hubble/src/prototype"

type Value interface {
    GetString() string
    GetInt() int
    GetArray() []interface{}
    GetMap() map[interface{}]interface{}
    GetType() int
    GetIterator(prototype.Object) Iterator
}

