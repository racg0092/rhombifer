package parsing

import (
	"os"

	"github.com/racg0092/rhombifer/utils"
)

func InputValidation() error {

	args := os.Args[1:]

	if len(args) == 0 {
		return utils.RhombiError{
			Message: "No arguments found",
			Code:    utils.NO_ARGS,
		}
	}

	return nil
}
