package register

import (
    "fmt"
    "github.com/rookie-xy/hubble/input"
)

func Input(name string, f input.Factory) {
    //name = name[strings.LastIndex(name, ".") + 1:]
    if name == "" {
        return
    }

    if _, exists := input.Inputs[name]; exists {
        fmt.Printf("This input '%v' already registered\n", name)
        return
    }

    input.Inputs[name] = f
}
