package domain

import models "cleanAchitech/entities"

type TodoItemRepository interface {
	// Create creates a new TodoItem.
	Create(item *models.TodoItem) error

	// GetByID retrieves a TodoItem by its id.
	GetByID(id uint) (*models.TodoItem, error)

	// GetAll retrieves all TodoItems.
	GetAll() ([]models.TodoItem, error)

	// Update updates a TodoItem..
	Update(item *models.TodoItem) error

	// Delete deletes a TodoItem by its id.
	Delete(id uint) error

	// GetByTitle retrieves a TodoItem by its title.
	GetByTitle(title string, item *models.TodoItem) error

	// CreateWithProcedure creates a new TodoItem using PostgreSQL procedure.
	CreateWithProcedure(item *models.TodoItem) error
}
