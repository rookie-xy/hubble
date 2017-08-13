package codec

import (
    "github.com/rookie-xy/worker/src/prototype"
)

type Factory func(*Config) (Codec, error)

type Codec interface {
    Encode(in prototype.Object) (prototype.Object, error)
    Decode(out []byte) (prototype.Object, error)
}

var Codecs = map[string]Factory{}

type Config struct {
    Name      string
    configure string
}
