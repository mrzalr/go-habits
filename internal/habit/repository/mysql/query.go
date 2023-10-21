package mysql

var (
	GetAllHabitsQuery = `
	SELECT 
		id, activity, description, start_time, end_time, created_at 
	FROM habit
	WHERE created_at BETWEEN :startDate AND :endDate`

	GetHabitByIDQuery = `
	SELECT 
		id, activity, description, start_time, end_time, created_at 
	FROM habit
	WHERE id = :id`

	CreateHabitQuery = `
	INSERT INTO 
		habit(id, activity, description, start_time, end_time, created_at)
	VALUES
		(:id, :activity, :description, :startTime, :endTime, :createdAt)`

	UpdateHabitQuery = `
	UPDATE habit 
	SET 
		activity = :activity,
		description = :description,
		start_time = :startTime,
		end_time = :endTime
	WHERE id = :id`
)
