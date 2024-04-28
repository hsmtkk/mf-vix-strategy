package weekstofirstmonth

import "time"

// 第1限月満期までの残り週数
func CalculateWeeksToFirstMonth(firstMonth time.Time) int {
	restDays := int((time.Until(firstMonth)).Hours()) / 24
	weeks := restDays / 7
	return weeks
}
