package context

import (
	"echo-example-package/context/response"
	"echo-example-package/util"
	"errors"
	"os"
	"time"

	"github.com/gocraft/dbr"
	"github.com/golang-jwt/jwt/v5"
)

// アカウントを表現します。
type Account struct {
	AccountId          string  `db:"account_id"`            /** アカウントID */
	AccountName        *string `db:"account_name"`          /** アカウント名 */
	Email              *string `db:"email"`                 /** メールアドレス */
	Authority          *string `db:"authority"`             /** 権限 */
	Salt               string  `db:"salt"`                  /** ソルト */
	Password           string  `db:"password"`              /** パスワード(ハッシュ済) */
	LastUpdatePassDate *string `db:"last_update_pass_date"` /** パスワード最終更新日 */
	AccessToken        *string `db:"access_token"`          /** アクセストークン(ハッシュ済) */
	RegistDate         *string `db:"regist_date"`           /** 登録日時 */
	UpdateDate         *string `db:"update_date"`           /** 更新日時 */
}

type AccountStore struct {
	util util.Util
}

func NewAccountStore() *AccountStore {
	return &AccountStore{
		util: *util.NewUtil(u),
	}
}

type LoginParam struct {
	AccountId string `json:"account_id" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type RegisterAccountParam struct {
	AccountId   string `json:"account_id" validate:"required"`
	AccountName string `json:"account_name" validate:"required"`
	Email       string `json:"email" validate:"email"`
	Password    string `json:"password" validate:"required"`
}

type ChangeAccountParam struct {
	AccountId   string `json:"account_id" validate:"required"`
	AccountName string `json:"account_name"`
	Email       string `json:"email"`
	Authority   string `json:"authority"`
}

type ResetPasswordParam struct {
	Password string `json:"password" validate:"required"`
}

func (as *AccountStore) GetAccountInfo(tx *dbr.Tx, id string) Account {
	var a Account
	tx.Select("account_id, account_name, email, authority").
		From("account").
		Where("account_id=?", id).
		Load(&a)
	return a
}

func (as *AccountStore) GetAccountAuth(tx *dbr.Tx, id string) Account {
	var a Account
	tx.Select("account_id, salt, password, last_update_pass_date, access_token").
		From("account").
		Where("account_id=?", id).
		Load(&a)
	return a
}

// ログインします。
func (as *AccountStore) Login(tx *dbr.Tx, p *LoginParam) response.Result {
	account := as.GetAccountAuth(tx, p.AccountId)
	if p.AccountId != account.AccountId || as.util.Convert.HashPW(account.Salt, p.Password) != account.Password {
		return response.Error(errors.New("login error."), response.ERROR_AUTH_LOGIN)
	}
	claims := &ClaimsRecord{
		account.AccountId,
		account.Password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return response.Error(err, response.ERROR_INTERNAL_SERVER)
	}

	as.updateAccessToken(tx, account.AccountId, t)

	return response.Success(t, response.SUCCESS_LOGIN)
}

// アクセストークンを更新します。
func (as *AccountStore) updateAccessToken(tx *dbr.Tx, id string, token string) error {
	t := map[string]interface{}{"access_token": as.util.Convert.HashToken(token)}
	_, err := tx.Update("account").
		SetMap(t).
		Where("account_id=?", id).
		Exec()
	if err != nil {
		return err
	}
	return nil
}

/**
* Accountを登録します。
* 最大権限を持つAccountのみ登録が可能
 */
func (as *AccountStore) RegisterAccount(tx *dbr.Tx, p *RegisterAccountParam) response.Result {
	reg := as.regAccount(p)
	_, err := tx.InsertInto("account").
		Columns("account_id", "account_name", "email", "authority", "salt", "password", "last_update_pass_date", "regist_date").
		Record(&reg).
		Exec()
	if err != nil {
		return response.Error(err, response.ERROR_INTERNAL_SERVER)
	}
	return response.Success(as.GetAccountInfo(tx, p.AccountId), response.SUCCESS_ACCOUNT_REGISTER)
}

// Accountを更新します。
func (as *AccountStore) ChangeAccount(tx *dbr.Tx, p *ChangeAccountParam) response.Result {
	chg := as.util.Convert.StructToMap(as.chgAccount(p))
	data, err := tx.Update("account").
		SetMap(chg).
		Where("account_id=?", p.AccountId).
		Exec()
	if err != nil {
		return response.Error(err, response.ERROR_INTERNAL_SERVER)
	}
	if row, _ := data.RowsAffected(); row == 0 {
		return response.Error(errors.New("request account not found."), response.ERROR_REQUEST_ACCOUNT_NOT_FOUND)
	}
	return response.Success(as.GetAccountInfo(tx, p.AccountId), response.SUCCESS_ACCOUNT_CHANGE)
}

/**
* パスワードを変更します。
* Account自身によるパスワード変更
 */
func (as *AccountStore) ResetPassword(tx *dbr.Tx, id string, pw string) response.Result {
	account := as.GetAccountAuth(tx, id)
	if id != account.AccountId || as.util.Convert.HashPW(account.Salt, pw) != account.Password {
		return response.Error(errors.New("login error"), response.ERROR_AUTH_LOGIN)
	}
	chg := as.util.Convert.StructToMap(as.chgPassword(pw))
	_, err := tx.Update("account").
		SetMap(chg).
		Where("account_id=?", id).
		Exec()
	if err != nil {
		return response.Error(err, response.ERROR_INTERNAL_SERVER)
	}
	return response.Success(nil, response.SUCCESS_USER_REGISTER)
}

// クレームパラメタ
type ClaimsRecord struct {
	AccountId string `json:"account_id" validate:"required"`
	Password  string `json:"password" validate:"required"`
	jwt.RegisteredClaims
}

// アカウント登録パラメタ
type regAccountRecord struct {
	AccountId          string `json:"account_id"`
	AccountName        string `json:"account_name"`
	Email              string `json:"email"`
	Authority          string `json:"authority"`
	Salt               string `json:"salt"`
	Password           string `json:"password"`
	LastUpdatePassDate string `json:"last_update_pass_date"`
	RegistDate         string `json:"regist_date"`
}

func (as *AccountStore) regAccount(p *RegisterAccountParam) regAccountRecord {
	salt := as.util.Id.Uuid()
	return regAccountRecord{
		AccountId:          p.AccountId,
		AccountName:        p.AccountName,
		Email:              p.Email,
		Authority:          "NORMAL",
		Salt:               salt,
		Password:           as.util.Convert.HashPW(salt, p.Password),
		LastUpdatePassDate: as.util.Time.NowDate(),
		RegistDate:         as.util.Time.NowDateTime(),
	}
}

// アカウント変更パラメタ
type chgAccountRecord struct {
	AccountId   string `json:"account_id"`
	AccountName string `json:"account_name"`
	Email       string `json:"email"`
	Authority   string `json:"authority"`
	UpdateDate  string `json:"update_date"`
}

func (as *AccountStore) chgAccount(p *ChangeAccountParam) chgAccountRecord {
	return chgAccountRecord{
		AccountId:   p.AccountId,
		AccountName: p.AccountName,
		Email:       p.Email,
		Authority:   p.Authority,
		UpdateDate:  as.util.Time.NowDateTime(),
	}
}

// パスワード変更パラメタ
type chgPwRecord struct {
	Salt               string `json:"salt"`
	Password           string `json:"password"`
	LastUpdatePassDate string `json:"last_update_pass_date"`
	UpdateDate         string `json:"update_date"`
}

func (as *AccountStore) chgPassword(pw string) chgPwRecord {
	salt := as.util.Id.Uuid()
	return chgPwRecord{
		Salt:               salt,
		Password:           as.util.Convert.HashPW(salt, pw),
		LastUpdatePassDate: as.util.Time.NowDateTime(),
		UpdateDate:         as.util.Time.NowDateTime(),
	}
}
