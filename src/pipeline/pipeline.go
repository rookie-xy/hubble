package pipeline

import (
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/event"
)

type Factory func(log.Log, int) (Pipeline, error)

type Pipeline interface {
    Clone() Pipeline
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

type Configure struct {
    Name string
    Size int
}
