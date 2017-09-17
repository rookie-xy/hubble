package pipeline

import (
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/types"
)

const Channel  = "channel"

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
    Enqueue(event.Event) int
}

type dequeue interface {
    Dequeue(size int) (event.Event, int)
}

type requeue interface {
    Requeue(event.Event) int
}

var Factories = map[string]Factory{}
var Queues    = map[string]Queue{}
