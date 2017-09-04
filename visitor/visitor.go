package visitor
/*
type Visitor interface {
    VisitLog(p *Log) string
}

type Place interface {
   	Accept(v Visitor) string
}

type track struct {

}

func (r *track) VisitLog(l *Log) {
    l.Collect()
}

type Scanner struct {
    places []Place
}

func (r *Scanner) Add(p Place) {
    r.places = append(r.places, p)
}

func (r *Scanner) Accept() string {
    var result string

    for _, p := range r.places {
        p.Accept(track{})
    }

    return result
}

func (r *Scanner) Scan() {

}

type Log struct {
}

func (r *Log) Accept(v Visitor) string {
    return v.VisitLog(r)
}

func (r *Log) Collect() string {
    return "Collect ..."
}

func main() {
    scanner := Scanner{}

    for {
        select {
        case:
            scanner.Scan()
        }
    }
}
*/

package pattern

import (
"fmt"
"math"
)

// 访问者模式（Visitor Pattern） 主要将数据结构与数据操作分离
//
// 1、对象结构中对象对应的类很少改变，但经常需要在此对象结构上定义新的操作。
// 2、需要对一个对象结构中的对象进行很多不同的并且不相关的操作，而需要避免让这些操作"污染"这些对象的类，也不希望在增加新操作时修改这些类。


type Visitor interface {
	Visit(DataStruct)
}

type DataStruct interface {
	Accept(Visitor)
}


type ABData struct {
	A int
	B int
}
func (as *ABData)Accept(vi Visitor){
	vi.Visit(as)
}

type AddVisitor struct {

}

func (av *AddVisitor)Visit(dataS DataStruct){
	data:=dataS.(*ABData)
	sum:=data.A+data.B
	fmt.Println("A+B=",sum)
}

type SubVisitor struct {

}

func (sv *SubVisitor)Visit(dataS DataStruct){
	data:=dataS.(*ABData)
	sum:=data.A-data.B
	fmt.Println("abs(A-B)=",math.Abs(float64(sum)))
}

func VisitorTest(){
	Data:=&ABData{A:8,B:10}
	add:=&AddVisitor{}
	sub:=&SubVisitor{}

	Data.Accept(add)
	Data.Accept(sub)
}

type People struct {
}

func (self *People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

func (self *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

func (self *People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

type City struct {
	places []Place //места посещения
}

func (self *City) Add(p Place) {
	self.places = append(self.places, p)
}

func (self *City) Accept(v Visitor) string {
	var result string
	for _, p := range self.places {
		result += p.Accept(v)
	}
	return result
}

type SushiBar struct {
}

func (self *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(self)
}

func (self *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

type Pizzeria struct {
}

func (self *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(self)
}

func (self *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

type BurgerBar struct {
}

func (self *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(self)
}

func (self *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}
