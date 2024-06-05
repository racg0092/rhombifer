package rhombifer

import (
	"fmt"
	"os"
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
	root.AddSub(builtin.HelpCommand(nil, nil))
	root.Subs["help"].Run()
}

func TestHelpPrintWithCommands(t *testing.T) {
	root := rhombifer.Root()
	cmd := rhombifer.Command{
		Name:      "echo",
		ShortDesc: "Prints hello world",
		LongDesc:  "Prints a simple hello world for the hell of it",
		Run: func(args ...string) error {
			fmt.Println("Hello World")
			return nil
		},
	}
	root.AddSub(cmd)
	root.AddSub(builtin.HelpCommand(nil, nil))
	root.Subs["help"].Run()
}

func TestExecCommand(t *testing.T) {
	root := rhombifer.Root()
	if root == nil {
		t.Errorf("Root command not set")
	}
	root.AddSub(builtin.HelpCommand(nil, nil))
	root.Run = func(args ...string) error {
		root.Subs["help"].Run()
		return nil
	}

	t.Run("empty command", func(t *testing.T) {
		err := rhombifer.ExecCommand("")
		if err != nil {
			t.Error(err)
		}
	})

}

func TestStartApp(t *testing.T) {
	os.Args = []string{"self"}
	root := rhombifer.Root()
	root.AddSub(builtin.HelpCommand(nil, nil))
	if err := rhombifer.Start(); err != nil {
		t.Error(err)
	}
}

func ExampleHelpCommand() {
	// get root command if any. Warning calling this function
	// creates root command
	root := rhombifer.Root()
	// add built in. can be added to any command
	root.AddSub(builtin.HelpCommand(nil, nil))
	// run help command
	root.Subs["help"].Run()
}
