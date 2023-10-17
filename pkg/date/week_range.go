package date

import (
	"time"
)

func GetWeekRange(weekday int) WeekRange {
	dayDuration := time.Hour * 24

	// Here I use standard, week starts Monday and ends Sunday.
	start := time.Now().Add(-dayDuration * time.Duration(weekday-1))
	end := time.Now().Add(dayDuration * time.Duration(7-weekday))
	return WeekRange{start, end}
}
