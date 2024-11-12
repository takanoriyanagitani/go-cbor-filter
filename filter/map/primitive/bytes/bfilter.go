package bfilter

import (
	"bytes"
	"time"

	mp "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive"
)

type FilterBytesExact struct {
	Value []byte
}

func (f FilterBytesExact) KeepBytes(b []byte) (keep bool) {
	return 0 == bytes.Compare(b, f.Value)
}

func (f FilterBytesExact) KeepUnsigned(u uint64) (keep bool) { return false }
func (f FilterBytesExact) KeepBool(b bool) (keep bool)       { return false }
func (f FilterBytesExact) KeepSigned(_ int64) (keep bool)    { return false }
func (f FilterBytesExact) KeepFloat(_ float64) (keep bool)   { return false }
func (f FilterBytesExact) KeepString(_ string) (keep bool)   { return false }
func (f FilterBytesExact) KeepNull() (keep bool)             { return false }
func (f FilterBytesExact) KeepTime(_ time.Time) (keep bool)  { return false }
func (f FilterBytesExact) KeepUndefined() (keep bool)        { return false }
func (f FilterBytesExact) KeepAny(_ any) (keep bool)         { return false }

func (f FilterBytesExact) AsFilterVal() mp.FilterVal { return f }
