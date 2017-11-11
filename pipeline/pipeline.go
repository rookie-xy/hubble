package pipeline

import (
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/types"
)

//mediator
type Factory func(log.Log, types.Value) (Queue, error)

type Queue interface {
    // prototype pattern
    Clone() Queue
    Close() int

    enqueue
    dequeue
    requeue
}

type enqueue interface {
    Enqueue(event.Event) error
}

type dequeue interface {
    Dequeue() (event.Event, error)
    Dequeues(size int) ([]event.Event, error)
}

type requeue interface {
    Requeue(event.Event) error
}

var Factories = map[string]Factory{}
var Queues    = map[string]Queue{}
