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

func (repo *CourseRepository) FindAll() ([]CourseModel, error) {
	rows, err := repo.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []CourseModel{}
	for rows.Next() {
		var id, name, description, categoryID string

		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}

		categories = append(categories, CourseModel{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}

	return categories, nil
}
