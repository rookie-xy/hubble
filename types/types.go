package types

const (
    ARRAY = iota
    MAP
    STRING
    INTEGER
)

type Map   map[interface{}]interface{}
type Array []interface{}
