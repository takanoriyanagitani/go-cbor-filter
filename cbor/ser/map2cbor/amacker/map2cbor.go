package map2cbor

import (
	"context"
	"io"

	fc "github.com/fxamacker/cbor/v2"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
)

type MapToCborToOutput struct {
	*fc.Encoder
}

func (o MapToCborToOutput) OutputMap(_ context.Context, m cf.CborMap) error {
	return o.Encoder.Encode(m)
}

func (o MapToCborToOutput) AsMapOutput() cf.MapOutput { return o.OutputMap }

func MapOutputFromWriter(wtr io.Writer) cf.MapOutput {
	return MapToCborToOutput{Encoder: fc.NewEncoder(wtr)}.AsMapOutput()
}
