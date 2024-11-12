package bfilter

import (
	"time"

	mp "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive"
)

type FilterStringExact struct {
	Value string
}

func (f FilterStringExact) KeepString(s string) (keep bool) {
	return s == f.Value
}

func (f FilterStringExact) KeepBytes(_ []byte) (keep bool)    { return false }
func (f FilterStringExact) KeepUnsigned(u uint64) (keep bool) { return false }
func (f FilterStringExact) KeepBool(b bool) (keep bool)       { return false }
func (f FilterStringExact) KeepSigned(_ int64) (keep bool)    { return false }
func (f FilterStringExact) KeepFloat(_ float64) (keep bool)   { return false }
func (f FilterStringExact) KeepNull() (keep bool)             { return false }
func (f FilterStringExact) KeepTime(_ time.Time) (keep bool)  { return false }
func (f FilterStringExact) KeepUndefined() (keep bool)        { return false }
func (f FilterStringExact) KeepAny(_ any) (keep bool)         { return false }

func (f FilterStringExact) AsFilterVal() mp.FilterVal { return f }

func FilterStringNew(val string) FilterStringExact {
	return FilterStringExact{Value: val}
}
