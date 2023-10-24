package date

import (
	"time"
)

func GetWeekRange(weekday int) WeekRange {
	dayDuration := time.Hour * 24

	// The standard week is starts from Monday 00:00 until Sunday 24:00
	start := time.Now().Add(-dayDuration * time.Duration(weekday-1))
	end := time.Now().Add(dayDuration * time.Duration(8-weekday))
	return WeekRange{resetTime(start), resetTime(end)}
}

func resetTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
