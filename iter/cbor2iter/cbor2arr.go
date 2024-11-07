package cbor2iter

import (
	"context"
	"io"
	"iter"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
	util "github.com/takanoriyanagitani/go-cbor-filter/util"
)

type CborToArrayIterator func(context.Context) iter.Seq[cf.CborArray]

func (i CborToArrayIterator) ToSerialized(
	ser cf.Serializer,
) CborIterator {
	return func(ctx context.Context) iter.Seq[cf.SerializedArray] {
		var original iter.Seq[cf.CborArray] = i(ctx)
		return func(yield func(cf.SerializedArray) bool) {
			for arr := range original {
				raw, e := ser(ctx, arr)
				if nil != e {
					return
				}
				if !yield(raw) {
					return
				}
			}
		}
	}
}

type CborIterator func(context.Context) iter.Seq[cf.SerializedArray]

func (i CborIterator) ToFiltered(filter cf.FilterSerialized) CborIterator {
	return func(ctx context.Context) iter.Seq[cf.SerializedArray] {
		var original iter.Seq[cf.SerializedArray] = i(ctx)
		return func(yield func(cf.SerializedArray) bool) {
			for item := range original {
				keep, e := filter(ctx, item)
				if nil != e {
					return
				}

				if keep {
					if !yield(item) {
						return
					}
				}
			}
		}
	}
}

func (i CborIterator) OutputAll(
	ctx context.Context,
	o cf.OutputSerialized,
) error {
	var err error = nil
	for item := range i(ctx) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		err = o(ctx, item)
		if nil != err {
			return err
		}
	}
	return nil
}

type ReaderToArrayIterator func(io.Reader) util.IO[CborToArrayIterator]
