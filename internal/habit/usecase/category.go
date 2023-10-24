package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/habit/model"
)

func (u *usecase) GetHabitCategories() ([]model.Category, error) {
	categories, err := u.repository.GetHabitCategories()
	if err != nil {
		return []model.Category{}, err
	}

	return categories, nil
}

func (u *usecase) CreateHabitCategory(category model.Category) (model.Category, error) {
	insertedID, err := u.repository.CreateHabitCategory(category)
	if err != nil {
		return model.Category{}, err
	}

	return u.repository.GetHabitCategoryByID(insertedID)
}

func (u *usecase) UpdateHabitCategory(id uuid.UUID, category model.Category) (model.Category, error) {
	foundCatagory, err := u.repository.GetHabitCategoryByID(id)
	if err != nil {
		return model.Category{}, err
	}

	category.ID = foundCatagory.ID
	category.UpdatedAt = time.Now()

	updatedID, err := u.repository.UpdateHabitCategory(category)
	if err != nil {
		return model.Category{}, err
	}

	return u.repository.GetHabitCategoryByID(updatedID)
}
