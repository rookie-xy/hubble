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
        panic(fmt.Sprintf("this filter '%v' already registered ", name))
    }

    filter.Filters[name] = f
}
