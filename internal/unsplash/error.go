package unsplash

type IllegalArgumentError struct {
	ErrString string
}

func (e *IllegalArgumentError) Error() string {
	return e.ErrString
}
