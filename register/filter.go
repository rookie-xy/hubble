package register


import (
    "fmt"
    "github.com/rookie-xy/hubble/filter"
)

func Filter(name string, f filter.Factory) {
    if name == "" {
        return
    }

    if _, exists := filter.Filters[name]; exists {
        fmt.Printf("This filter '%v' already registered\n", name)
        return
    }

    filter.Filters[name] = f
}
