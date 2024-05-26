package handler

import (
	"echo-example-package/context"
	"echo-example-package/context/response"
	"echo-example-package/usecase"

	"github.com/labstack/echo/v4"
)

type accountHandler struct {
	account usecase.AccountService
}

func NewAccountHandler() *accountHandler {
	return &accountHandler{
		account: *usecase.NewAccountService(),
	}
}

// アカウントを登録します。
func (ah *accountHandler) registerAccount(c echo.Context) error {
	param := new(context.RegisterAccountParam)
	c.Bind(param)
	if result := c.Validate(param); result != nil {
		return c.JSON(response.Validate(result))
	}
	result := ah.account.RegisterAccount(param)
	return c.JSON(result.Status, result)
}

// アカウントを変更します。
func (ah *accountHandler) changeAccount(c echo.Context) error {
	param := new(context.ChangeAccountParam)
	c.Bind(param)
	if result := c.Validate(param); result != nil {
		return c.JSON(response.Validate(result))
	}
	result := ah.account.ChangeAccount(param)
	return c.JSON(result.Status, result)
}
