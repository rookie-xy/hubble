package types

type Carer interface {
	Rase() string
}

type Enginer interface {
	GetSound() string
}

type Car struct {
	engine Enginer
}

func (self *Car) Rase() string {
	return self.engine.GetSound()
}

type EngineSuzuki struct {
}

func (self *EngineSuzuki) GetSound() string {
	return "SssuuuuZzzuuuuKkiiiii"
}

type EngineHonda struct {
}

func (self *EngineHonda) GetSound() string {
	return "HhoooNnnnnnnnnDddaaaaaaa"
}

type EngineLada struct {
}

func (self *EngineLada) GetSound() string {
	return "PhhhhPhhhhPhPhPhPhPh"
}
