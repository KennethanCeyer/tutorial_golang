package main

import (
	"time"
)

func GetDateTime(date string) (time.Time, error) {
	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}
	return dateTime, nil
}

func GetTimeUnit(time int) string {
	if time < 6 {
		return "새벽"
	} else if time < 10 {
		return "아침"
	} else if time < 13 {
		return "점심"
	} else if time < 17 {
		return "낮"
	} else if time < 21 {
		return "저녁"
	}
	return "심야"
}
