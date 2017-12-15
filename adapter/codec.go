package adapter

import "github.com/rookie-xy/hubble/codec"

type LogDecoder interface {
	codec.Decoder
	LogDecode([]byte, bool) (int, []byte, error)
}

func ToLogDecoder(de codec.Decoder) LogDecoder {
    return de.(LogDecoder)
}
