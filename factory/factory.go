package factory

import (
    "fmt"

    "github.com/rookie-xy/hubble/codec"
    "github.com/rookie-xy/hubble/observer"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/pipeline"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/input"
    "github.com/rookie-xy/hubble/output"
    "github.com/rookie-xy/hubble/prototype"
)

// factory method
func Input(name string, l log.Log, v types.Value) (input.Input, error) {
    key := "json"
    if name != "" {
        key = name
    }

    factory := input.Inputs[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' input is not available", key)
    }

    return factory(l, v)
}

func Codec(name string, l log.Log, v types.Value) (codec.Codec, error) {
    key := "json"
    if name != "" {
        key = name
    }

    factory := codec.Codecs[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' codec is not available", key)
    }

    return factory(l, v)
}

func Pipeline(name string, l log.Log, v types.Value) (pipeline.Queue, error) {
    key := "channel"
    if name != "" {
        key = name
    }

    factory := pipeline.Factories[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' pipeline is not available", key)
    }

    return factory(l, v)
}

func Subject(name string) observer.Subject {
    key := ""
    if name != key {
        key = name
    }

    subject := observer.Subjects[key]
    if subject == nil {
        fmt.Println("Not found subject:", key)
        return nil
    }

    return subject
}

func Observer(name string) observer.Observer {
    key := ""
    if name != key {
        key = name
    }

    observer := observer.Observers[key]
    if observer == nil {
        fmt.Errorf("'%v' observer is not available", key)
        return nil
    }

    return observer
}

func Queue(name string) pipeline.Queue {
    key := ""
    if name != key {
        key = name
    }

    queue := pipeline.Queues[key]
    if queue == nil {
        fmt.Errorf("'%v' queue is not available", key)
        return nil
    }

    return prototype.Pipeline(queue)
}

func Forward(name string) (proxy.Forward, error) {
    key := ""
    if name != key {
        key = name
    }

    if forward, ok := proxy.Forwards[key]; ok {
        return forward, nil
    }

    return nil, fmt.Errorf("'%v' forward is not available", key)
}

func Client(name string, l log.Log, v types.Value) (proxy.Forward, error) {
    key := ""
    if name != key {
        key = name
    }

    factory := proxy.Clients[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' client is not available", key)
    }

    return factory(l, v)
}

func Output(name string, l log.Log, v types.Value) (output.Output, error) {
    key := ""
    if name != key {
        key = name
    }

    factory := output.Outputs[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' output is not available", key)
    }

    return factory(l, v)
}

func Server(name string, l log.Log, v types.Value) (proxy.Reverse, error) {
    key := ""
    if name != key {
        key = name
    }

    factory := proxy.Reverses[key]
    if factory == nil {
        return nil, fmt.Errorf("'%v' server is not available", key)
    }

    return factory(l, v)
}
