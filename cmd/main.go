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
	// // Khởi tạo repository (giả định)
	// userRepo := repository.NewUserRepository() // Cần triển khai
	// // Khởi tạo use case
	// authUseCase := usecases.NewAuthUseCase(userRepo)

	// // Khởi tạo handler
	// authHandler := handlers.NewAuthHandler(authUseCase)
	// Set up router
	router := gin.Default()

	// Define API routes
	v1 := router.Group("/v1")
	{
		items := v1.Group("/items")
		{
			// Create a new Todoitems
			items.POST("/create", todoItemHandler.CreateItem)

			// Get a Todoitems by id
			items.GET("/getitem/:id", todoItemHandler.GetItem)

			// Get all Todoitems
			items.GET("/getall", todoItemHandler.GetItems)

			// Update a Todoitems
			items.PUT("/update/:id", todoItemHandler.UpdateItem)

			// Delete a Todoitems
			items.DELETE("/delete/:id", todoItemHandler.DeleteItem)

			// Restore a deleted Todoitems
			items.POST("/restore/:id", todoItemHandler.RestoreItem)
		}

	}

	// Start the server
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
