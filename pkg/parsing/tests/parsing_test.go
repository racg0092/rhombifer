package tests

import (
	"testing"

	"github.com/racg0092/rhombifer/pkg/models"
	"github.com/racg0092/rhombifer/pkg/parsing"
)

func TestFlagsLookup(t *testing.T) {

	t.Run("testing short format", func(t *testing.T) {
		input := "-r"
		flags := make(map[string]models.Flag)
		flags["r"] = models.Flag{
			Name:        "recursive",
			ShortFormat: "r",
		}
		_, err := parsing.FlagsLookup(flags, input)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("testing long format", func(t *testing.T) {
		input := "--recursive"
		flags := make(map[string]models.Flag)
		flags["recursive"] = models.Flag{
			Name:        "recursive",
			ShortFormat: "r",
		}
		_, err := parsing.FlagsLookup(flags, input)
		if err != nil {
			t.Error(err)
		}
	})
}
