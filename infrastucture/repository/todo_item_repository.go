package repository

import (
	"cleanAchitech/config/domain"
	models "cleanAchitech/entities"

	"gorm.io/gorm"
)

// TodoItemRepository struct represents the repository for todo items
type TodoItemRepository struct {
	db *gorm.DB // The database connection
}

// NewTodoItemRepository creates a new instance of TodoItemRepository
func NewTodoItemRepository(db *gorm.DB) domain.TodoItemRepository {
	return &TodoItemRepository{db: db}
}

// Create creates a new todo item in the repository
func (r *TodoItemRepository) Create(item *models.TodoItem) error {
	return r.db.Create(item).Error
}

// GetByID retrieves a todo item from the repository by its ID
func (r *TodoItemRepository) GetByID(id uint) (*models.TodoItem, error) {
	var item models.TodoItem
	err := r.db.First(&item, id).Error
	return &item, err
}

// GetAll retrieves all todo items from the repository
func (r *TodoItemRepository) GetAll() ([]models.TodoItem, error) {
	var items []models.TodoItem
	err := r.db.Find(&items).Error
	return items, err
}

// Update updates a todo item in the repository
func (r *TodoItemRepository) Update(item *models.TodoItem) error {
	return r.db.Save(item).Error
}

// Delete deletes a todo item from the repository by its ID
func (r *TodoItemRepository) Delete(id uint) error {
	return r.db.Delete(&models.TodoItem{}, id).Error
}
