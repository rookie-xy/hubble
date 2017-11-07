package adapter

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/pipeline"
    "github.com/rookie-xy/hubble/types"
)

type FileEvent interface {
    event.Event

    GetHeader() types.Map
    GetBody() MessageEvent
    GetFooter() []byte
}

func ToFileEvent(e event.Event) FileEvent {
    return e.(FileEvent)
}

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

type SinceDBEvent interface {
    event.Event
}

func ToSinceDBEvent(e event.Event) SinceDBEvent {
    return e.(SinceDBEvent)
}
