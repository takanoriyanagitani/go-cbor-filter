package w0filter

import (
	"context"

	w0 "github.com/tetratelabs/wazero"
)

type Runtime struct {
	w0.Runtime
	w0.ModuleConfig
	InitializeInputBuffer string
	Filter                string
}

func (r Runtime) Close(ctx context.Context) error {
	return r.Runtime.Close(ctx)
}

func (r Runtime) Compile(
	ctx context.Context,
	wasm []byte,
) (Compiled, error) {
	compiled, e := r.Runtime.CompileModule(
		ctx,
		wasm,
	)
	return Compiled{
		CompiledModule:        compiled,
		ModuleConfig:          r.ModuleConfig,
		InitializeInputBuffer: r.InitializeInputBuffer,
		Filter:                r.Filter,
	}, e
}
