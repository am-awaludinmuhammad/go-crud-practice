package main

import (
	"go-crud-practice/internal/categories"
	"go-crud-practice/internal/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db.LoadEnv()

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = db.AutoMigrate(&categories.Category{})
	if err != nil {
		log.Fatal("failed to migrate", err)
	}
	r := gin.Default()

	// Init category module
	categoryRepo := categories.NewCategoryRepository(db)
	categoryService := categories.NewCategoryService(categoryRepo)
	categoryHandler := categories.NewCategoryHandler(categoryService)

	// register routes
	categories.RegisterCategoryRoutes(r, categoryHandler)

	if err := r.Run(); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
