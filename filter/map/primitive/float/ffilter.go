package bfilter

import (
	"strconv"
	"time"

	util "github.com/takanoriyanagitani/go-cbor-filter/util"

	mp "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive"
)

type FilterFloatExact struct {
	Value float64
}

func (f FilterFloatExact) KeepFloat(d float64) (keep bool) {
	return d == f.Value
}

func (f FilterFloatExact) KeepUnsigned(u uint64) (keep bool) {
	var converted uint64 = uint64(f.Value)
	return float64(converted) == f.Value && converted == u
}

func (f FilterFloatExact) KeepSigned(i int64) (keep bool) {
	var converted int64 = int64(f.Value)
	return float64(converted) == f.Value && converted == i
}

func (f FilterFloatExact) KeepBool(b bool) (keep bool)      { return false }
func (f FilterFloatExact) KeepBytes(_ []byte) (keep bool)   { return false }
func (f FilterFloatExact) KeepString(_ string) (keep bool)  { return false }
func (f FilterFloatExact) KeepNull() (keep bool)            { return false }
func (f FilterFloatExact) KeepTime(_ time.Time) (keep bool) { return false }
func (f FilterFloatExact) KeepUndefined() (keep bool)       { return false }
func (f FilterFloatExact) KeepAny(_ any) (keep bool)        { return false }

func (f FilterFloatExact) AsFilterVal() mp.FilterVal { return f }

func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func FilterFloatNew(val float64) FilterFloatExact {
	return FilterFloatExact{Value: val}
}

var FilterFloatFromString func(c string) (FilterFloatExact, error) = util.
	ComposeErr(
		StringToFloat,
		util.ErrFn(FilterFloatNew),
	)
