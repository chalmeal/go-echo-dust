package usecase

import (
	"echo-example-package/connect"
	"echo-example-package/context"
	"echo-example-package/context/response"
	"echo-example-package/model"
	"echo-example-package/store"

	"github.com/gocraft/dbr"
)

type UserService struct {
	sess  *dbr.Session
	dtx   *context.DTx
	audit *context.AuditStore
	store *store.Store
}

func NewUserService() *UserService {
	return &UserService{
		sess:  connect.DbConnect(),
		audit: &context.AuditStore{},
		store: store.NewStore(),
	}
}

// ユーザ一覧を取得します。
func (us *UserService) FindUser() response.Result {
	return us.dtx.TxReadOnly(us.sess, func(tx *dbr.Tx) response.Result {
		return us.store.User.FindUser(tx)
	})
}

// ユーザを取得します。
func (us *UserService) GetUser(id string) response.Result {
	return us.dtx.TxReadOnly(us.sess, func(tx *dbr.Tx) response.Result {
		return us.store.User.GetUser(tx, id)
	})
}

// ユーザを登録します。
func (us *UserService) RegisterUser(param *model.RegUser) response.Result {
	return us.audit.WriteAudit(us.sess, "register userId is: "+param.UserId,
		func() response.Result {
			return us.dtx.Tx(us.sess, func(tx *dbr.Tx) response.Result {
				return us.store.User.RegisterUser(tx, param)
			})
		})
}

// ユーザを編集します。
func (us *UserService) ChangeUser(param *model.ChgUser) response.Result {
	return us.audit.WriteAudit(us.sess, "update userId is: "+param.UserId,
		func() response.Result {
			return us.dtx.Tx(us.sess, func(tx *dbr.Tx) response.Result {
				return us.store.User.ChangeUser(tx, param)
			})
		})
}
