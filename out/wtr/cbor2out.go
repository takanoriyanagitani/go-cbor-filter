package wtr

import (
	"context"
	"io"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
	util "github.com/takanoriyanagitani/go-cbor-filter/util"
)

type OutputCbor struct{ io.Writer }

func (o OutputCbor) Write(_ context.Context, raw cf.SerializedArray) error {
	_, e := o.Writer.Write(raw)
	return e
}

func (o OutputCbor) AsOutputSerialized() cf.OutputSerialized {
	return o.Write
}

func WriterToOutput(w io.Writer) util.IO[cf.OutputSerialized] {
	return func(_ context.Context) (cf.OutputSerialized, error) {
		oc := OutputCbor{Writer: w}
		return oc.AsOutputSerialized(), nil
	}
}

var WriterToOut WriterToOutputSerialized = WriterToOutput
