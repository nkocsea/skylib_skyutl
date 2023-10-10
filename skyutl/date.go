package skyutl

import (
	"time"
)

const YYYYMMDD = "20060102"
const YYYYMMDDHHMM = "200601020304"

//function FormatDateTime return "yyyyMMddHHMM"
func FormatDateTime(millisecond int64) string {
	return formatDateTime(millisecond, YYYYMMDDHHMM)
}

//function FormatDate return "yyyyMMdd"
func FormatDate(millisecond int64) string {
	return formatDateTime(millisecond, YYYYMMDD)
}

//function FormatDateWithLayout
func FormatDateWithLayout(millisecond int64, layout string) string {
	return formatDateTime(millisecond, layout)
}

//function StandardFormatDate return "dd/MM/yyyyy"
func StandardFormatDate(millisecond int64) string {
	return formatDateTime(millisecond, "02/01/2006")
}

func formatDateTime(millisecond int64, layout string) string {
	if millisecond == 0 {
		return ""
	}

	return time.Unix(0, millisecond*int64(time.Millisecond)).In(time.UTC).Format(layout)
}


func MakeDateTimeWithDiffHour(millisecond int64, diffHour float64) int64 {
	return time.UnixMilli(millisecond).Add(time.Hour * time.Duration(diffHour)).In(time.UTC).UnixMilli()
}