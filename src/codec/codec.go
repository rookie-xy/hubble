package codec

import (
    "github.com/rookie-xy/hubble/src/prototype"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/command"
)

type Factory func(log.Log, *command.Command) (Codec, error)

type Codec interface {
    Encode(in prototype.Object) (prototype.Object, error)
    Decode(out []byte) (prototype.Object, error)
}

var Codecs = map[string]Factory{}
