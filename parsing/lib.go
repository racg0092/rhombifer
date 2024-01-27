package parsing

import (
	"os"

	"github.com/racg0092/rhombifer/utils"
)

//todo: Leave the user decide weather to implement of just print  available commands if no args

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
