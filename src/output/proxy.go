package output

type Subject interface {
	Send() string
}

type Proxy struct {
	realSubject Subject
}

func (self *Proxy) Send() string {
	if self.realSubject == nil {
		self.realSubject = &RealSubject{}
	}
	return "<strong>" + self.realSubject.Send() + "</strong>"
}

type RealSubject struct {
}

func (self *RealSubject) Send() string {
	return "I’ll be back!"
}
