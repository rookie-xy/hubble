package codec

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
)

type Factory func(log.Log, types.Value) (Codec, error)

type Codec interface {
    Encoder
    Decoder
}

var Codecs = map[string]Factory{}
