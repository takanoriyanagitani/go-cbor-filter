package cbor2iter

import (
	"context"
	"io"
	"iter"

	fc "github.com/fxamacker/cbor/v2"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
	util "github.com/takanoriyanagitani/go-cbor-filter/util"

	ca "github.com/takanoriyanagitani/go-cbor-filter/iter/cbor2iter"
)

type CborToArrToIter struct {
	*fc.Decoder
}

func (c CborToArrToIter) ToIter(_ context.Context) iter.Seq[cf.CborArray] {
	var buf cf.CborArray
	var err error
	return func(yield func(cf.CborArray) bool) {
		for {
			clear(buf)
			buf = buf[:0]

			err = c.Decoder.Decode(&buf)
			if nil != err {
				return
			}

			if !yield(buf) {
				return
			}
		}
	}
}

func (c CborToArrToIter) AsCborIterSource() ca.CborToArrayIterator {
	return c.ToIter
}

func IterSourceFromRdr(rdr io.Reader) util.IO[ca.CborToArrayIterator] {
	return func(_ context.Context) (ca.CborToArrayIterator, error) {
		c2a := CborToArrToIter{
			Decoder: fc.NewDecoder(rdr),
		}
		return c2a.AsCborIterSource(), nil
	}
}

var RdrToArrIter ca.ReaderToArrayIterator = IterSourceFromRdr

type CborToMapToIter struct{ *fc.Decoder }

func (c CborToMapToIter) ToIter(_ context.Context) iter.Seq[cf.CborMap] {
	var buf cf.CborMap
	var err error = nil

	return func(yield func(cf.CborMap) bool) {
		for {
			clear(buf)

			err = c.Decoder.Decode(&buf)
			if nil != err {
				return
			}

			if !yield(buf) {
				return
			}
		}
	}
}

func (c CborToMapToIter) AsCborMapSource() ca.CborToMapIterator {
	return c.ToIter
}

func MapSourceFromRdr(rdr io.Reader) util.IO[ca.CborToMapIterator] {
	return func(_ context.Context) (ca.CborToMapIterator, error) {
		c2m := CborToMapToIter{
			Decoder: fc.NewDecoder(rdr),
		}
		return c2m.AsCborMapSource(), nil
	}
}

var RdrToMapIter ca.ReaderToMapIterator = MapSourceFromRdr
