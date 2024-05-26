package usecase

import (
	"echo-example-package/connect"
	"echo-example-package/context"
	"echo-example-package/context/response"

	"github.com/gocraft/dbr"
)

type SystemService struct {
	sess  *dbr.Session
	dtx   *context.DTx
	audit *context.AuditStore
}

func NewSystemService() *SystemService {
	return &SystemService{
		sess:  connect.DbConnect(),
		audit: &context.AuditStore{},
	}
}

// 利用者監査ログ一覧を取得します。
func (ss *SystemService) FindAudit() response.Result {
	return ss.dtx.TxReadOnly(ss.sess, func(tx *dbr.Tx) response.Result {
		err := ss.audit.FindAudit(tx)
		return err
	})
}

// 利用者監査ログを取得します。
func (ss *SystemService) GetUser(id string) response.Result {
	return ss.dtx.TxReadOnly(ss.sess, func(tx *dbr.Tx) response.Result {
		err := ss.audit.GetAudit(tx, id)
		return err
	})
}
