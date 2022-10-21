package clipboard

import "fmt"

type noob struct{}

// NewForNoob new instance for Noob
func NewForNoob() RWable {
	fmt.Println("HP 12")
	return &noob{}
}

func (*noob) Read() (string, error) {
	return "", ErrOSNotFound
}

func (*noob) Write(s string) error {
	return ErrOSNotFound
}
