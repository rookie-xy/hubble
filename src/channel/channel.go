package channel

import (
    "github.com/rookie-xy/worker/src/log"
    "github.com/rookie-xy/worker/src/event"
)

type Factory func(log.Log, int) (Channel, error)

type Channel interface {
    Clone() Channel
    Push
    Pull
}

type Push interface {
    Push(event.Event) int
}

type Pull interface {
    Pull(size int) (event.Event, int)
}

var Channels = map[string]Factory{}
var Publish = map[string]Channel{}
var Subscribe = map[string]Channel{}

type Configure struct {
    Name string
    Size int
}
