package register

import (
    "fmt"
//    "strings"
    "github.com/rookie-xy/hubble/codec"
)

func Codec(name string, f codec.Factory) {
    //name = name[strings.LastIndex(name, ".") + 1:]
    if name == "" {
        return
    }

    if _, exists := codec.Codecs[name]; exists {
        panic(fmt.Sprintf("codec '%v' already registered ", name))
    }

    codec.Codecs[name] = f
}
