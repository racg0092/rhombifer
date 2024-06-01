package rhombifer

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	rhombitext "github.com/racg0092/rhombifer/pkg/text"
)

func UseBuiltInHelp(short, long *string) Command {
	help := Command{
		Name:      "help",
		ShortDesc: "Displays help information",
		LongDesc: `
		Displays help information for the specified command or the root command if no command is specified.
		`,
		Leaf: true,
		Run: func(args []string) error {

			if len(args) == 0 {
				fmt.Printf("\n%s", rhombitext.LightGray(strings.ToUpper(root.Name)))
				if root.LongDesc != "" {
					fmt.Printf("%s", root.LongDesc)
				} else {
					fmt.Printf("%s", root.ShortDesc)
				}
				w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
				fmt.Fprintf(w, "\n%v", rhombitext.Bold("Commands"))
				for _, sub := range root.Subs {
					fmt.Fprintf(w, "\n\t%s\t%s", sub.Name, sub.ShortDesc)
				}
				fmt.Fprintf(w, "\n\n")
				w.Flush()
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
