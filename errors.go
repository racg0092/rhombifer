package rhombifer

import "github.com/racg0092/rhombifer/pkg/errs"

var (
	ErroNoRootRunFunc = errs.NewError("Root has no Run function. Therefore it can't interpret any flags")
)
