package query

func (h *habit) GetAllHabits() string {
	return `
	SELECT 
		h.id, c.name category, h.activity, h.description, h.created_at, h.updated_at
	FROM habits h
	INNER JOIN habit_categories c ON h.category_id = c.id
	WHERE h.created_at BETWEEN :startDate AND :endDate`
}

func (h *habit) GetHabitByID() string {
	return `
	SELECT 
		h.id, c.name category, h.activity, h.description, h.created_at, h.updated_at 
	FROM habits h
	INNER JOIN habit_categories c ON h.category_id = c.id
	WHERE c.id = :id`
}

func (h *habit) CreateHabit() string {
	return `
	INSERT INTO 
		habits(id, category_id, activity, description, created_at, updated_at)
	VALUES
		(:id, :categoryID, :activity, :description, :createdAt, :updatedAt)`
}

func (h *habit) UpdateHabit() string {
	return `
	UPDATE habits
	SET
		category_id = :categoryID,
		activity = :activity,
		description = :description,
		updated_at = :updatedAt
	WHERE id = :id`
}
