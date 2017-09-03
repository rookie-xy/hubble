package types

// Iterator
type Iterator interface {
    Index() int
    Iterm() *Iterm
    Has() bool
    Next()
    Prev()
    Reset()
    End()
}

type Aggregate interface {
    Iterator() Iterator
}

type Iterm struct {
    Key   Value
    Value Object
}
