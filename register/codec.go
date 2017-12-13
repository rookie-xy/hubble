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
        fmt.Printf("This codec '%v' already registered\n", name)
        return
    }
    codec.Codecs[name] = f
}

func Encoder(name string, e codec.Encoding) {
    //name = name[strings.LastIndex(name, ".") + 1:]
    if name == "" {
        return
    }

    if _, exists := codec.Encodings[name]; exists {
        fmt.Printf("This encoder '%v' already registered\n", name)
        return
    }
    codec.Encodings[name] = e
}

func Decoder(name string, e codec.Decoding) {
    //name = name[strings.LastIndex(name, ".") + 1:]
    if name == "" {
        return
    }

    if _, exists := codec.Decodings[name]; exists {
        fmt.Printf("This decoder '%v' already registered\n", name)
        return
    }
    codec.Decodings[name] = e
}
