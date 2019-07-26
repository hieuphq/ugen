package clipboard

import "errors"

// ErrOSNotFound not found any os to match with your machine
var ErrOSNotFound = errors.New("OS not found")

// Readable ..
type Readable interface {
	Read() (string, error)
}

// Writable ..
type Writable interface {
	Write(string) error
}

// RWable ..
type RWable interface {
	Readable
	Writable
}

// NewClipboard by current your os
func NewClipboard(os string) RWable {
	switch os {
	case "darwin":
		return NewForMacOS()

	default:
		return NewForNoob()
	}
}
