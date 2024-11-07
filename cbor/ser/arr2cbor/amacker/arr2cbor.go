package arr2cbor

import (
	"bytes"
	"context"

	fc "github.com/fxamacker/cbor/v2"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
)

func ArrToCborToBuf(
	_ context.Context,
	arr cf.CborArray,
	buf *bytes.Buffer,
) error {
	return fc.MarshalToBuffer(arr, buf)
}

var ArraySerializerBufDefault cf.SerializerBuffered = ArrToCborToBuf
