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

type CodecPrototype interface {
    Prototype
    codec.Codec
}

func Codec(this codec.Codec) codec.Codec {
    prototype := this.(CodecPrototype)
    return prototype.Clone().(codec.Codec)
}
