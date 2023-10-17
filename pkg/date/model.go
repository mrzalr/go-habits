package date

import "time"

type WeekRange struct {
	StartDate time.Time `db:"startDate"`
	EndDate   time.Time `db:"endDate"`
}
