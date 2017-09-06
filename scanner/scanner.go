package scanner

//strategy
type Scanner interface {
    Scan() bool
}

func New() {

}


/*
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
*/
