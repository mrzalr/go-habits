package query

type (
	category    struct{}
	habit       struct{}
	habitDetail struct{}
)

var (
	Category    = &category{}
	Habit       = &habit{}
	HabitDetail = &habitDetail{}
)
