package utility

import "time"

// 時間を日本時間(GMT+9)に変換する
func ConvertJST(clock time.Time) time.Time {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	clockJST := clock.In(jst)
	return clockJST
}
