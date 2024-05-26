package handler

import (
	"echo-example-package/usecase"

	"github.com/labstack/echo/v4"
)

type systemHandler struct {
	system usecase.SystemService
}

func NewSystemHandler() *systemHandler {
	return &systemHandler{
		system: *usecase.NewSystemService(),
	}
}

// 利用者監査ログ一覧を取得します。
func (sh *systemHandler) findAudit(c echo.Context) error {
	result := sh.system.FindAudit()
	return c.JSON(result.Status, result)
}

// 利用者監査ログを取得します。
func (sh *systemHandler) getAudit(c echo.Context) error {
	id := c.Param("id")
	result := sh.system.GetUser(id)
	return c.JSON(result.Status, result)
}
