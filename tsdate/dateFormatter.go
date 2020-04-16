package tsdate

import (
	"time"
)

var dateTimeTZFormat = "2006-01-02T15:04:05Z"
var dateFormat = "2006-01-02"

const timeMilli = int64(time.Millisecond)

func DateTimeTZFormatter(dateTime *time.Time) string {
	return dateTime.UTC().Format(dateTimeTZFormat)
}

func DateFormatter(date *time.Time) string {
	return date.UTC().Format(dateFormat)
}

func GetDateTimeTZFormat() *string {
	return &dateTimeTZFormat
}

func GetDateFormat() *string {
	return &dateFormat
}

func GetTimeInMilli(timeValue time.Time) int64 {
	return int64(timeValue.UnixNano() / timeMilli)
}

func GetDateAndTime(date string) (formattedDate time.Time, err error) {
	formattedDate, err = time.Parse(dateTimeTZFormat, date)
	return
}
