package rhombifer

import (
	"fmt"
	"testing"

	"github.com/racg0092/rhombifer"
	"github.com/racg0092/rhombifer/pkg/builtin"
)

func TestRoot(t *testing.T) {
	rootCmd := rhombifer.Root()
	rootCmd.Run = func(args ...string) error {
		fmt.Println("Yay we are up and running")
		return nil
	}
	rootCmd.Run()
}

func TestBuiltInHelp(t *testing.T) {
	root := rhombifer.Root()
	root.AddSub(builtin.UseBuiltInHelp(nil, nil))
	root.Subs["help"].Run()
}

func ExapleUseBuiltInHelp() {
	// get root command if any. Warning calling this function
	// creates root command
	root := rhombifer.Root()
	// add built in. can be added to any command
	root.AddSub(builtin.UseBuiltInHelp(nil, nil))
	// run help command
	root.Subs["help"].Run()
}
