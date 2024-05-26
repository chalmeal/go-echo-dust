package context

import (
	"echo-example-package/util"
)

var (
	u util.Util
)

type Context struct {
	Account AccountStore
}

func NewContext() *Context {
	return &Context{
		Account: *NewAccountStore(),
	}
}
