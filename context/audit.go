package context

import (
	"echo-example-package/context/response"
	"echo-example-package/util"

	"github.com/gocraft/dbr"
)

// 利用者監査ログを表現します。
type Audit struct {
	AuditId        *int    `db:"audit_id"`         /** 監査ログID */
	AccountId      *string `db:"account_id"`       /** アカウントID */
	Message        *string `db:"message"`          /** メッセージ */
	ErrorReason    *string `db:"error_reason"`     /** エラー事由 */
	Status         *string `db:"status"`           /** 状態 */
	StartDateTime  *string `db:"start_date_time"`  /** 処理開始日時 */
	FinishDateTime *string `db:"finish_date_time"` /** 処理終了日時 */
}

type AuditStore struct {
	dtx  *DTx
	util util.Util
}

// 利用者監査ログ一覧を取得します。
func (as *AuditStore) FindAudit(tx *dbr.Tx) response.Result {
	var a []Audit
	_, err := tx.Select("*").From("audit").Load(&a)
	if err != nil {
		return response.Error(err, response.ERROR_INTERNAL_SERVER)
	}
	return response.Success(a, response.SUCCESS_GENERAL)
}

// 利用者監査ログを取得します。
func (as *AuditStore) GetAudit(tx *dbr.Tx, id string) response.Result {
	var a Audit
	_, err := tx.Select("*").From("audit").Where("audit_id=?", id).Load(&a)
	if err != nil {
		return response.Error(err, response.ERROR_INTERNAL_SERVER)
	}
	return response.Success(a, response.SUCCESS_GENERAL)
}

/* 利用者監査ログを書き込みます。
* WAL(Write ahead log)でのログ書き込み処理
 */
func (as *AuditStore) WriteAudit(sess *dbr.Session, message string,
	write func() response.Result) (err response.Result) {
	var id int64
	var e error
	as.dtx.Tx(sess, func(tx *dbr.Tx) response.Result {
		id, e = as.aheadWriteAudit(tx, message)
		if e != nil {
			return response.Error(e, response.ERROR_INTERNAL_SERVER)
		}
		return response.Success(id, response.SUCCESS_GENERAL)
	})
	defer func() {
		as.dtx.Tx(sess, func(tx *dbr.Tx) response.Result {
			if as.laterWriteAudit(tx, id, err) != nil {
				return response.Error(e, response.ERROR_INTERNAL_SERVER)
			}
			return response.Success(err.Result, response.SUCCESS_GENERAL)
		})
	}()
	return write()
}

/* 利用者監査ログの先行書き込みを行います。
* store実行前のユースケース開始書き込み処理
 */
func (as *AuditStore) aheadWriteAudit(tx *dbr.Tx, message string) (int64, error) {
	reg := aheadAuditParam{
		AccountId:     "account-id",
		Message:       message,
		Status:        "PROCESSING",
		StartDateTime: as.util.Time.NowDateTimeMsec(),
	}
	audit, e := tx.InsertInto("audit").
		Columns("account_id", "message", "status", "start_date_time").
		Record(&reg).
		Exec()
	tx.Commit()

	id, _ := audit.LastInsertId()

	return id, e
}

/* 利用者監査ログの実行結果書き込みを行います。
* store実行後のユースケース終了書き込み処理
 */
func (as *AuditStore) laterWriteAudit(tx *dbr.Tx, id int64, err response.Result) error {
	var param *laterAuditParam
	if err.Status == 200 {
		param = as.finish(err)
	} else {
		param = as.error(err)
	}

	chg := as.util.Convert.StructToMap(param)
	_, e := tx.Update("audit").SetMap(chg).Where("audit_id=?", id).Exec()
	if e != nil {
		return e
	}
	tx.Commit()

	return nil
}

// 利用者監査ログを完了状態にします。
func (as *AuditStore) finish(err response.Result) *laterAuditParam {
	return &laterAuditParam{
		Status:         "FINISH",
		ErrorReason:    "",
		FinishDateTime: as.util.Time.NowDateTimeMsec(),
	}
}

// 利用者監査ログをエラー状態にします。
func (as *AuditStore) error(err response.Result) *laterAuditParam {
	return &laterAuditParam{
		Status:         "ERROR",
		ErrorReason:    err.ErrorReason,
		FinishDateTime: as.util.Time.NowDateTimeMsec(),
	}
}

// 開始パラメタ
type aheadAuditParam struct {
	AccountId     string `json:"account_id"`
	Message       string `json:"message"`
	Status        string `json:"status"`
	StartDateTime string `json:"start_date_time"`
}

// 終了パラメタ
type laterAuditParam struct {
	Status         string      `json:"status"`
	ErrorReason    interface{} `json:"error_reason"`
	FinishDateTime string      `json:"finish_date_time"`
}
