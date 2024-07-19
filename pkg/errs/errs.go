package errs

type Err struct {
	s string
}

func (e *Err) Error() string {
	return e.s
}

func NewError(s string) error {
	return &Err{s}
}
