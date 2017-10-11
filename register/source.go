package register

import (
    "fmt"
    "strings"
    "github.com/rookie-xy/hubble/source"
)

func Source(name string, f source.Factory) {
    name = name[strings.LastIndex(name, ".") + 1:]
    if name == "" {
        return
    }

    if _, exists := source.Sources[name]; exists {
        panic(fmt.Sprintf("codec '%v' already registered ", name))
    }

    source.Sources[name] = f
}
