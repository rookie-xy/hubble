package iterator

import "github.com/rookie-xy/worker/src/prototype"

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
