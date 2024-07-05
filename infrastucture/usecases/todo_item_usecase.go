package usecases

import (
	"cleanAchitech/config/domain"
	models "cleanAchitech/entities"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ErrItemTitleAlreadyExists = "Item title already exists"

// TodoItemUsecase is a struct that holds the repository for Todoitems
type TodoItemUsecase struct {
	repo domain.TodoItemRepository
}

// NewTodoItemUsecase returns a new instance of TodoItemUsecase
func NewTodoItemUsecase(repo domain.TodoItemRepository) *TodoItemUsecase {
	return &TodoItemUsecase{repo: repo}
}

// CreateItem creates a new Todoitems
func (u *TodoItemUsecase) CreateItem(item *models.TodoItem) error {
	// Check if item title is already exist
	var existingItem models.TodoItem
	err := u.repo.GetByTitle(item.Title, &existingItem)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == nil && existingItem.ID != item.ID {
		return &gin.Error{
			Err:  errors.New(ErrItemTitleAlreadyExists),
			Type: gin.ErrorTypePublic,
			Meta: ErrItemTitleAlreadyExists,
		}
	}
	return u.repo.Create(item)
}

// GetItem retrieves a Todoitems by its id
func (u *TodoItemUsecase) GetItem(id uint) (*models.TodoItem, error) {
	return u.repo.GetByID(id)
}

// GetItems retrieves all Todoitems
func (u *TodoItemUsecase) GetItems(filter func(db *gorm.DB) *gorm.DB) ([]models.TodoItem, error) {
	return u.repo.GetAll()
}

// UpdateItem updates a Todoitems
func (u *TodoItemUsecase) UpdateItem(item *models.TodoItem) error {
	// Check if item title is already exist
	var existingItem models.TodoItem
	err := u.repo.GetByTitle(item.Title, &existingItem)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	// Update only specified fields
	if item.Title != "" {
		existingItem.Title = item.Title
	}
	if item.Description != "" {
		existingItem.Description = item.Description
	}
	if item.Status != "" {
		existingItem.Status = item.Status
	}

	if err == nil && existingItem.ID != item.ID {
		return &gin.Error{
			Err:  errors.New(ErrItemTitleAlreadyExists),
			Type: gin.ErrorTypePublic,
			Meta: ErrItemTitleAlreadyExists,
		}
	}
	return u.repo.Update(&existingItem)
}

// DeleteItem deletes a Todoitems by its id
func (u *TodoItemUsecase) DeleteItem(id uint) error {
	return u.repo.Delete(id)
}

// CreateItem creates a new Todoitems using stored procedure
func (u *TodoItemUsecase) CreateItemWithProcedure(item *models.TodoItem) error {
	// Check if item title already exists
	var existingItem models.TodoItem
	err := u.repo.GetByTitle(item.Title, &existingItem)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == nil {
		return &gin.Error{
			Err:  errors.New(ErrItemTitleAlreadyExists),
			Type: gin.ErrorTypePublic,
			Meta: ErrItemTitleAlreadyExists,
		}
	}

	// Call stored procedure to create item
	return u.repo.CreateWithProcedure(item)
}
