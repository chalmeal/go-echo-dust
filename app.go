package main

import (
	"echo-example-package/handler"
	mw "echo-example-package/middleware"
	"echo-example-package/validator"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load("env/.env")

	e := echo.New()
	e.Validator = validator.NewValidator()
	e.Use(middleware.Logger())
	mw.CorsConfig(e)

	a := handler.NewAuthHandler()
	auth := e.Group("")
	a.Auth(auth)

	mw.JwtConfig(e)

	e.Logger.Fatal(e.Start(os.Getenv("APP_PORT")))
}
