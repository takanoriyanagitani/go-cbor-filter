package bfilter

import (
	"fmt"
	"strconv"
	"time"

	util "github.com/takanoriyanagitani/go-cbor-filter/util"

	mp "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive"
)

type FilterUnsignedExact struct {
	Value uint64
}

func (f FilterUnsignedExact) KeepUnsigned(i uint64) (keep bool) {
	return i == f.Value
}

func (f FilterUnsignedExact) KeepFloat(v float64) (keep bool)  {
	var converted uint64 = uint64(v)
	return float64(converted) == v && converted == f.Value
}

func (f FilterUnsignedExact) KeepBool(b bool) (keep bool)      { return false }
func (f FilterUnsignedExact) KeepSigned(_ int64) (keep bool)   { return false }
func (f FilterUnsignedExact) KeepBytes(_ []byte) (keep bool)   { return false }
func (f FilterUnsignedExact) KeepString(_ string) (keep bool)  { return false }
func (f FilterUnsignedExact) KeepNull() (keep bool)            { return false }
func (f FilterUnsignedExact) KeepTime(_ time.Time) (keep bool) { return false }
func (f FilterUnsignedExact) KeepUndefined() (keep bool)       { return false }
func (f FilterUnsignedExact) KeepAny(_ any) (keep bool)        { return false }

func (f FilterUnsignedExact) AsFilterVal() mp.FilterVal { return f }

func (f FilterUnsignedExact) String() string {
	return fmt.Sprintf("type=unsigned,value=%v", f.Value)
}

func StringToInteger64u(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func FilterUnsignedNew(val uint64) FilterUnsignedExact {
	return FilterUnsignedExact{Value: val}
}

var FilterUnsignedFromString func(c string) (FilterUnsignedExact, error) = util.
	ComposeErr(
		StringToInteger64u,
		util.ErrFn(FilterUnsignedNew),
	)
