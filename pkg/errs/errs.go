package errs

type Err struct {
	s string
}

func (e *Err) Error() string {
	return e.s
}

func NewError(s string) *Err {
	return &Err{s}
}

// Appends message s to default error message
func (e *Err) AppendMessage(s string) {
	e.s = e.s + " " + s
}
