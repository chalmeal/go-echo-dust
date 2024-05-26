package util

import (
	"fmt"
	"strings"
)

type Check struct{}

/**
* SQLエラー判定
* switch文の中で判定しているエラーコードはMySQLに依存します。
* PostgreSQLなど他のDBを利用する場合はDBに応じてエラーコードを書き換えてください。
 */
func (c *Check) SqlErrorCheck(err error) string {
	code := "INTERNAL"
	e := strings.Split(fmt.Sprint(err.Error()), ":")
	switch e[0] {
	// 主キー重複
	case "Error 1062 (23000)":
		code = "PRIMARY"
	}
	return code
}

/**
* TODO:
* パスワード設定時強度判定
*
* 設定するパスワードは以下の条件を満たす必要があります
* 1. 8文字以上の文字列であること
* 2. 大文字小文字それぞれ1文字以上含まれていること
* 3. 数字が1文字以上含まれていること
* 4. 記号が1文字以上含まれていること
*
 */
func (c *Check) SetPasswordCheck() {}
