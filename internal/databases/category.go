package databases

import (
	"database/sql"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

type CategoryModel struct {
	ID, Name, Description string
}

func NewCategory(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Create(name, description string) (*CategoryModel, error) {
	id := uuid.NewString()

	_, err := repo.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, description,
	)
	if err != nil {
		return nil, err
	}

	return &CategoryModel{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}
