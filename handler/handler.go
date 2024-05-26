package handler

import (
	"echo-example-package/context"

	"github.com/labstack/echo/v4"
)

var (
	c echo.Context
)

type Handler struct {
	ctx     context.Context
	auth    authHandler
	account accountHandler
	system  systemHandler
	user    userHandler
}

func NewHandler() *Handler {
	return &Handler{
		ctx:     *context.NewContext(),
		auth:    *NewAuthHandler(),
		account: *NewAccountHandler(),
		system:  *NewSystemHandler(),
		user:    *NewUserHandler(),
	}
}
