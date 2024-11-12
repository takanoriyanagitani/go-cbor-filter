package bfilter

import (
	"strconv"
	"time"

	util "github.com/takanoriyanagitani/go-cbor-filter/util"

	mp "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive"
)

type FilterSignedExact struct {
	Value int64
}

func (f FilterSignedExact) KeepSigned(i int64) (keep bool) {
	return i == f.Value
}

func (f FilterSignedExact) KeepUnsigned(u uint64) (keep bool) {
	return 0 < f.Value && u == uint64(f.Value)
}

func (f FilterSignedExact) KeepFloat(v float64) (keep bool)  {
	var converted int64 = int64(v)
	return float64(converted) == v && converted == f.Value
}

func (f FilterSignedExact) KeepBool(b bool) (keep bool)      { return false }
func (f FilterSignedExact) KeepBytes(_ []byte) (keep bool)   { return false }
func (f FilterSignedExact) KeepString(_ string) (keep bool)  { return false }
func (f FilterSignedExact) KeepNull() (keep bool)            { return false }
func (f FilterSignedExact) KeepTime(_ time.Time) (keep bool) { return false }
func (f FilterSignedExact) KeepUndefined() (keep bool)       { return false }
func (f FilterSignedExact) KeepAny(_ any) (keep bool)        { return false }

func (f FilterSignedExact) AsFilterVal() mp.FilterVal { return f }

func StringToInteger64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func FilterSignedNew(val int64) FilterSignedExact {
	return FilterSignedExact{Value: val}
}

var FilterSignedFromString func(cfg string) (FilterSignedExact, error) = util.
	ComposeErr(
		StringToInteger64,
		util.ErrFn(FilterSignedNew),
	)
