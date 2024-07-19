package rhombifer

import (
	"fmt"
	"os"
	"strings"
	"testing"

	rhombi "github.com/racg0092/rhombifer"
	"github.com/racg0092/rhombifer/pkg/builtin"
	"github.com/racg0092/rhombifer/pkg/models"
)

// sample user input
func mimicOsArgs(params string) []string {
	var input string
	if params == "" {
		input = "./myprogram"
	} else {
		input = "./myprogram " + params
	}

	args := make([]string, 0)
	args = append(args, strings.Split(input, " ")...)
	return args
}

func AddSampleFlags(cmd *rhombi.Command) {

	r := models.Flag{
		Name:        "recursive",
		ShortFormat: "r",
	}

	cmd.AddFlag(r)
}

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
		rhombi.RunHelpIfNoInput = true
		help := builtin.HelpCommand(nil, nil)
		root.AddSub(help)
		if err := rhombi.Start(); err != nil {
			t.Error(err)
		}
	})

	t.Run("running root with flag", func(t *testing.T) {
		os.Args = mimicOsArgs("--lol")
		rhombi.RunHelpIfNoInput = true
		help := builtin.HelpCommand(nil, nil)
		root.AddSub(help)
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
		rhombi.RunHelpIfNoInput = true
		help := builtin.HelpCommand(nil, nil)
		root.AddSub(help)
		AddSampleFlags(root)
		root.Run = func(a ...string) error {
			fmt.Println("Yay from root")
			return nil
		}
		if err := rhombi.Start(); err != nil {
			t.Error(err)
		}
	})

}
