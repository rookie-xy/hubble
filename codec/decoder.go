package codec

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
)

type Decoding func(log.Log, types.Value) (Decoder, error)

type Decoder interface {
    Decode([]byte) (types.Object, error)
}

var Decodings = map[string]Decoding{}
