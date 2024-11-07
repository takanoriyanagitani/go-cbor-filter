package w0filter

import (
	"context"

	wa "github.com/tetratelabs/wazero/api"
)

type Instance struct {
	wa.Module
	InitializeInputBuffer string
	Filter                string
}

func (i Instance) Close(ctx context.Context) error {
	return i.Module.Close(ctx)
}

func (i Instance) ToFilter() (FilterW0, error) {
	f := FilterW0{
		InitializeInputBuffer: InitializeInputBuffer{
			Function: i.Module.ExportedFunction(i.InitializeInputBuffer),
		},
		Memory: Memory{Memory: i.Module.Memory()},
		Filter: Filter{
			Function: i.Module.ExportedFunction(i.Filter),
		},
	}
	var err error = f.Validate()
	return f, err
}
