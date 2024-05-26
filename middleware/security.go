package middleware

import (
	"echo-example-package/connect"
	"echo-example-package/context"
	"echo-example-package/context/response"
	"echo-example-package/handler"
	"echo-example-package/util"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gocraft/dbr"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	dtx *context.DTx
	u   util.Util
)

/** CORS Config */
func CorsConfig(e *echo.Echo) {

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("EXAMPLE_LOCAL_PORT")},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

}

/** JWT Config */
func JwtConfig(e *echo.Echo) {

	jwt := middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		ParseTokenFunc: func(tokenString string, c echo.Context) (interface{}, error) {
			var token *jwt.Token
			dtx.TxReadOnly(connect.DbConnect(), func(tx *dbr.Tx) response.Result {
				keyFunc := func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(os.Getenv("JWT_SECRET_KEY")), nil
				}

				t, err := jwt.Parse(tokenString, keyFunc)
				if err != nil {
					return response.Error(err, response.ERROR_AUTH_GENERAL)
				}
				a := &context.AccountStore{}
				accountId := t.Claims.(jwt.MapClaims)["account_id"]
				account := a.GetAccountAuth(tx, accountId.(string))
				if !t.Valid || *account.AccessToken != u.Convert.HashToken(t.Raw) {
					return response.Error(errors.New("invalid token error"), response.ERROR_AUTH_INVALID)
				}

				handler.NewHandler()
				token = t

				return response.Success(nil, response.SUCCESS_GENERAL)
			})
			if token == nil {
				return nil, errors.New("verify token error")
			}
			return token, nil
		},
	}

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(jwt))
	h := handler.NewHandler()
	h.Router(api)

}
