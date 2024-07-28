package tests

import (
	"fmt"
	"github.com/racg0092/rhombifer"
	"testing"
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

	//todo find one flag test

}
