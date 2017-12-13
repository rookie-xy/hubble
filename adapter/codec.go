package adapter

import "github.com/rookie-xy/hubble/codec"
/*
type ValueCodec interface {
	codec.Codec
    ValueDecode(out []byte, atEOF bool) (int, types.Object, error)
}

func ToValueCodec(c codec.Codec) ValueCodec {
    return c.(ValueCodec)
}
*/

type LogDecoder interface {
	codec.Decoder
	LogDecode([]byte, bool) (int, []byte, error)
}

func ToLogDecoder(de codec.Decoder) LogDecoder {
    return de.(LogDecoder)
}
