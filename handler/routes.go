package handler

import (
	"github.com/labstack/echo/v4"
)

func (a *authHandler) Auth(auth *echo.Group) {

	auth.POST("/login", a.login)

}

func (h *Handler) Router(api *echo.Group) {

	// account
	accounts := api.Group("/account")
	accounts.POST("/register", h.account.registerAccount)
	accounts.POST("/change", h.account.changeAccount)

	// system
	systems := api.Group("/system")
	systems.GET("/audit", h.system.findAudit)
	systems.GET("/audit/:id", h.system.getAudit)

	// user
	users := api.Group("/user")
	users.GET("", h.user.findUser)
	users.GET("/:id", h.user.getUser)
	users.POST("/register", h.user.registerUser)
	users.POST("/change", h.user.changeUser)

}
