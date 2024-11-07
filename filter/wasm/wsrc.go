package wasm

import (
	"context"
	"io"
	"io/fs"
	"os"

	util "github.com/takanoriyanagitani/go-cbor-filter/util"
)

type WasmSource util.IO[[]byte]

type WasmFsSource struct {
	fs.FS
	WasmBasename string
	WasmMaxBytes int64
}

func (w WasmFsSource) ToSource() WasmSource {
	return func(_ context.Context) ([]byte, error) {
		f, e := w.FS.Open(w.WasmBasename)
		if nil != e {
			return nil, e
		}
		defer f.Close()

		limited := &io.LimitedReader{
			R: f,
			N: w.WasmMaxBytes,
		}
		return io.ReadAll(limited)
	}
}

type WasmFsConfig struct {
	WasmDirName  string
	WasmBasename string
	WasmMaxBytes int64
}

func (c WasmFsConfig) ToWasmSource() WasmFsSource {
	return WasmFsSource{
		FS:           os.DirFS(c.WasmDirName),
		WasmBasename: c.WasmBasename,
		WasmMaxBytes: c.WasmMaxBytes,
	}
}

var WasmFsConfigDefault WasmFsConfig = WasmFsConfig{
	WasmDirName:  "./modules.d",
	WasmBasename: "filter.wasm",
	WasmMaxBytes: 16777216,
}

func WasmFsConfigNewDefault(dirname string) WasmFsConfig {
	var base WasmFsConfig = WasmFsConfigDefault
	base.WasmDirName = dirname
	return base
}
