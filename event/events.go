package event

type Events interface {
    //Event
    Put(Event) int
    Batch() []Event
}
