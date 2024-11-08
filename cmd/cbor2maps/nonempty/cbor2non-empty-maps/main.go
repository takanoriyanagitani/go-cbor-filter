package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
	util "github.com/takanoriyanagitani/go-cbor-filter/util"

	ca "github.com/takanoriyanagitani/go-cbor-filter/iter/cbor2iter"
	am "github.com/takanoriyanagitani/go-cbor-filter/iter/cbor2iter/amacker"

	ne "github.com/takanoriyanagitani/go-cbor-filter/empty"

	ma "github.com/takanoriyanagitani/go-cbor-filter/cbor/ser/map2cbor/amacker"

	ap "github.com/takanoriyanagitani/go-cbor-filter/app/cbor2maps2filtered2out"
)

var rdr2maps ca.ReaderToMapIterator = am.RdrToMapIter
var mfilter cf.FilterMap = ne.FilterMapNonEmpty

type IoConfig struct {
	io.Reader
	io.Writer
}

func (i IoConfig) ToMapIterSource() util.IO[ca.CborToMapIterator] {
	return rdr2maps(i.Reader)
}

func (i IoConfig) ToMapOutput() cf.MapOutput {
	return ma.MapOutputFromWriter(i.Writer)
}

func (i IoConfig) ToAppSource() util.IO[ap.App] {
	return util.ComposeIo(
		i.ToMapIterSource(),
		func(s ca.CborToMapIterator) ap.App {
			return ap.App{
				CborToMapIterator: s,
				FilterMap:         mfilter,
				MapOutput:         i.ToMapOutput(),
			}
		},
	)
}

func rdr2wtr(ctx context.Context, rdr io.Reader, wtr io.Writer) error {
	icfg := IoConfig{
		Reader: rdr,
		Writer: wtr,
	}
	app, e := icfg.ToAppSource()(ctx)
	if nil != e {
		return e
	}
	return app.OutputAll(ctx)
}

func stdin2stdout(ctx context.Context) error {
	var br io.Reader = bufio.NewReader(os.Stdin)
	var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer bw.Flush()
	return rdr2wtr(ctx, br, bw)
}

func main() {
	e := stdin2stdout(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
