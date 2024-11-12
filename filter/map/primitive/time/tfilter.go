package bfilter

import (
	"time"

	mp "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive"
)

type FilterTimeExact struct {
	Value time.Time
}

func (f FilterTimeExact) KeepTime(t time.Time) (keep bool) {
	return t == f.Value
}

func (f FilterTimeExact) KeepBytes(b []byte) (keep bool)    { return false }
func (f FilterTimeExact) KeepUnsigned(u uint64) (keep bool) { return false }
func (f FilterTimeExact) KeepBool(b bool) (keep bool)       { return false }
func (f FilterTimeExact) KeepSigned(_ int64) (keep bool)    { return false }
func (f FilterTimeExact) KeepFloat(_ float64) (keep bool)   { return false }
func (f FilterTimeExact) KeepString(_ string) (keep bool)   { return false }
func (f FilterTimeExact) KeepNull() (keep bool)             { return false }
func (f FilterTimeExact) KeepUndefined() (keep bool)        { return false }
func (f FilterTimeExact) KeepAny(_ any) (keep bool)         { return false }

func (f FilterTimeExact) AsFilterVal() mp.FilterVal { return f }
