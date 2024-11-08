package empty

import (
	"context"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
)

func KeepNonEmptyArray(_ context.Context, arr cf.CborArray) (keep bool, e error) {
	keep = 0 < len(arr)
	return
}

func KeepNonEmptyMap(_ context.Context, m cf.CborMap) (keep bool, e error) {
	keep = 0 < len(m)
	return
}

var FilterArrNonEmpty cf.FilterArray = KeepNonEmptyArray

var FilterMapNonEmpty cf.FilterMap = KeepNonEmptyMap
