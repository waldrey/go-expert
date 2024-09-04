package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/waldrey/go-expert/sqlc/internal/db"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	/*
		err = queries.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          uuid.New().String(),
			Name:        "Backend",
			Description: sql.NullString{String: "Backend description", Valid: true},
		})
		if err != nil {
			log.Fatalf("failed to create category: %v", err)
		}

		categories, err := queries.ListCategories(ctx)
		if err != nil {
			log.Fatalf("failed to list categories: %v", err)
		}

		for _, category := range categories {
			log.Printf("category: %+v", category)
		}*/

	/*
		err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
			ID:          "9f94e9ef-4765-4b84-b596-01d1ce46eef3",
			Name:        "Backend updated",
			Description: sql.NullString{String: "Backend description updated", Valid: true},
		})
		if err != nil {
			log.Fatalf("failed to update category: %v", err)
		}

		categories, err := queries.ListCategories(ctx)
		if err != nil {
			log.Fatalf("failed to list categories: %v", err)
		}

		for _, category := range categories {
			log.Printf("category: %+v", category)
		}
	*/
	err = queries.DeleteCategory(ctx, "9f94e9ef-4765-4b84-b596-01d1ce46eef3")
	if err != nil {
		log.Fatalf("failed to delete category: %v", err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		log.Fatalf("failed to list categories: %v", err)
	}

	for _, category := range categories {
		log.Printf("category: %+v", category)
	}
}
