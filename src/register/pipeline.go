package register

import (
    "fmt"
    "github.com/rookie-xy/hubble/src/pipeline"
)

func Pipeline(name string, f pipeline.Factory) {
    if name == "" {
        return
    }

    if _, exists := pipeline.Pipelines[name]; exists {
        panic(fmt.Sprintf("this channel '%v' already registered ", name))
    }

    pipeline.Pipelines[name] = f
}
