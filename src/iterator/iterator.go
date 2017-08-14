package iterator

import "github.com/rookie-xy/hubble/src/prototype"

type Iterator interface {
   	Index() int
   	Value() prototype.Object
   	Has() bool
   	Next()
   	Prev()
   	Reset()
   	End()
}

type Aggregate interface {
   	Iterator() Iterator
}
