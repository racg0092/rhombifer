package tests

import (
	// "fmt"
	"strings"
	"testing"

	"github.com/racg0092/rhombifer/pkg/models"
	"github.com/racg0092/rhombifer/pkg/parsing"
)

// Takes a sintetic input and returns what the command line would feed as the
// arguments
func virtaulArgs(input string) []string {
	return strings.Split(input, " ")
}

// Create sample flags to teste against
func sampleFlags() []models.Flag {
	flags := make([]models.Flag, 0)

	r := models.Flag{
		Name:        "recursive",
		ShortFormat: "r",
	}

	flags = append(flags, r)
	return flags
}

func TestFlagsLookup(t *testing.T) {

	t.Run("testing short format", func(t *testing.T) {
		input := virtaulArgs("-r")
		flags := sampleFlags()
		foundFlags, err := parsing.FlagsLookup(flags, input...)
		if err != nil {
			t.Error(err)
		}
		if len(foundFlags) != 1 {
			t.Errorf("Expected 1 flag found. but found %d", len(foundFlags))
		}
	})

	t.Run("testing short format with values", func(t *testing.T) {
		input := virtaulArgs("-r hello")
		flags := sampleFlags()
		fls, err := parsing.FlagsLookup(flags, input...)
		if err != nil {
			t.Error(err)
		}
		if len(fls) == 0 {
			t.Errorf("Expected recursive flag to be found. Slice was empty")
		}
		if len(fls[0].Values) != 1 {
			t.Errorf("Expected to find one value appended to the flage. But found %d", len(fls[0].Values))
		}
	})

	t.Run("testing long format", func(t *testing.T) {
		input := virtaulArgs("--recursive")
		flags := sampleFlags()
		foundFlags, err := parsing.FlagsLookup(flags, input...)
		if err != nil {
			t.Error(err)
		}
		if foundFlags == nil {
			t.Errorf("Expected to get a slice of flag pointers. Got nil")
		}
	})

	t.Run("testing long format parsing with value", func(t *testing.T) {
		input := virtaulArgs("--recursive hello")
		flags := sampleFlags()
		foundFlags, err := parsing.FlagsLookup(flags, input...)
		if err != nil {
			t.Error(err)
		}
		if foundFlags == nil {
			t.Errorf("Expected to get a slice of flag pointers. Got nil")
		}
		if len(foundFlags[0].Values) != 1 {
			t.Errorf("Expected to find one value appended to the flag. But found %d", len(foundFlags[0].Values))
		}
	})

}
