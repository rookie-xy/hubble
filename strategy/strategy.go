package strategy


/*
import (
    "context"
)

const (
      =
)

type Context struct {
    context.Context
    Kill context.CancelFunc
}

func NewContext() *Context {
    return &Context{}
}

func CreateContext(action string) context.Context {
    switch action {
    case "backgr"

    }
    return context.Background()
}

func CreateTodoContext() context.Context {
    return context.TODO()
}

func (r *Context) WithCancel(c context.Context) int {
    if this, kill := context.WithCancel(c); this != nil {
        r.Context = this
        r.Kill = kill

    } else {
        return Error
    }

    return Ok
}
*/

type StrategySort interface {
	Sort([]int)
}

type BubbleSort struct {
}

func (self *BubbleSort) Sort(s []int) {
	size := len(s)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
}

type InsertionSort struct {
}

func (self *InsertionSort) Sort(s []int) {
	size := len(s)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff int = s[i]
		for j = i - 1; j >= 0; j-- {
			if s[j] < buff {
				break
			}
			s[j+1] = s[j]
		}
		s[j+1] = buff
	}
}

type Context struct {
	strategy StrategySort
}

func (self *Context) Algorithm(a StrategySort) {
	self.strategy = a
}

func (self *Context) Sort(s []int) {
	self.strategy.Sort(s)
}
