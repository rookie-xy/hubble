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
        fmt.Printf("This pipeline '%v' already registered\n", name)
        return
    }

    pipeline.Factories[name] = f
}

func Queue(name string, Q pipeline.Queue) {
    if name == "" {
        return
    }

    if _, exists := pipeline.Queues[name]; exists {
        fmt.Printf("This pipeline object '%v' already registered\n", name)
        return
    }

    pipeline.Queues[name] = Q
}
