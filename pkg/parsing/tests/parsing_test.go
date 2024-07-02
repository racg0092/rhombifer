package tests

import (
	"fmt"
	"github.com/racg0092/rhombifer/pkg/models"
	"github.com/racg0092/rhombifer/pkg/parsing"
	"testing"
)

func TestFlagsLookup(t *testing.T) {

	t.Run("testing short format", func(t *testing.T) {
		input := "-r"
		flags := make([]models.Flag, 0)
		rflag := models.Flag{
			Name:        "recursive",
			ShortFormat: "r",
		}
		flags = append(flags, rflag)
		fls, err := parsing.FlagsLookup(flags, input)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%v\n", fls)
	})

	// t.Run("testing long format", func(t *testing.T) {
	// 	input := "--recursive"
	// 	flags := make([]models.Flag, 0)
	// 	recursive := models.Flag{
	// 		Name:        "recursive",
	// 		ShortFormat: "r",
	// 	}
	// 	flags = append(flags, recursive)
	// 	_, err := parsing.FlagsLookup(flags, input)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// })
	//
}
