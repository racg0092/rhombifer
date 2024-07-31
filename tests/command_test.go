package tests

import (
	// "fmt"
	"github.com/racg0092/rhombifer"
	"testing"
)

func TestFlagValidation(t *testing.T) {

	t.Run("validating short flag required", func(t *testing.T) {
		osargs := mimicOsArgs("-r")
		cmd := rhombifer.Command{
			Name: "foo",
		}
		addSampleFlags(&cmd)
		valid := cmd.ValidateRequiredFlags(osargs)
		if !valid {
			t.Errorf("expected true but got %v", valid)
		}
	})

	t.Run("validating long flag required", func(t *testing.T) {
		osargs := mimicOsArgs("--recursive")
		cmd := rhombifer.Command{
			Name: "foo",
		}
		addSampleFlags(&cmd)
		valid := cmd.ValidateRequiredFlags(osargs)
		if !valid {
			t.Errorf("expected true but got %v", valid)
		}
	})

}
