package w0filter

import (
	"context"
	"errors"

	wa "github.com/tetratelabs/wazero/api"

	util "github.com/takanoriyanagitani/go-cbor-filter/util"
)

var (
	ErrInvalidInitializer error = errors.New("invalid initializer")
	ErrUnableToInitialize error = errors.New("unable to initialize")
)

type InitializeInputBuffer struct {
	wa.Function
}

func (i InitializeInputBuffer) InitializeRaw(
	ctx context.Context,
	size uint64,
	init uint64,
) ([]uint64, error) {
	return i.Function.Call(ctx, size, init)
}

func (i InitializeInputBuffer) ToIo(sz uint64, init uint64) util.IO[[]uint64] {
	return func(ctx context.Context) ([]uint64, error) {
		return i.InitializeRaw(ctx, sz, init)
	}
}

func (i InitializeInputBuffer) ToSingle(sz uint64, init uint64) util.IO[uint64] {
	return util.ComposeIoErr(
		i.ToIo(sz, init),
		func(results []uint64) (uint64, error) {
			switch len(results) {
			case 1:
				return results[0], nil
			default:
				return 0, ErrInvalidInitializer
			}
		},
	)
}

func (i InitializeInputBuffer) ToOffset(sz uint64, init uint64) util.IO[uint32] {
	return util.ComposeIoErr(
		i.ToSingle(sz, init),
		func(result uint64) (uint32, error) {
			var ioffset int32 = wa.DecodeI32(result)
			if ioffset < 0 {
				return 0, ErrUnableToInitialize
			}
			return uint32(ioffset), nil
		},
	)
}

func (i InitializeInputBuffer) Initialize(
	ctx context.Context,
	sz uint64,
	init uint64,
) (offset uint32, e error) {
	return i.ToOffset(sz, init)(ctx)
}

func (i InitializeInputBuffer) InitializeDefault(
	ctx context.Context,
	sz uint64,
) (offset uint32, e error) {
	return i.Initialize(ctx, sz, 0)
}
