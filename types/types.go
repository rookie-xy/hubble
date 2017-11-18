package types

const (
    ARRAY = iota
    MAP
    STRING
    INTEGER
    Unknown
)

type Map   map[interface{}]interface{}
type Array []interface{}
