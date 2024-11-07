package w0filter

import (
	"errors"

	wa "github.com/tetratelabs/wazero/api"
)

var (
	ErrUnableToWrite error = errors.New("unable to write")
)

type Memory struct {
	wa.Memory
}

func (m Memory) Write(offset uint32, data []byte) error {
	var ok bool = m.Memory.Write(offset, data)
	switch ok {
	case true:
		return nil
	default:
		return ErrUnableToWrite
	}
}
