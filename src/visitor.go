package input


/*
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

type Place interface {
	Accept(v Visitor) string
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
*/
