package databases

import (
	"database/sql"

	"github.com/google/uuid"
)

type CourseRepository struct {
	db *sql.DB
}

type CourseModel struct {
	ID, Name, Description, CategoryID string
}

func NewCourse(db *sql.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (repo *CourseRepository) Create(name, description, categoryID string) (*CourseModel, error) {
	id := uuid.NewString()

	_, err := repo.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)",
		id, name, description, categoryID,
	)
	if err != nil {
		return nil, err
	}

	return &CourseModel{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}
