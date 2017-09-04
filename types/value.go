package types

type Value interface {
    GetString() string
    GetInt() int
    GetArray() []interface{}
    GetMap() map[interface{}]interface{}
    GetType() int
    GetBool() bool
    GetIterator(Object) Iterator
}
