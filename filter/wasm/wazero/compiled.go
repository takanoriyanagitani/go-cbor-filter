package w0filter

import (
	"context"

	w0 "github.com/tetratelabs/wazero"
)

type Compiled struct {
	w0.CompiledModule
	w0.ModuleConfig
	InitializeInputBuffer string
	Filter                string
}

func (c Compiled) Close(ctx context.Context) error {
	return c.CompiledModule.Close(ctx)
}

func (c Compiled) Instantiate(
	ctx context.Context,
	rtm w0.Runtime,
) (Instance, error) {
	mdl, e := rtm.InstantiateModule(ctx, c.CompiledModule, c.ModuleConfig)
	return Instance{
		Module:                mdl,
		InitializeInputBuffer: c.InitializeInputBuffer,
		Filter:                c.Filter,
	}, e
}
