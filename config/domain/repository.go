package domain

import models "cleanAchitech/entities"

type TodoItemRepository interface {
	// Create creates a new todo item.
	Create(item *models.TodoItem) error

	// GetByID retrieves a todo item by its id.
	GetByID(id uint) (*models.TodoItem, error)

	// GetAll retrieves all todo items.
	GetAll() ([]models.TodoItem, error)

	// Update updates a todo item..
	Update(item *models.TodoItem) error

	// Delete deletes a todo item by its id.
	Delete(id uint) error

	// GetByTitle retrieves a todo item by its id.
	GetByTitle(title string) (*models.TodoItem, error)
}
