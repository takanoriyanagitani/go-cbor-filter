package cbor2maps2out

import (
	"context"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
	ic "github.com/takanoriyanagitani/go-cbor-filter/iter/cbor2iter"
)

type App struct {
	ic.CborToMapIterator
	cf.FilterMap
	cf.MapOutput
}

func (a App) OutputAll(ctx context.Context) error {
	var filtered = a.
		CborToMapIterator.
		ToFiltered(a.FilterMap)
	return filtered.OutputAll(ctx, a.MapOutput)
}
