package channel

//channel可以用中介模式

/*
type Cycle interface {
    Start()
    Stop()
}

type Routine interface {
    SetCycle(cycle Cycle)
}

type RoutineCycle struct {
    name string
}

func (r *RoutineCycle) Start() {
}

func (r *RoutineCycle) Stop() {
}
*/


type Mediator interface {
	ConnectColleagues()
	Send(msg string)
}

type Colleague interface {
	setMediator(mediator Mediator)
}

type ConcreteMediator struct {
	*Farmer
	*Cannery
	*Shop
}

func (self *ConcreteMediator) ConnectColleagues() {
	self.Farmer.setMediator(self)
	self.Cannery.setMediator(self)
	self.Shop.setMediator(self)
}

func (self *ConcreteMediator) Send(msg string) {
	if msg == "Farmer: Tomato complete..." {
		self.Cannery.money -= 15000.00
		self.Farmer.money += 15000.00
		self.Cannery.MakeKetchup(self.Farmer.GetTomato())
	} else if msg == "Cannery: Ketchup complete..." {
		self.Shop.money -= 30000.00
		self.Cannery.money += 30000.00
		self.Shop.SellKetchup(self.Cannery.GetKetchup())
	}
}

func NewMediator() *ConcreteMediator {
	mediator := &ConcreteMediator{}
	mediator.ConnectColleagues()
	mediator.Farmer.money = 7500.00
	mediator.Cannery.money = 15000.00
	mediator.Shop.money = 30000.00
	return mediator
}

type Farmer struct {
	mediator Mediator
	tomato   int
	money    float64
}

func (self *Farmer) setMediator(mediator Mediator) {
	self.mediator = mediator
}

func (self *Farmer) GrowTomato(tomato int) {
	self.tomato = tomato
	self.money -= 7500.00
	self.mediator.Send("Farmer: Tomato complete...")
}

func (self *Farmer) GetTomato() int {
	return self.tomato
}

type Cannery struct {
	mediator Mediator
	ketchup  int
	money    float64
}

func (self *Cannery) setMediator(mediator Mediator) {
	self.mediator = mediator
}

func (self *Cannery) MakeKetchup(tomato int) {
	self.ketchup = tomato
	self.mediator.Send("Cannery: Ketchup complete...")
}

func (self *Cannery) GetKetchup() int {
	return self.ketchup
}

type Shop struct {
	mediator Mediator
	money    float64
}

func (self *Shop) setMediator(mediator Mediator) {
	self.mediator = mediator
}

func (self *Shop) SellKetchup(ketchup int) {
	self.money = float64(ketchup) * 54.75
}

func (self *Shop) GetMoney() float64 {
	return self.money
}
