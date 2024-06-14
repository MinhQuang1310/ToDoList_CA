package usecases

import (
	"cleanAchitech/config/db"
	"cleanAchitech/config/domain"
	models "cleanAchitech/entities"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TodoItemUsecase is a struct that holds the repository for todo items
type TodoItemUsecase struct {
	repo domain.TodoItemRepository
}

// NewTodoItemUsecase returns a new instance of TodoItemUsecase
func NewTodoItemUsecase(repo domain.TodoItemRepository) *TodoItemUsecase {
	return &TodoItemUsecase{repo: repo}
}

// CreateItem creates a new todo item
func (u *TodoItemUsecase) CreateItem(item *models.TodoItem) error {
	// Check if item title is already exist
	var existingItem models.TodoItem
	result := db.InitDB().Where("title = ?", item.Title).First(&existingItem)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}
	if result.RowsAffected > 0 {
		return &gin.Error{
			Err:  errors.New("item title already exists"),
			Type: gin.ErrorTypePublic,
			Meta: "Item title already exists",
		}
	}
	return u.repo.Create(item)
}

// GetItem retrieves a todo item by its id
func (u *TodoItemUsecase) GetItem(id uint) (*models.TodoItem, error) {
	return u.repo.GetByID(id)
}

// GetItems retrieves all todo items
func (u *TodoItemUsecase) GetItems() ([]models.TodoItem, error) {
	return u.repo.GetAll()
}

// UpdateItem updates a todo item
func (u *TodoItemUsecase) UpdateItem(item *models.TodoItem) error {
	// Check if item title is already exist
	var existingItem models.TodoItem
	result := db.InitDB().Where("title = ?", item.Title).First(&existingItem)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}
	if result.RowsAffected > 0 {
		return &gin.Error{
			Err:  errors.New("item title already exists"),
			Type: gin.ErrorTypePublic,
			Meta: "Item title already exists",
		}
	}
	return u.repo.Update(item)
}

// DeleteItem deletes a todo item by its id
func (u *TodoItemUsecase) DeleteItem(id uint) error {
	return u.repo.Delete(id)
}
