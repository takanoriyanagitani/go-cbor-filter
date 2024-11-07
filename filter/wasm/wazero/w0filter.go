package w0filter

import (
	"context"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
)

type FilterW0 struct {
	InitializeInputBuffer
	Memory
	Filter
}

func (f FilterW0) Validate() error {
	oks := []bool{
		nil != f.InitializeInputBuffer.Function,
		nil != f.Memory.Memory,
		nil != f.Filter.Function,
	}
	for _, ok := range oks {
		var ng bool = !ok
		if ng {
			return ErrInvalidFilter
		}
	}
	return nil
}

func (f FilterW0) FilterArray(
	ctx context.Context,
	arr cf.SerializedArray,
) (keep bool, e error) {
	var sz int = len(arr)

	offset, e := f.InitializeInputBuffer.InitializeDefault(ctx, uint64(sz))
	if nil != e {
		return false, e
	}

	e = f.Memory.Write(offset, arr)
	if nil != e {
		return false, e
	}

	return f.Filter.FilterArray(ctx)
}

func (f FilterW0) AsFilterSerialized() cf.FilterSerialized {
	return f.FilterArray
}
