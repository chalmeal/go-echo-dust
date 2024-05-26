package usecase

import (
	"echo-example-package/connect"
	"echo-example-package/context"
	"echo-example-package/context/response"

	"github.com/gocraft/dbr"
)

type AuthService struct {
	sess  *dbr.Session
	dtx   *context.DTx
	audit *context.AuditStore
	auth  *context.AccountStore
}

func NewAuthService() *AuthService {
	return &AuthService{
		sess:  connect.DbConnect(),
		audit: &context.AuditStore{},
		auth:  &context.AccountStore{},
	}
}

// ログインします。
func (as *AuthService) Login(param *context.LoginParam) response.Result {
	return as.audit.WriteAudit(as.sess, "login accountId is: "+param.AccountId,
		func() response.Result {
			return as.dtx.Tx(as.sess, func(tx *dbr.Tx) response.Result {
				return as.auth.Login(tx, param)
			})
		})
}
