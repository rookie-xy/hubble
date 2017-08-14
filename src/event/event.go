package event

// adapter
type Event interface {
    MakeHeader()
    MakeBody()
    MakeFooter()
    Get()
    Set()
}

type event struct {
    Event
    name string
}

func New() *event {
    return &event{}
}

func (r *event) Set() {
    r.Set()
}

func (r *event) Get() {
    r.Get()
}
