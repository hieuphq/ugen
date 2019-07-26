package clipboard

type noob struct{}

// NewForNoob new instance for Noob
func NewForNoob() RWable {
	return &noob{}
}

func (*noob) Read() (string, error) {
	return "", ErrOSNotFound
}

func (*noob) Write(s string) error {
	return ErrOSNotFound
}
