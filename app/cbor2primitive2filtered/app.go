package app

import (
	"context"

	cf "github.com/takanoriyanagitani/go-cbor-filter"

	ic "github.com/takanoriyanagitani/go-cbor-filter/iter/cbor2iter"

	mp "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive"

	cfg "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive/config"
)

type App struct {
	cfg.Configs
	ic.CborToMapIterator
	cf.MapOutput
}

func (a App) ToFilters() ([]mp.Filter, error) {
	return a.Configs.ToFilters()
}

func (a App) ToFilterMap() (cf.FilterMap, error) {
	f, e := a.ToFilters()
	return mp.Filters(f).ToFilterMap(), e
}

func (a App) ToFilteredIterSource() ic.CborToMapIterator {
	fm, e := a.ToFilterMap()
	if nil != e {
		return a.CborToMapIterator.ToFiltered(cf.FilterMapStaticErr(e))
	}
	return a.CborToMapIterator.ToFiltered(fm)
}

func (a App) OutputAll(ctx context.Context) error {
	return a.ToFilteredIterSource().OutputAll(ctx, a.MapOutput)
}
