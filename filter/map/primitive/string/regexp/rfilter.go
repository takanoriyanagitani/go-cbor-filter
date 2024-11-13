package rfilter

import (
	"regexp"
	"time"

	util "github.com/takanoriyanagitani/go-cbor-filter/util"

	mp "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive"
)

type FilterStringRegexp struct {
	*regexp.Regexp
}

func (f FilterStringRegexp) KeepString(s string) (keep bool) {
	return f.Regexp.MatchString(s)
}

func (f FilterStringRegexp) KeepBytes(_ []byte) (keep bool)    { return false }
func (f FilterStringRegexp) KeepUnsigned(u uint64) (keep bool) { return false }
func (f FilterStringRegexp) KeepBool(b bool) (keep bool)       { return false }
func (f FilterStringRegexp) KeepSigned(_ int64) (keep bool)    { return false }
func (f FilterStringRegexp) KeepFloat(_ float64) (keep bool)   { return false }
func (f FilterStringRegexp) KeepNull() (keep bool)             { return false }
func (f FilterStringRegexp) KeepTime(_ time.Time) (keep bool)  { return false }
func (f FilterStringRegexp) KeepUndefined() (keep bool)        { return false }
func (f FilterStringRegexp) KeepAny(_ any) (keep bool)         { return false }

func (f FilterStringRegexp) AsFilterVal() mp.FilterVal { return f }

func StringToPattern(s string) (*regexp.Regexp, error) {
	return regexp.Compile(s)
}

func FilterRegexpNew(pat *regexp.Regexp) FilterStringRegexp {
	return FilterStringRegexp{Regexp: pat}
}

var FilterRegexpFromString func(c string) (FilterStringRegexp, error) = util.
	ComposeErr(
		StringToPattern,
		util.ErrFn(FilterRegexpNew),
	)
