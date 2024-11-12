package bfilter

import (
	"strconv"
	"time"

	util "github.com/takanoriyanagitani/go-cbor-filter/util"

	mp "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive"
)

type FilterBool struct {
	Value bool
}

func (f FilterBool) KeepBool(b bool) (keep bool) { return b == f.Value }

func (f FilterBool) KeepUnsigned(_ uint64) (keep bool) { return false }
func (f FilterBool) KeepSigned(_ int64) (keep bool)    { return false }
func (f FilterBool) KeepFloat(_ float64) (keep bool)   { return false }
func (f FilterBool) KeepBytes(_ []byte) (keep bool)    { return false }
func (f FilterBool) KeepString(_ string) (keep bool)   { return false }
func (f FilterBool) KeepNull() (keep bool)             { return false }
func (f FilterBool) KeepTime(_ time.Time) (keep bool)  { return false }
func (f FilterBool) KeepUndefined() (keep bool)        { return false }
func (f FilterBool) KeepAny(_ any) (keep bool)         { return false }

func (f FilterBool) AsFilterVal() mp.FilterVal { return f }

func FilterBoolNew(val bool) FilterBool { return FilterBool{Value: val} }

var FilterBoolFromString func(config string) (FilterBool, error) = util.
	ComposeErr(
		strconv.ParseBool,
		util.ErrFn(FilterBoolNew),
	)
