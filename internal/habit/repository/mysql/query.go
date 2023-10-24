package mysql

var (

	// Habit category query

	GetHabitCategoriesQuery = `
	SELECT id, name, created_at, updated_at
	FROM habit_categories`

	GetHabitCategoryByIDQuery = `
	SELECT id, name, created_at, updated_at
	FROM habit_categories
	WHERE id = :id`

	CreateHabitCategoryQuery = `
	INSERT INTO habit_categories(id, name, created_at, updated_at)
	VALUES (:id, :name, :createdAt, :updatedAt)`

	UpdateHabitCategoryQuery = `
	UPDATE habit_categories
	SET
		name = :name,
		updated_at = :updatedAt
	WHERE id = :id`

	// Habit query

	GetAllHabitsQuery = `
	SELECT 
		id, category_id, activity, description, created_at 
	FROM habits
	WHERE created_at BETWEEN :startDate AND :endDate`

	GetHabitByIDQuery = `
	SELECT 
		id, category_id, activity, description, created_at 
	FROM habits
	WHERE id = :id`

	CreateHabitQuery = `
	INSERT INTO 
		habits(id, category_id, activity, description, created_at)
	VALUES
		(:id, :categoryID, :activity, :description, :createdAt)`

	// Habit detail query

	GetHabitDetailByIDQuery = `
	SELECT 
		id, habit_id, start_time, end_time, remark, valid
	FROM habit_details
	WHERE id = :id`

	CreateHabitDetailQuery = `
	INSERT INTO 
		habit_details(id, habit_id, start_time, end_time, remark, valid)
	VALUES
		(:id, :habitID, :startTime, :endTime, :remark, :valid)`

	GetLastHabitDetailStartedQuery = `
	SELECT 
		id, habit_id, start_time, end_time, remark, valid
	FROM habit_details
	WHERE habit_id = :habit_id 
	AND end_time = :end_time
	ORDER BY start_time DESC
	LIMIT 1`

	UpdateHabitDetailQuery = `
	UPDATE habit_details
	SET 
		end_time = :endTime,
		remark = :remark,
		valid = :valid
	WHERE id = :id`
)
