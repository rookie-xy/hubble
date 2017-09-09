package register

import (
    "fmt"
    "github.com/rookie-xy/hubble/pipeline"
)

func Pipeline(name string, f pipeline.Factory) {
    if name == "" {
        return
    }

    if _, exists := pipeline.Queues[name]; exists {
        panic(fmt.Sprintf("this pipeline '%v' already registered ", name))
    }

    pipeline.Queues[name] = f
}

func Clones(name string, q pipeline.Queue) {
    if name == "" {
        return
    }

    if _, exists := pipeline.Clones[name]; exists {
        panic(fmt.Sprintf("this pipeline name '%v' already registered ", name))
    }

    pipeline.Clones[name] = q
}
