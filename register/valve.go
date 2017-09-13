package register


import (
    "fmt"
    "github.com/rookie-xy/hubble/valve"
)

func Valve(name string, f valve.Factory) {
    if name == "" {
        return
    }

    if _, exists := valve.Valves[name]; exists {
        panic(fmt.Sprintf("this valve '%v' already registered ", name))
    }

    valve.Valves[name] = f
}
