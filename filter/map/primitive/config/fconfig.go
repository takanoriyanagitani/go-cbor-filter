package fconfig

import (
	"errors"
	"log"
	"strings"

	mp "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive"

	bfilter "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive/bool"
	ffilter "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive/float"
	ifilter "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive/signed"
	sfilter "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive/string"
	ufilter "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive/unsigned"
)

var (
	ErrInvalidFilter error = errors.New("invalid filter")
	ErrInvalidType   error = errors.New("invalid filter type")
)

type FilterType string

func (f FilterType) String() string { return string(f) }

const (
	FilterTypeBool FilterType = "b"

	FilterTypeUnsigned FilterType = "u"

	FilterTypeSigned FilterType = "s"

	FilterTypeFloat FilterType = "f"

	FilterTypeBytes FilterType = "B"

	FilterTypeString FilterType = "S"

	FilterTypeNull FilterType = "n"

	FilterTypeTime FilterType = "t"

	FilterTypeUndefined FilterType = "U"

	FilterTypeAny FilterType = "A"
)

func FilterTypeFromString(s string) (FilterType, error) {
	switch FilterType(s) {
	case FilterTypeBool:
		return FilterTypeBool, nil
	case FilterTypeUnsigned:
		return FilterTypeUnsigned, nil
	case FilterTypeSigned:
		return FilterTypeSigned, nil
	case FilterTypeFloat:
		return FilterTypeFloat, nil
	case FilterTypeBytes:
		return FilterTypeBytes, nil
	case FilterTypeString:
		return FilterTypeString, nil
	case FilterTypeNull:
		return FilterTypeNull, nil
	case FilterTypeTime:
		return FilterTypeTime, nil
	case FilterTypeUndefined:
		return FilterTypeUndefined, nil
	case FilterTypeAny:
		return FilterTypeAny, nil
	default:
		return FilterTypeBool, ErrInvalidType
	}
}

type FilterConfig struct {
	Key string
	FilterType
	Config string
}

func (c FilterConfig) ToFilterVal() (mp.FilterVal, error) {
	switch c.FilterType {
	case FilterTypeBool:
		return bfilter.FilterBoolFromString(c.Config)
	case FilterTypeUnsigned:
		return ufilter.FilterUnsignedFromString(c.Config)
	case FilterTypeSigned:
		return ifilter.FilterSignedFromString(c.Config)
	case FilterTypeFloat:
		return ffilter.FilterFloatFromString(c.Config)
	case FilterTypeString:
		return sfilter.FilterStringNew(c.Config), nil
	default:
		log.Printf("invalid filter. type: %v\n", c.FilterType)
		return nil, ErrInvalidFilter
	}
}

func (c FilterConfig) ToFilter() (mp.Filter, error) {
	v, e := c.ToFilterVal()
	return mp.Filter{
		FilterKey: mp.FilterKey(c.Key),
		FilterVal: v,
	}, e
}

func FilterConfigNew(key, typ, cfg string) (FilterConfig, error) {
	ftyp, e := FilterTypeFromString(typ)
	return FilterConfig{
		Key:        key,
		FilterType: ftyp,
		Config:     cfg,
	}, e
}

func FilterConfigFromStrings(cfg []string) (FilterConfig, error) {
	if 3 != len(cfg) {
		log.Printf("unexpected config: %v\n", cfg)
		return FilterConfig{}, ErrInvalidFilter
	}
	return FilterConfigNew(
		cfg[0],
		cfg[1],
		cfg[2],
	)
}

// Parses the filter config.
//
// Format: [key]=[typ]:[val]
//
// Examples
//
//	| filter type | key      | config | config string     |
//	|:-----------:|:--------:|:------:|:-----------------:|
//	| bool        | removed  | false  | removed=b:false   |
//	| unsigned    | height   | 634    | height=u:634      |
//	| signed      | id       | -1     | id=i:-1           |
//	| float       | distance | 42.195 | distance=f:42.195 |
//	| string      | name     | helo   | name=S:helo       |
func FilterConfigFromString(cfg string) (FilterConfig, error) {
	var pair []string = strings.SplitN(cfg, "=", 2)
	if 2 != len(pair) {
		log.Printf("unexpected config: %v\n", cfg)
		return FilterConfig{}, ErrInvalidFilter
	}

	var key string = pair[0]
	var config string = pair[1]

	pair = strings.SplitN(config, ":", 2)
	if 2 != len(pair) {
		log.Printf("unexpected config: %v\n", cfg)
		return FilterConfig{}, ErrInvalidFilter
	}

	return FilterConfigNew(
		key,
		pair[0],
		pair[1],
	)
}

func ConfigsFromStrings(configs []string) ([]FilterConfig, error) {
	var ret []FilterConfig
	for _, cfg := range configs {
		parsed, e := FilterConfigFromString(cfg)
		if nil != e {
			return nil, e
		}
		ret = append(ret, parsed)
	}
	return ret, nil
}

type Configs []FilterConfig

func (c Configs) ToFilters() ([]mp.Filter, error) {
	var ret []mp.Filter
	for _, cfg := range c {
		converted, e := cfg.ToFilter()
		if nil != e {
			return nil, e
		}
		ret = append(ret, converted)
	}
	return ret, nil
}

func ConfigsToFilters(cfgs Configs) ([]mp.Filter, error) {
	return cfgs.ToFilters()
}
