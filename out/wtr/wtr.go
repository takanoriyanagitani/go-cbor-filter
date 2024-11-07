package wtr

import (
	"io"

	cf "github.com/takanoriyanagitani/go-cbor-filter"
	util "github.com/takanoriyanagitani/go-cbor-filter/util"
)

type WriterToOutputSerialized func(io.Writer) util.IO[cf.OutputSerialized]
