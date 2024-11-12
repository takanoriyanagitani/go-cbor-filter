package main

import (
	"bufio"
	"context"
	"errors"
	"io"
	"log"
	"os"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
	util "github.com/takanoriyanagitani/go-cbor-filter/util"

	ca "github.com/takanoriyanagitani/go-cbor-filter/iter/cbor2iter"
	am "github.com/takanoriyanagitani/go-cbor-filter/iter/cbor2iter/amacker"

	ne "github.com/takanoriyanagitani/go-cbor-filter/empty"

	ma "github.com/takanoriyanagitani/go-cbor-filter/cbor/ser/map2cbor/amacker"

	pc "github.com/takanoriyanagitani/go-cbor-filter/filter/map/primitive/config"

	ap "github.com/takanoriyanagitani/go-cbor-filter/app/cbor2primitive2filtered"
)

var argSource util.IO[[]string] = func(_ context.Context) ([]string, error) {
	return os.Args[1:], nil
}

var rawCfgSource util.IO[[]pc.FilterConfig] = util.ComposeIoErr(
	argSource,
	pc.ConfigsFromStrings,
)

var cfgSource util.IO[pc.Configs] = util.ComposeIoErr(
	rawCfgSource,
	func(cfgs []pc.FilterConfig) (pc.Configs, error) {
		return pc.Configs(cfgs), nil
	},
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
	return func(ctx context.Context) (ap.App, error) {
		cfg, ec := cfgSource(ctx)
		it, ei := i.ToMapIterSource()(ctx)
		return ap.App{
			Configs:           cfg,
			CborToMapIterator: it,
			MapOutput:         i.ToMapOutput(),
		}, errors.Join(ec, ei)
	}
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
