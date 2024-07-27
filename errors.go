package rhombifer

import "github.com/racg0092/rhombifer/pkg/errs"

var (
	ErroNoRootRunFunc = errs.NewError("root has no Run function. Therefore it can't interpret any flags")

	ErroFlagUndefined   = errs.NewError("flag is undefined.")
	ErroFlagHasNoValues = errs.NewError("flag has no values.")
)
