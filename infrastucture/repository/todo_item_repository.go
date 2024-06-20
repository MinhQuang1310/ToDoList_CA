package repository

import (
	"cleanAchitech/config/domain"
	models "cleanAchitech/entities"

	"gorm.io/gorm"
)

// TodoItemRepository struct represents the repository for Todoitems
type TodoItemRepository struct {
	db *gorm.DB // The database connection
}

// NewTodoItemRepository creates a new instance of TodoItemRepository
func NewTodoItemRepository(db *gorm.DB) domain.TodoItemRepository {
	return &TodoItemRepository{db: db}
}

// Create creates a new Todoitems in the repository
func (r *TodoItemRepository) Create(item *models.TodoItem) error {
	return r.db.Create(item).Error
}

// GetByID retrieves a Todoitems from the repository by its ID
func (r *TodoItemRepository) GetByID(id uint) (*models.TodoItem, error) {
	var item models.TodoItem
	err := r.db.First(&item, id).Error
	return &item, err
}

// GetByTitle retrieves a Todoitems from the repository by its title
func (r *TodoItemRepository) GetByTitle(title string, item *models.TodoItem) error {
	return r.db.Where("title = ?", title).First(item).Error
}

// GetAll retrieves all Todoitemss from the repository
func (r *TodoItemRepository) GetAll() ([]models.TodoItem, error) {
	var items []models.TodoItem
	err := r.db.Find(&items).Error
	return items, err
}

// Update updates a Todoitems in the repository
func (r *TodoItemRepository) Update(item *models.TodoItem) error {
	return r.db.Save(item).Error
}

// Delete deletes a Todoitems from the repository by its ID
func (r *TodoItemRepository) Delete(id uint) error {
	return r.db.Delete(&models.TodoItem{}, id).Error
}

// CreateWithProcedure creates a new Todoitems using PostgreSQL procedure.
func (r *TodoItemRepository) CreateWithProcedure(item *models.TodoItem) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Exec("CALL insert_book(?, ?, ?)", item.Title, item.Description, item.Status).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
