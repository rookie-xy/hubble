package prototype

import "github.com/rookie-xy/hubble/codec"

type DecoderPrototype interface {
    Prototype
    codec.Decoder
}

func Decoder(this codec.Decoder) codec.Decoder {
    prototype := this.(DecoderPrototype)
    return prototype.Clone().(codec.Decoder)
}
