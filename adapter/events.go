package adapter

import "github.com/rookie-xy/hubble/event"

type Events interface {
    event.Event
    Put(event.Event) int
    Batch() []event.Event
}

func FileEvents(e event.Event) Events {
    return e.(Events)
}
