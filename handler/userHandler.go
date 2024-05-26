package handler

import (
	"echo-example-package/context/response"
	"echo-example-package/model"
	"echo-example-package/usecase"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	user usecase.UserService
}

func NewUserHandler() *userHandler {
	return &userHandler{
		user: *usecase.NewUserService(),
	}
}

// ユーザ一覧を取得します。
func (uh *userHandler) findUser(c echo.Context) error {
	result := uh.user.FindUser()
	return c.JSON(result.Status, result)
}

// ユーザを取得します。
func (uh *userHandler) getUser(c echo.Context) error {
	id := c.Param("id")
	result := uh.user.GetUser(id)
	return c.JSON(result.Status, result)
}

// ユーザを登録します。
func (uh *userHandler) registerUser(c echo.Context) error {
	param := new(model.RegUser)
	c.Bind(param)
	if result := c.Validate(param); result != nil {
		return c.JSON(response.Validate(result))
	}
	result := uh.user.RegisterUser(param)
	return c.JSON(result.Status, result)
}

// ユーザを編集します。
func (uh *userHandler) changeUser(c echo.Context) error {
	param := new(model.ChgUser)
	c.Bind(param)
	if result := c.Validate(param); result != nil {
		return c.JSON(response.Validate(result))
	}
	result := uh.user.ChangeUser(param)
	return c.JSON(result.Status, result)
}
