package cbor2filtered

import (
	"bytes"
	"context"
)

type CborArray []any
type SerializedArray []byte

type FilterArray func(context.Context, CborArray) (keep bool, err error)
type FilterSerialized func(context.Context, SerializedArray) (keep bool, e error)

type Serializer func(context.Context, CborArray) (SerializedArray, error)
type SerializerBuffered func(context.Context, CborArray, *bytes.Buffer) error

func (b SerializerBuffered) ToSerializer() Serializer {
	var buf bytes.Buffer
	var err error = nil
	return func(ctx context.Context, arr CborArray) (SerializedArray, error) {
		buf.Reset()
		err = b(ctx, arr, &buf)
		return buf.Bytes(), err
	}
}

func (f FilterSerialized) ToFilter(s Serializer) FilterArray {
	return func(ctx context.Context, arr CborArray) (keep bool, err error) {
		serialized, e := s(ctx, arr)
		if nil != e {
			return false, e
		}
		return f(ctx, serialized)
	}
}

type OutputSerialized func(context.Context, SerializedArray) error
