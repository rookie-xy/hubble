package register

import (
    "fmt"
    "github.com/rookie-xy/hubble/output"
)

func Output(name string, o output.Factory) {
//    name = name[strings.LastIndex(name, ".") + 1:]
    if name == "" {
        return
    }

    if _, exists := output.Outputs[name]; exists {
        fmt.Printf("This output '%v' already registered\n", name)
        return
    }

    output.Outputs[name] = o
}
