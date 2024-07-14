package rhombifer

import (
	"os"
	"strings"
	"testing"

	rhombi "github.com/racg0092/rhombifer"
	"github.com/racg0092/rhombifer/pkg/builtin"
	"github.com/racg0092/rhombifer/pkg/models"
)

// sample user input
func mimicOsArgs() []string {
	input := "./myprogram --lol"

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
	os.Args = mimicOsArgs()
	root := rhombi.Root()

	t.Run("running root with no values", func(t *testing.T) {
		if err := rhombi.Start(); err != nil {
			t.Error(err)
		}
	})

	t.Run("running root wiht not values and help as default", func(t *testing.T) {
		rhombi.RunHelpIfNoInput = true
		help := builtin.HelpCommand(nil, nil)
		root.AddSub(help)
		if err := rhombi.Start(); err != nil {
			t.Error(err)
		}
	})
}
