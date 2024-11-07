package w0filter

import (
	"context"
	"errors"

	wa "github.com/tetratelabs/wazero/api"

	util "github.com/takanoriyanagitani/go-cbor-filter/util"
	//pa "github.com/takanoriyanagitani/go-cbor-filter/util/pair"
)

var (
	ErrInvalidFilter error = errors.New("invalid filter")
)

type Filter struct {
	wa.Function
}

func (f Filter) ToIo() util.IO[[]uint64] {
	return func(ctx context.Context) ([]uint64, error) {
		return f.Call(ctx)
	}
}

func (f Filter) ToIoSingle() util.IO[uint64] {
	return util.ComposeIoErr(
		f.ToIo(),
		func(results []uint64) (uint64, error) {
			switch len(results) {
			case 1:
				return results[0], nil
			default:
				return 0, ErrInvalidFilter
			}
		},
	)
}

func (f Filter) ToIoFilter() util.IO[bool] {
	return util.ComposeIo(
		f.ToIoSingle(),
		func(u uint64) bool { return 0 != u },
	)
}

func (f Filter) FilterArray(ctx context.Context) (keep bool, e error) {
	return f.ToIoFilter()(ctx)
}
