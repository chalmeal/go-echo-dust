package model

// ユーザを表現します。
type User struct {
	UserId     *string `db:"user_id"`     /** ユーザーID */
	Name       *string `db:"name"`        /** 氏名 */
	NameKana   *string `db:"name_kana"`   /** 氏名カナ */
	RegistDate *string `db:"regist_date"` /** 登録日時 */
	UpdateDate *string `db:"update_date"` /** 更新日時 */
}

// ユーザを登録します。
type RegUser struct {
	UserId   string `json:"user_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	NameKana string `json:"name_kana" validate:"required"`
}

// ユーザを編集します。
type ChgUser struct {
	UserId   string `json:"user_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	NameKana string `json:"name_kana" validate:"required"`
}
