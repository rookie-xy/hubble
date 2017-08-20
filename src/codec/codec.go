package codec

import (
    "github.com/rookie-xy/hubble/src/types"
    "github.com/rookie-xy/hubble/src/log"
)

type Factory func(log.Log, types.Value) (Codec, error)

type Codec interface {
    Encode(in types.Object) (types.Object, error)
    Decode(out []byte) (types.Object, error)
}

var Codecs = map[string]Factory{}
