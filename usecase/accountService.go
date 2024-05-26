package usecase

import (
	"echo-example-package/connect"
	ctx "echo-example-package/context"
	"echo-example-package/context/response"

	"github.com/gocraft/dbr"
)

type AccountService struct {
	sess    *dbr.Session
	dtx     *ctx.DTx
	audit   *ctx.AuditStore
	account *ctx.AccountStore
}

func NewAccountService() *AccountService {
	return &AccountService{
		sess:    connect.DbConnect(),
		audit:   &ctx.AuditStore{},
		account: ctx.NewAccountStore(),
	}
}

// アカウントを登録します。
func (as *AccountService) RegisterAccount(param *ctx.RegisterAccountParam) response.Result {
	return as.audit.WriteAudit(as.sess, "register accountId is: "+param.AccountId,
		func() response.Result {
			return as.dtx.Tx(as.sess, func(tx *dbr.Tx) response.Result {
				return as.account.RegisterAccount(tx, param)
			})
		})
}

// Accountを変更します。
func (as *AccountService) ChangeAccount(param *ctx.ChangeAccountParam) response.Result {
	return as.audit.WriteAudit(as.sess, "change accountId is: "+param.AccountId,
		func() response.Result {
			return as.dtx.Tx(as.sess, func(tx *dbr.Tx) response.Result {
				return as.account.ChangeAccount(tx, param)
			})
		})
}

/**
* パスワードを変更します。
* Account自身によるパスワード変更
 */
func (as *AccountService) ResetPassword(param *ctx.ResetPasswordParam) response.Result {
	return as.dtx.Tx(as.sess, func(tx *dbr.Tx) response.Result {
		return as.account.ResetPassword(tx, "wip-account-id", param.Password) // todo: get account id
	})
}

/**
* パスワードをリセットします。
* 最大権限を持つAccountによるパスワードリセット
 */
func (as *AccountService) ResetPasswordAdmin() response.Result {
	return response.Result{}
}
