package codec

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
)

type Encoding func(log.Log, types.Value) (Encoder, error)

type Encoder interface {
    Encode(types.Object) ([]byte, error)
}

var Encodings = map[string]Encoding{}
