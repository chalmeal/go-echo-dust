/*
* フォーマット変換による共通処理を表現します。
* bool及びerrorによるハンドリングは考慮していません。
 */
package util

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Time struct{}

/*
* 現在日付をString形式で返します。
* YYYY-MM-DD
 */
func (t *Time) NowDate() string {
	l, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err)
		return ""
	}
	jst := time.Now().UTC()
	now := strings.Split(fmt.Sprint(jst.In(l)), " ")
	return now[0]
}

/*
* 現在日時をString形式で返します。
* YYYY-MM-DD hh:mm:ss
 */
func (t *Time) NowDateTime() string {
	l, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err)
		return ""
	}
	jst := time.Now().UTC()
	now := strings.Split(fmt.Sprint(jst.In(l)), ".")
	return now[0]
}

/*
* 現在日時をString形式で返します。
* YY-MM-DD hh:mm:ss.mm
 */
func (t *Time) NowDateTimeMsec() string {
	l, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err)
		return ""
	}
	jst := time.Now().UTC()
	now := strings.Split(fmt.Sprint(jst.In(l)), ".")
	return now[0] + "." + now[1][:2]
}
