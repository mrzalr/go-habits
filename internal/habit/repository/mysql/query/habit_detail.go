package query

func (hb *habitDetail) GetHabitDetailByID() string {
	return `
	SELECT 
		id, habit_id, start_time, end_time, remark, valid
	FROM habit_details
	WHERE id = :id`
}

func (hb *habitDetail) CreateHabitDetail() string {
	return `
	INSERT INTO 
		habit_details(id, habit_id, start_time, end_time, remark, valid)
	VALUES
		(:id, :habitID, :startTime, :endTime, :remark, :valid)`
}

func (hb *habitDetail) GetLastHabitDetailStarted() string {
	return `
	SELECT 
		id, habit_id, start_time, end_time, remark, valid
	FROM habit_details
	WHERE habit_id = :habit_id 
	AND end_time = :end_time
	ORDER BY start_time DESC
	LIMIT 1`
}

func (hb *habitDetail) UpdateHabitDetail() string {
	return `
	UPDATE habit_details
	SET 
		end_time = :endTime,
		remark = :remark,
		valid = :valid
	WHERE id = :id`
}
