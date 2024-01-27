package rhombifer

import (
	"testing"

	"github.com/racg0092/rhombifer/parsing"
)

func TestArgsValidation(t *testing.T) {
	got := parsing.InputValidation()

	if got != nil {
		t.Errorf("got %q want %q", got.Error(), "nil")
	}
}
