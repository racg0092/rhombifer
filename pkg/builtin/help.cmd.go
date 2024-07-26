// Package builtin provide some common commands use by cli applications.
//
// The packages should be used if you intend to add some of the builtin commands
// provided
package builtin

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/racg0092/rhombifer"
	text "github.com/racg0092/rhombifer/pkg/text"
)

// Generates a default built in `help` command. It takes an optional `short` and `long`
// parameters for command descriptions. If non are provided default built in are used.
//
//   - `short` default = "help"
//   - `long` default = "Displays help information"
func HelpCommand(short, long *string) rhombifer.Command {
	help := rhombifer.Command{
		Name:      "help",
		ShortDesc: "Displays help information",
		LongDesc: `
		Displays help information for the specified command or the root command if no command is specified.
		`,
		Leaf: true,
		Run: func(args ...string) error {
			root := rhombifer.Root()
			if len(args) == 0 {
				fmt.Printf("\n%s", text.LightGray(strings.ToUpper(root.Name)))
				if root.LongDesc != "" {
					fmt.Printf("%s", root.LongDesc)
				} else {
					fmt.Printf("%s", root.ShortDesc)
				}
				w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
				fmt.Fprintf(w, "\n%v", text.Bold("Commands"))
				for _, sub := range root.Subs {
					fmt.Fprintf(w, "\n\t%s\t%s", sub.Name, sub.ShortDesc)
				}
				fmt.Fprintf(w, "\n\n")
				w.Flush()
			} else {
				if root == nil {
					return fmt.Errorf("Root is not defined")
				}
				cmd, found := root.Subs[args[0]]
				if !found {
					return fmt.Errorf("Command %s is not recognized", args[0])
				}
				subHelp(cmd)
			}
			return nil
		},
	}

	if short != nil {
		help.ShortDesc = *short
	}

	if long != nil {
		help.LongDesc = *long
	}

	return help
}

// Handles help function for sub commands of the root command
func subHelp(cmd *rhombifer.Command) {
	fmt.Print("\n")
	fmt.Printf("%s\n", cmd.Name)
	if cmd.LongDesc != "" {
		fmt.Printf("\n%s\n\n", cmd.LongDesc)
	} else {
		fmt.Printf("\n%s\n\n", cmd.ShortDesc)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "%v", text.Bold("Flags:"))
	if cmd.Flags != nil {
		for _, f := range cmd.Flags {
			fmt.Fprintf(w, "\n\t--%s\t%s", f.Name, f.Short)
		}
		fmt.Fprintf(w, "\n")
		w.Flush()
	}
}
