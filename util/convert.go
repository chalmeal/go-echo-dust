package util

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
)

type Convert struct{}

/*
* パラメタコンバート
* structで定義したパラメタをmap[string]interface{}型に変換
 */
func (c *Convert) StructToMap(param interface{}) map[string]interface{} {
	var data map[string]interface{}
	jsonStr, err := json.Marshal(param)
	if err != nil {
		log.Println(err)
		return nil
	}

	out := new(bytes.Buffer)
	err = json.Indent(out, jsonStr, "", "    ")
	if err != nil {
		log.Println(err)
		return nil
	}
	if err := json.Unmarshal([]byte(out.String()), &data); err != nil {
		log.Println(err)
		return nil
	}

	return data
}

/**
* パスワードハッシュ化(SHA-256)
 */
func (c *Convert) HashPW(salt string, password string) string {
	pw := sha256.Sum256([]byte(salt + password))
	return fmt.Sprintf("%x", pw)
}

/**
* アクセストークンハッシュ化(SHA-256)
 */
func (c *Convert) HashToken(token string) string {
	t := sha256.Sum256([]byte(token))
	return fmt.Sprintf("%x", t)
}
