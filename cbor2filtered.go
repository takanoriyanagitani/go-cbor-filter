package cbor2filtered

import (
	"bytes"
	"context"
)

type CborArray []any
type SerializedArray []byte

type SerializedMap []byte

type CborAny any

type FilterAny func(context.Context, CborAny) (keep bool, err error)

type CborMap map[string]any

type FilterMap func(context.Context, CborMap) (keep bool, err error)

type FilterArray func(context.Context, CborArray) (keep bool, err error)

type FilterSerialized func(context.Context, SerializedArray) (keep bool, e error)

type Serializer func(context.Context, CborArray) (SerializedArray, error)
type SerializerBuffered func(context.Context, CborArray, *bytes.Buffer) error

type MapSerializer func(context.Context, CborMap) (SerializedMap, error)
type MapSerializerBuffered func(context.Context, CborMap, *bytes.Buffer) error

type MapOutput func(context.Context, CborMap) error

func (m MapSerializerBuffered) ToSerializer() MapSerializer {
	var buf bytes.Buffer
	var err error = nil
	return func(ctx context.Context, mp CborMap) (SerializedMap, error) {
		buf.Reset()
		err = m(ctx, mp, &buf)
		return buf.Bytes(), err
	}
}

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
