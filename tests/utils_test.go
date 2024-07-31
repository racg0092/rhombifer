package tests

import (
	"fmt"
	"testing"

	"github.com/racg0092/rhombifer"
)

func TestFindFlags(t *testing.T) {

	t.Run("find flags", func(t *testing.T) {
		root := rhombifer.Root()
		cmd := rhombifer.Command{
			Name: "cmd",
			Run: func(args ...string) error {
				fmt.Println("Hello World form cmd")
				return nil
			},
		}
		addSampleFlags(&cmd)
		root.AddSub(&cmd)
		rhombifer.ExecCommand(cmd.Name, "--recursive", "--foo")
		flags, err := rhombifer.FindFlags("--recursive", "--foo")
		if err != nil {
			t.Error(err)
		}

		if len(flags) < 2 {
			t.Errorf("did not found all flags")
		}

		if flags[0].Name != "recursive" && flags[1].Name != "foo" {
			t.Errorf("flags found in wrong order")
		}

	})

	t.Run("find one flag", func(t *testing.T) {
		root := rhombifer.Root()
		cmd := rhombifer.Command{
			Name: "cmd",
			Run: func(args ...string) error {
				fmt.Println("Hello world from cmd")
				return nil
			},
		}
		addSampleFlags(&cmd)
		root.AddSub(&cmd)
		rhombifer.ExecCommand(cmd.Name, "--recursive", "--foo")
		flag, err := rhombifer.FindFlag("--recursive")
		if err != nil {
			t.Error(err)
		}

		if flag.Name != "recursive" {
			t.Errorf("found wrong flag %s", flag.Name)
		}
	})
}
