package builtin

import (
	"os"
	"strings"
	"testing"

	"github.com/racg0092/rhombifer"
	// "github.com/racg0092/rhombifer/pkg/models"
)

func FuggazziSubs() []rhombifer.Command {
	commands := make([]rhombifer.Command, 0)

	rcmd := rhombifer.Command{
		Name:     "Recursive",
		LongDesc: "Foo Bar and the woo woo gang do shit together",
	}

	commands = append(commands, rcmd)

	return commands
}

func OsArgs(expand string) []string {
	program := "myprg"
	if "" != expand {
		program = program + " " + expand
	}
	program = strings.ReplaceAll(program, "  ", " ")
	return strings.Split(program, " ")
}

func TestHelpWithValue(t *testing.T) {
	os.Args = OsArgs("help rcmd")
	root := rhombifer.Root()
	root.AddSub(HelpCommand(nil, nil))
	root.Subs["rcmd"] = FuggazziSubs()[0]
	rhombifer.RunHelpIfNoInput = true
	if err := rhombifer.Start(); err != nil {
		t.Error(err)
	}
}

func ExampleHelpCommand() {
	// get root command if any. Warning calling this function
	// creates root command
	root := rhombifer.Root()
	// add built in. can be added to any command
	root.AddSub(HelpCommand(nil, nil))
	// run help command
	root.Subs["help"].Run()
}
