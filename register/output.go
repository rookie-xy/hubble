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
        panic(fmt.Sprintf("output '%v' already registered ", name))
    }

    output.Outputs[name] = o
}
