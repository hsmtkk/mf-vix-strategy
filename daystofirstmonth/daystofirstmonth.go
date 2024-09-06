package daystofirstmonth

import "time"

// 第1限月満期までの残り週数
func CalculateDaysToFirstMonth(firstMonth time.Time) int {
	restDays := int((time.Until(firstMonth)).Hours()) / 24
	return restDays
}
