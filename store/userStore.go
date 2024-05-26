package store

import (
	"echo-example-package/context/response"
	"echo-example-package/model"
	"echo-example-package/util"

	"github.com/gocraft/dbr"
)

type UserStore struct {
	util util.Util
}

func NewUserStore() *UserStore {
	return &UserStore{
		util: *util.NewUtil(u),
	}
}

func (us *UserStore) find(tx *dbr.Tx) []model.User {
	var u []model.User
	tx.Select("*").From("user").Load(&u)
	return u
}

func (us *UserStore) get(tx *dbr.Tx, id string) model.User {
	var u model.User
	tx.Select("*").From("user").Where("user_id=?", id).Load(&u)
	return u
}

// ユーザ一覧を取得します。
func (us *UserStore) FindUser(tx *dbr.Tx) response.Result {
	var u []model.User
	_, err := tx.Select("*").From("user").Load(&u)
	if err != nil {
		return response.Error(err, response.ERROR_INTERNAL_SERVER)
	}
	return response.Success(u, response.SUCCESS_GENERAL)
}

// ユーザを取得します。
func (us *UserStore) GetUser(tx *dbr.Tx, id string) response.Result {
	var u model.User
	_, err := tx.Select("*").From("user").Where("user_id=?", id).Load(&u)
	if err != nil {
		return response.Error(err, response.ERROR_INTERNAL_SERVER)
	}
	return response.Success(u, response.SUCCESS_GENERAL)
}

// ユーザを登録します。
func (us *UserStore) RegisterUser(tx *dbr.Tx, p *model.RegUser) response.Result {
	reg := us.regUser(p)
	_, err := tx.InsertInto("user").
		Columns("user_id", "name", "name_kana", "regist_date").
		Record(&reg).
		Exec()
	if err != nil {
		return response.Error(err, us.regErrorCheck(err))
	}
	return response.Success(us.get(tx, p.UserId), response.SUCCESS_USER_REGISTER)
}

// ユーザを編集します。
func (us *UserStore) ChangeUser(tx *dbr.Tx, p *model.ChgUser) response.Result {
	chg := us.util.Convert.StructToMap(us.chgUser(p))
	data, err := tx.Update("user").
		SetMap(chg).
		Where("user_id=?", p.UserId).
		Exec()
	if err != nil {
		return response.Error(err, response.ERROR_INTERNAL_SERVER)
	}
	if row, _ := data.RowsAffected(); row == 0 {
		return response.Error(err, response.ERROR_USER_CHANGE_NOT_FOUND)
	}
	return response.Success(us.get(tx, p.UserId), response.SUCCESS_USER_CHANGE)
}

// 登録パラメタ
type regUserRecord struct {
	UserId     string `json:"user_id"`
	Name       string `json:"name"`
	NameKana   string `json:"name_kana"`
	RegistDate string `json:"regist_date"`
}

func (us *UserStore) regUser(p *model.RegUser) regUserRecord {
	return regUserRecord{
		UserId:     p.UserId,
		Name:       p.Name,
		NameKana:   p.NameKana,
		RegistDate: us.util.Time.NowDateTime(),
	}
}

func (us *UserStore) regErrorCheck(err error) response.Result {
	res := response.ERROR_INTERNAL_SERVER
	switch us.util.Check.SqlErrorCheck(err) {
	case "PRIMARY":
		res = response.ERROR_USER_REGIST_PRIMARY
	}
	return res
}

// 変更パラメタ
type chgUserRecord struct {
	Name       string `json:"name"`
	NameKana   string `json:"name_kana"`
	UpdateDate string `json:"update_date"`
}

func (us *UserStore) chgUser(p *model.ChgUser) chgUserRecord {
	return chgUserRecord{
		Name:       p.Name,
		NameKana:   p.NameKana,
		UpdateDate: us.util.Time.NowDateTime(),
	}
}
