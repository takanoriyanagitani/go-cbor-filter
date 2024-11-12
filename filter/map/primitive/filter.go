package pfilter

import (
	"context"
	"time"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
)

type FilterKey string

type FilterVal interface {
	KeepBool(b bool) (keep bool)
	KeepUnsigned(u uint64) (keep bool)
	KeepSigned(s int64) (keep bool)
	KeepFloat(f float64) (keep bool)
	KeepBytes(b []byte) (keep bool)
	KeepString(s string) (keep bool)
	KeepNull() (keep bool)
	KeepTime(t time.Time) (keep bool)
	KeepUndefined() (keep bool)
	KeepAny(a any) (keep bool)
}

type Filter struct {
	FilterKey
	FilterVal
}

func (f Filter) Keep(m map[string]any) (keep bool) {
	if nil == f.FilterVal {
		return false
	}

	val, found := m[string(f.FilterKey)]
	var missing bool = !found
	if missing {
		return f.KeepUndefined()
	}

	switch t := val.(type) {
	case bool:
		return f.FilterVal.KeepBool(t)
	case uint64:
		return f.FilterVal.KeepUnsigned(t)
	case int64:
		return f.FilterVal.KeepSigned(t)
	case float64:
		return f.FilterVal.KeepFloat(t)
	case []byte:
		return f.FilterVal.KeepBytes(t)
	case string:
		return f.FilterVal.KeepString(t)
	case nil:
		return f.FilterVal.KeepNull()
	case time.Time:
		return f.FilterVal.KeepTime(t)
	default:
		return f.FilterVal.KeepAny(val)
	}
}

type Filters []Filter

func (f Filters) Keep(m map[string]any) (keep bool) {
	for _, filt := range f {
		keep = filt.Keep(m)
		if !keep {
			return false
		}
	}
	return true
}

func (f Filters) ToFilterMap() cf.FilterMap {
	return func(_ context.Context, m cf.CborMap) (keep bool, e error) {
		return f.Keep(m), nil
	}
}
