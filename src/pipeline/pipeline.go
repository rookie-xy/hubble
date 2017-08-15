package pipeline

import (
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/event"
    "github.com/rookie-xy/hubble/src/command"
)

type Factory func(log.Log, *command.Command) (Pipeline, error)

type Pipeline interface {
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
