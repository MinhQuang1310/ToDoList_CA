package main

import (
	"cleanAchitech/api/handlers"
	"cleanAchitech/config/db"
	models "cleanAchitech/entities"

	"cleanAchitech/infrastucture/repository"
	"cleanAchitech/infrastucture/usecases"

	"github.com/gin-gonic/gin"
)

// main is the entry point of the application.
// It initializes the database, migrates the schema,
// sets up the repository, usecase, and handler,
// configures the router, and starts the server.
func main() {
	// Initialize database connection
	gormDB := db.InitDB()
	sqlDB, err := gormDB.DB()
	if err != nil {
		panic("failed to get sql.DB")
	}
	defer sqlDB.Close()

	// Migrate the schema
	gormDB.AutoMigrate(&models.TodoItem{})

	// Set up repository, usecase, and handler
	todoItemRepo := repository.NewTodoItemRepository(gormDB)
	todoItemUsecase := usecases.NewTodoItemUsecase(todoItemRepo)
	todoItemHandler := handlers.NewTodoItemHandler(todoItemUsecase)

	// Set up router
	router := gin.Default()

	// Define API routes
	v1 := router.Group("/v1")
	{
		items := v1.Group("/items")
		{
			// Create a new todo item
			items.POST("/create", todoItemHandler.CreateItem)

			// Get a todo item by id
			items.GET("/getitem/:id", todoItemHandler.GetItem)

			// Get all todo items
			items.GET("/getall", todoItemHandler.GetItems)

			// Update a todo item
			items.PUT("/update/:id", todoItemHandler.UpdateItem)

			// Delete a todo item
			items.DELETE("/delete/:id", todoItemHandler.DeleteItem)
		}
	}

	// Start the server
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
