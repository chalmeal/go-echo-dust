package connect

import (
	"echo-example-package/env"
	"log"

	"github.com/gocraft/dbr"
)

func DbConnect() *dbr.Session {
	d, i := env.DbInfo()
	db, err := dbr.Open(d, i, nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	sess := db.NewSession(nil)
	return sess
}
