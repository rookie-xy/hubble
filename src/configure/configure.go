package configure

const Name = "configure"

var Event chan []byte = make(chan []byte)

/*
type Configure struct {
    log.Log
    observers []observer.Observer
    Data prototype.Object
}

func New(log log.Log) *Configure {
    return &Configure{
        Log: log,
    }
}

func (r *Configure) Attach(obs observer.Observer) {
    r.observers = append(r.observers, obs)
}

func (r *Configure) Notify() {
    for _, observer := range r.observers {
         observer.Update(r.Data)
    }
}
*/
