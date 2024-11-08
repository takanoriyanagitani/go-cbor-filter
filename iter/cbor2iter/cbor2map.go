package cbor2iter

import (
	"context"
	"io"
	"iter"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
	util "github.com/takanoriyanagitani/go-cbor-filter/util"
)

type CborToMapIterator func(context.Context) iter.Seq[cf.CborMap]

type ReaderToMapIterator func(io.Reader) util.IO[CborToMapIterator]

func (i CborToMapIterator) ToFiltered(f cf.FilterMap) CborToMapIterator {
	return func(ctx context.Context) iter.Seq[cf.CborMap] {
		return func(yield func(cf.CborMap) bool) {
			var original = i(ctx)
			for item := range original {
				keep, e := f(ctx, item)
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

func (i CborToMapIterator) OutputAll(
	ctx context.Context,
	out cf.MapOutput,
) error {
	for item := range i(ctx) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		e := out(ctx, item)
		if nil != e {
			return e
		}
	}
	return nil
}
