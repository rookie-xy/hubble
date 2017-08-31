package pipeline

import (
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/event"
    "github.com/rookie-xy/hubble/src/types"
)

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

var Queues = map[string]Factory{}

/*
package pipeline

import (
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/event"
    "github.com/rookie-xy/hubble/src/types"
)

type Factory func(log.Log, types.Value) (Pipeline, error)

type Pipeline interface {
    // prototype pattern
    Clone() Pipeline
    Close() int
    Push
    Pull
}

type Push interface {
    Push(event.Event) int
}

type Pull interface {
    Pull(size int) (event.Event, int)
}

var Pipelines = map[string]Factory{}
var Publish = map[string]Pipeline{}
var Subscribe = map[string]Pipeline{}
*/

