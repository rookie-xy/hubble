package event

type Events interface {
    Put(Event) int
    Batch() []Event
}
