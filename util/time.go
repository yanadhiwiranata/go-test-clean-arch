package util

import (
	"strconv"
	"time"

	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
)

func SameDay(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func GenerateSampleTestTime() (time.Time, time.Time, time.Time, time.Time) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)
	the_day_after_tomorrow := now.AddDate(0, 0, 2)
	return yesterday, now, tomorrow, the_day_after_tomorrow
}

func GetTimestampFromUnixString(s string) (time.Time, error) {
	if len(s) == 0 {
		return time.Time{}, domain.ErrBadParamInput
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, domain.ErrBadParamInput
	}
	return time.Unix(i, 0), nil
}
