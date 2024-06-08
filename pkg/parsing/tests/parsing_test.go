package tests

import (
	"testing"

	"github.com/racg0092/rhombifer/pkg/models"
	"github.com/racg0092/rhombifer/pkg/parsing"
)

func TestFlagsLookup(t *testing.T) {
	input := "-rfs"
	flags := make(map[string]models.Flag)
  flags["r"] = models.Flag{
    Name: "recursive"
  }
	_, err := parsing.FlagsLookup(flags, input)
	if err != nil {
		t.Error(err)
	}
}
