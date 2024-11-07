package w0filter

import (
	"context"

	w0 "github.com/tetratelabs/wazero"
)

type Config struct {
	w0.ModuleConfig
	InitializeInputBuffer string
	Filter                string
}

func (c Config) ToRuntime(ctx context.Context) Runtime {
	var rtm w0.Runtime = w0.NewRuntime(ctx)
	return Runtime{
		Runtime:               rtm,
		ModuleConfig:          c.ModuleConfig,
		InitializeInputBuffer: c.InitializeInputBuffer,
		Filter:                c.Filter,
	}
}

var ConfigDefault Config = Config{
	ModuleConfig:          w0.NewModuleConfig().WithName(""),
	InitializeInputBuffer: "initialize_input_buffer",
	Filter:                "filter",
}
