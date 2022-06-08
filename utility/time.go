package utility

import (
	"html/template"
	"time"
)

// 時間を日本時間(GMT+9)に変換する
func ConvertJST(clock time.Time) time.Time {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	clockJST := clock.In(jst)
	return clockJST
}

// 時間を整形してテンプレートに返す
func TimeFormat(clock time.Time) template.HTML {
	return template.HTML(clock.Format(time.ANSIC))
}
