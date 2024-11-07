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

	ic "github.com/takanoriyanagitani/go-cbor-filter/iter/cbor2iter"
	ca "github.com/takanoriyanagitani/go-cbor-filter/iter/cbor2iter/amacker"

	ow "github.com/takanoriyanagitani/go-cbor-filter/out/wtr"

	aa "github.com/takanoriyanagitani/go-cbor-filter/cbor/ser/arr2cbor/amacker"

	fw "github.com/takanoriyanagitani/go-cbor-filter/filter/wasm"
	ww "github.com/takanoriyanagitani/go-cbor-filter/filter/wasm/wazero"

	ap "github.com/takanoriyanagitani/go-cbor-filter/app/cbor2iter2filtered2output"
)

var rdr2arr ic.ReaderToArrayIterator = ca.RdrToArrIter
var wtr2out ow.WriterToOutputSerialized = ow.WriterToOut

var ser2buf cf.SerializerBuffered = aa.ArraySerializerBufDefault
var serializer cf.Serializer = ser2buf.ToSerializer()

func GetEnvNewByKey(key string) util.IO[string] {
	return func(_ context.Context) (string, error) {
		return os.Getenv(key), nil
	}
}

var WasmFsConfigNew util.IO[fw.WasmFsConfig] = util.ComposeIo(
	GetEnvNewByKey("ENV_WASM_MODULE_DIR"),
	fw.WasmFsConfigNewDefault,
)

var WasmFsSourceNew util.IO[fw.WasmFsSource] = util.ComposeIo(
	WasmFsConfigNew,
	func(cfg fw.WasmFsConfig) fw.WasmFsSource { return cfg.ToWasmSource() },
)

var WasmSourceNew util.IO[fw.WasmSource] = util.ComposeIo(
	WasmFsSourceNew,
	func(fsrc fw.WasmFsSource) fw.WasmSource { return fsrc.ToSource() },
)

type IoConfig struct {
	io.Reader
	io.Writer
}

func (i IoConfig) ToIterSource() util.IO[ic.CborToArrayIterator] {
	return rdr2arr(i.Reader)
}

func (i IoConfig) ToOutputSer() util.IO[cf.OutputSerialized] {
	return wtr2out(i.Writer)
}

func (i IoConfig) Filter2app(filter cf.FilterSerialized) util.IO[ap.App] {
	return func(ctx context.Context) (ap.App, error) {
		c2ai, ei := i.ToIterSource()(ctx)
		oser, eo := i.ToOutputSer()(ctx)
		return ap.App{
			CborToArrayIterator: c2ai,
			FilterSerialized:    filter,
			OutputSerialized:    oser,
			Serializer:          serializer,
		}, errors.Join(ei, eo)
	}
}

func stdin2stdout(ctx context.Context, filter cf.FilterSerialized) error {
	var br io.Reader = bufio.NewReader(os.Stdin)

	var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer bw.Flush()

	icfg := IoConfig{
		Reader: br,
		Writer: bw,
	}

	app, e := icfg.Filter2app(filter)(ctx)
	if nil != e {
		return e
	}
	return app.OutputAll(ctx)
}

func sub(ctx context.Context) error {
	var cfg ww.Config = ww.ConfigDefault
	var rtm ww.Runtime = cfg.ToRuntime(ctx)
	defer rtm.Close(ctx)

	wsrc, e := WasmSourceNew(ctx)
	if nil != e {
		return e
	}

	wasm, e := wsrc(ctx)
	if nil != e {
		return e
	}

	compiled, e := rtm.Compile(ctx, wasm)
	if nil != e {
		return e
	}
	defer compiled.Close(ctx)

	instance, e := compiled.Instantiate(ctx, rtm.Runtime)
	if nil != e {
		return e
	}
	defer instance.Close(ctx)

	filter, e := instance.ToFilter()
	if nil != e {
		return e
	}

	var fser cf.FilterSerialized = filter.AsFilterSerialized()

	return stdin2stdout(ctx, fser)
}

func main() {
	e := sub(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
