package domain

import (
	"fmt"
	"strconv"
	"time"
)

type CustomRequestTime struct {
	time.Time
}

func (t CustomRequestTime) MarshalJSON() ([]byte, error) {
	date := fmt.Sprintf(`%d`, t.Time.Unix())
	return []byte(date), nil
}

func (t *CustomRequestTime) UnmarshalJSON(b []byte) (err error) {
	t.Time, _ = GetTimestampFromUnixString(string(b))
	return
}

type CustomResposeTime struct {
	time.Time
}

func (t CustomResposeTime) MarshalJSON() ([]byte, error) {
	date := fmt.Sprintf(`%d`, t.Time.Unix())
	return []byte(date), nil
}

func (t *CustomResposeTime) UnmarshalJSON(b []byte) (err error) {
	time, err := time.Parse("\"2006-01-02T15:04:05Z07:00\"", string(b))
	if err != nil {
		return err
	}
	t.Time = time
	return
}

func GetTimestampFromUnixString(s string) (time.Time, error) {
	if len(s) == 0 {
		return time.Time{}, ErrBadParamInput
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, ErrBadParamInput
	}
	return time.Unix(i, 0), nil
}
