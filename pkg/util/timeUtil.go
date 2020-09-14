package util

import (
	"time"
)

const (
	UTC_TIME_LAYOUT    = "2006-01-02 15:04:05 -0700 UTC"
	NORMAL_TIME_LAYOUT = "2006-01-02 15:04:05"
)

func ParseTime(timestr string) (time.Time, error) {
	parsedTime, err := time.Parse(UTC_TIME_LAYOUT, timestr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func GetTimeLocation(location string) (*time.Location, error) {
	return time.LoadLocation(location)
}

func ParseFromInt64(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
