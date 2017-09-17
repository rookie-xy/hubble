package register

import (
    "fmt"
    "github.com/rookie-xy/hubble/pipeline"
)

func Pipeline(name string, f pipeline.Factory) {
    if name == "" {
        return
    }

    if _, exists := pipeline.Factories[name]; exists {
        panic(fmt.Sprintf("this pipeline '%v' already registered ", name))
    }

    pipeline.Factories[name] = f
}

func Queue(name string, Q pipeline.Queue) {
    if name == "" {
        return
    }

    if _, exists := pipeline.Queues[name]; exists {
        panic(fmt.Sprintf("this pipeline clone name '%v' already registered ", name))
    }

    pipeline.Queues[name] = Q
}
