package prototype

import "github.com/rookie-xy/hubble/codec"

type CodecPrototype interface {
	Prototype
	codec.Codec
}

func Codec(this codec.Codec) codec.Codec {
    prototype := this.(CodecPrototype)
    return prototype.Clone().(codec.Codec)
}
