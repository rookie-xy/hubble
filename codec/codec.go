package codec

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
)

type Factory func(log.Log, types.Value) (Codec, error)

type Codec interface {
    Encode(in types.Object) (types.Object, error)
    Decode(out []byte, atEOF bool) (int, types.Object, error)
}

var Codecs = map[string]Factory{}
