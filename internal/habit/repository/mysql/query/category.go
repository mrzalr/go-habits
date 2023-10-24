package query

func (c *category) GetHabitCategories() string {
	return `
	SELECT id, name, created_at, updated_at
	FROM habit_categories`
}

func (c *category) GetHabitCategoryByID() string {
	return `
	SELECT id, name, created_at, updated_at
	FROM habit_categories
	WHERE id = :id`
}

func (c *category) CreateHabitCategory() string {
	return `
	INSERT INTO habit_categories(id, name, created_at, updated_at)
	VALUES (:id, :name, :createdAt, :updatedAt)`
}

func (c *category) UpdateHabitCategory() string {
	return `
	UPDATE habit_categories
	SET
		name = :name,
		updated_at = :updatedAt
	WHERE id = :id`
}
