package adapter

import (
	"github.com/rookie-xy/hubble/codec"
	"github.com/rookie-xy/hubble/types"
)

type ValueCodec interface {
	codec.Codec
    ValueDecode(out []byte, atEOF bool) (int, types.Object, error)
}

func ToValueCodec(c codec.Codec) ValueCodec {
    return c.(ValueCodec)
}
