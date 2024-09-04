package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/waldrey/go-expert/sqlc/internal/db"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %v", errRb, err)
		}
		return err
	}

	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			Price:       argsCourse.Price,
			CategoryID:  argsCategory.ID,
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	courses, err := queries.ListCourses(ctx)
	if err != nil {
		log.Fatalf("error on list courses: %v", err)
	}

	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f\n", course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}

	/*
		courseArgs := CourseParams{
			ID:          uuid.New().String(),
			Name:        "Go",
			Description: sql.NullString{String: "Go programming language", Valid: true},
			Price:       230.99,
		}

		categoryArgs := CategoryParams{
			ID:          uuid.New().String(),
			Name:        "Programming",
			Description: sql.NullString{String: "Programming category", Valid: true},
		}

		courseDB := NewCourseDB(dbConn)
		err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
		if err != nil {
			log.Fatalf("error on create course and category: %v", err)
		}
	*/
}
