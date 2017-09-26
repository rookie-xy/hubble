package types

import "time"

type Value interface {
    GetString() string
    GetBytes() []byte
    GetInt() int
    GetUint64() uint64
    GetBool() bool
    GetDuration() time.Duration
    GetArray() []interface{}
    GetMap() map[interface{}]interface{}
    GetIterator(Object) Iterator
    GetType() int
}
