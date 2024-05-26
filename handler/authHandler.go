package handler

import (
	"echo-example-package/context"
	"echo-example-package/context/response"
	"echo-example-package/usecase"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	auth usecase.AuthService
}

func NewAuthHandler() *authHandler {
	return &authHandler{
		auth: *usecase.NewAuthService(),
	}
}

// ログインします。
func (ah *authHandler) login(c echo.Context) error {
	param := new(context.LoginParam)
	c.Bind(param)
	if result := c.Validate(param); result != nil {
		return c.JSON(response.Validate(result))
	}
	result := ah.auth.Login(param)
	return c.JSON(result.Status, result)
}
