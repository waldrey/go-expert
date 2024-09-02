package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryId string) (Course, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES (?, ?, ?, ?)", id, name, description, categoryId)
	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description, CategoryId: categoryId}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryId); err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func (c *Course) FindByCategoryID(categoryId string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = ?", categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryId); err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}
