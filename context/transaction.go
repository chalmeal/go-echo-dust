/*
* トランザクションラッパー処理
 */
package context

import (
	"echo-example-package/context/response"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

type DTx struct {
}

/*
* Read-Only Transaction Wrapper
 */
func (dtx *DTx) TxReadOnly(sess *dbr.Session,
	txFunc func(*dbr.Tx) response.Result) (err response.Result) {
	tx, e := sess.Begin()
	if e != nil {
		return response.Error(e, response.ERROR_INTERNAL_SERVER)
	}
	defer func() {
		if p := recover(); p != nil {
			return
		} else if e != nil {
			return
		} else {
			e = tx.Commit()
		}
	}()
	return txFunc(tx)
}

/*
* Transaction Wrapper
 */
func (dtx *DTx) Tx(sess *dbr.Session,
	txFunc func(*dbr.Tx) response.Result) (err response.Result) {
	tx, e := sess.Begin()
	if e != nil {
		return response.Error(e, response.ERROR_INTERNAL_SERVER)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			return
		} else if err.Status != 200 {
			tx.Rollback()
			return
		} else {
			e = tx.Commit()
		}
	}()
	return txFunc(tx)
}
