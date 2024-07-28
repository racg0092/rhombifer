package tests

import (
	"fmt"
	rhombi "github.com/racg0092/rhombifer"
	"github.com/racg0092/rhombifer/pkg/builtin"
	"os"
	"testing"
)

func TestRootAndExe(t *testing.T) {
	root := rhombi.Root()

	t.Run("running root with no values", func(t *testing.T) {
		os.Args = mimicOsArgs("")
		if err := rhombi.Start(); err != nil {
			t.Error(err)
		}
	})

	t.Run("running root wiht not values and help as default", func(t *testing.T) {
		os.Args = mimicOsArgs("")
		rhombi.GetConfig().RunHelpIfNoInput = true
		help := builtin.HelpCommand(nil, nil)
		root.AddSub(&help)
		if err := rhombi.Start(); err != nil {
			t.Error(err)
		}
	})

	t.Run("running root with flag", func(t *testing.T) {
		os.Args = mimicOsArgs("--lol")
		rhombi.GetConfig().RunHelpIfNoInput = true
		help := builtin.HelpCommand(nil, nil)
		root.AddSub(&help)
		root.Run = func(a ...string) error {
			fmt.Println("Yay from root")
			return nil
		}
		if err := rhombi.Start(); err != nil {
			t.Error(err)
		}
	})

	t.Run("running root with flag and values", func(t *testing.T) {
		os.Args = mimicOsArgs("-r foo bar")
		rhombi.GetConfig().RunHelpIfNoInput = true
		help := builtin.HelpCommand(nil, nil)
		root.AddSub(&help)
		addSampleFlags(root)
		root.Run = func(a ...string) error {
			fmt.Println("Yay from root")
			return nil
		}
		if err := rhombi.Start(); err != nil {
			t.Error(err)
		}
	})

}
