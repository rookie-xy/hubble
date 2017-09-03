package register

import (
    "fmt"
    "github.com/rookie-xy/hubble/src/pipeline"
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
