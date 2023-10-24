package mysql

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/habit/model"
)

func (r *repository) GetHabitCategories() ([]model.Category, error) {
	stmt, err := r.db.Prepare(GetHabitCategoriesQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrDataNotFound
		}
		return nil, err
	}
	defer rows.Close()

	categories := []model.Category{}
	for rows.Next() {
		category := model.Category{}
		err := rows.Scan(
			&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *repository) GetHabitCategoryByID(id uuid.UUID) (model.Category, error) {
	nstmt, err := r.db.PrepareNamed(GetHabitCategoryByIDQuery)
	if err != nil {
		return model.Category{}, err
	}
	defer nstmt.Close()

	params := queryParams{"id": id}

	category := model.Category{}
	err = nstmt.QueryRow(params).Scan(
		&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Category{}, model.ErrDataNotFound
		}
		return model.Category{}, err
	}

	return category, nil
}

func (r *repository) CreateHabitCategory(category model.Category) (uuid.UUID, error) {
	nstmt, err := r.db.PrepareNamed(CreateHabitCategoryQuery)
	if err != nil {
		return uuid.UUID{}, err
	}
	defer nstmt.Close()

	_, err = nstmt.Exec(category)
	if err != nil {
		return uuid.UUID{}, err
	}

	return category.ID, nil
}

func (r *repository) UpdateHabitCategory(category model.Category) (uuid.UUID, error) {
	nstmt, err := r.db.PrepareNamed(UpdateHabitCategoryQuery)
	if err != nil {
		return uuid.UUID{}, err
	}
	defer nstmt.Close()

	_, err = nstmt.Exec(category)
	if err != nil {
		return uuid.UUID{}, err
	}

	return category.ID, nil
}
