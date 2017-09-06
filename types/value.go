package types

type Value interface {
    GetString() string
    GetInt() int
    GetUint64() uint64
    GetArray() []interface{}
    GetMap() map[interface{}]interface{}
    GetType() int
    GetBool() bool
    GetIterator(Object) Iterator
}
