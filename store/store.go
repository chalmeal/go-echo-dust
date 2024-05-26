package store

import (
	"echo-example-package/util"
)

var (
	u util.Util
)

type Store struct {
	User UserStore
}

func NewStore() *Store {
	return &Store{
		User: *NewUserStore(),
	}
}
