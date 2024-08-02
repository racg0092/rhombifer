package tests

import (
	"fmt"
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
func sampleFlags() []*models.Flag {
	flags := make([]*models.Flag, 0)

	r := models.Flag{
		Name:        "recursive",
		ShortFormat: "r",
	}

	b := models.Flag{
		Name:        "bobby",
		ShortFormat: "b",
	}

	flags = append(flags, &r, &b)
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
		input := virtaulArgs("-r hello doo foo lol")
		expect := len(input[1:])
		flags := sampleFlags()
		fls, err := parsing.FlagsLookup(flags, input...)
		if err != nil {
			t.Error(err)
		}
		if len(fls) == 0 {
			t.Errorf("Expected recursive flag to be found. Slice was empty")
		}
		if len(fls[0].Values) != expect {
			t.Errorf("Expected to find one value appended to the flag. But found %d", len(fls[0].Values))
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

	t.Run("testing long format parsing with values", func(t *testing.T) {
		input := virtaulArgs("--recursive hello foo doo lol")
		expect := len(input[1:])
		flags := sampleFlags()
		foundFlags, err := parsing.FlagsLookup(flags, input...)
		if err != nil {
			t.Error(err)
		}
		if foundFlags == nil {
			t.Errorf("Expected to get a slice of flag pointers. Got nil")
		}
		if len(foundFlags[0].Values) != expect {
			t.Errorf("Expected to find one value appended to the flag. But found %d", len(foundFlags[0].Values))
		}
	})

	t.Run("testing multiple flags", func(t *testing.T) {
		input := virtaulArgs("--recursive -b")
		flags := sampleFlags()
		foundFlags, err := parsing.FlagsLookup(flags, input...)
		if err != nil {
			t.Error(err)
		}
		for _, f := range foundFlags {
			fmt.Println(*f)
		}
	})

	t.Run("testing multiple flags with values", func(t *testing.T) {
		input := virtaulArgs("-b lol bob --recursive yay hello")
		flags := sampleFlags()
		foundFlags, err := parsing.FlagsLookup(flags, input...)
		if err != nil {
			t.Error(err)
		}
		for _, f := range foundFlags {
			fmt.Println(*f)
		}
	})

}

func TestSingleValueFlag(t *testing.T) {
	input := virtaulArgs("--recursive lol bob")
	flag := models.Flag{
		Name:        "recursive",
		SingleValue: true,
	}
	flags := make([]*models.Flag, 0)
	flags = append(flags, &flag)
	_, err := parsing.FlagsLookup(flags, input...)
	if err == nil {
		t.Errorf("multiple values for single flag error not identify")
	}
}

func TestExtractionFlagUtil(t *testing.T) {
	input := virtaulArgs("-b lol bob --recurse action doubled -l ignore rest")
	flags, err := parsing.ExtractFlagsFromArgs(input...)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(flags)
	//todo: can be improved
}
