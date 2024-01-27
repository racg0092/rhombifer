package utils

const (
	NO_ARGS = -1
)

type RhombiError struct {
	Code    int
	Message string
}

func (e RhombiError) Error() string {
	return e.Message
}
