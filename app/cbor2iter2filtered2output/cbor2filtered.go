package iter2filtered

import (
	"context"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
	ic "github.com/takanoriyanagitani/go-cbor-filter/iter/cbor2iter"
)

type App struct {
	ic.CborToArrayIterator
	cf.FilterSerialized
	cf.OutputSerialized
	cf.Serializer
}

func (a App) ToSerialized() ic.CborIterator {
	return a.CborToArrayIterator.ToSerialized(a.Serializer)
}

func (a App) OutputAll(ctx context.Context) error {
	var filtered ic.CborIterator = a.ToSerialized().ToFiltered(
		a.FilterSerialized,
	)
	return filtered.OutputAll(ctx, a.OutputSerialized)
}
