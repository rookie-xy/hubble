package adapter

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/pipeline"
)

type MessageEvent interface {
    event.Event

    ID() uint64
    GetContent() []byte
    Json() string
}

func ToMessageEvent(e event.Event) MessageEvent {
    return e.(MessageEvent)
}

type PipelineEvent interface {
    event.Event
    pipeline.Queue
}

func ToPipelineEvent(e event.Event) PipelineEvent {
    return e.(PipelineEvent)
}

func Pipeline2Event(Q pipeline.Queue) event.Event {
    pe := Q.(PipelineEvent)
    return pe.(event.Event)
}
