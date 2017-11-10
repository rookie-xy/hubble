package types

import "time"

type Value interface {
    GetString() string
    GetBytes() []byte
    GetInt() (int, error)
    GetUint64() (uint64, error)
    GetBool() bool
    GetDuration() (time.Duration, error)
    GetArray() []interface{}
    GetMap() map[interface{}]interface{}
    GetIterator(Object) Iterator
    GetType() int
}
