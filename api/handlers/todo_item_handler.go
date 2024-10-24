package handlers

import (
	models "cleanAchitech/entities"
	"cleanAchitech/infrastucture/usecases"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var errItemNotFound = errors.New("item not found")

// formatResponse formats the response data with additional properties
func formatResponse(_ *gin.Context, code string, message string, data interface{}, total int, pageNumber int, pageSize int) gin.H {
	return gin.H{
		"response": gin.H{
			"responseId":      uuid.New().String(),
			"responseCode":    code,
			"responseMessage": message,
			"responseTime":    time.Now().Format(time.RFC3339),
		},
		"result": gin.H{
			"data":       data,
			"total":      total,
			"pageNumber": pageNumber,
			"pageSize":   pageSize,
		},
	}
}

// TodoItemHandler handles TodoItem related HTTP requests
type TodoItemHandler struct {
	usecase *usecases.TodoItemUsecase
}

// NewTodoItemHandler creates a new TodoItemHandler instance
func NewTodoItemHandler(u *usecases.TodoItemUsecase) *TodoItemHandler {
	return &TodoItemHandler{usecase: u}
}

// CreateItem handles HTTP POST requests to create a new TodoItem
func (h *TodoItemHandler) CreateItem(c *gin.Context) {
	var item models.TodoItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(c, strconv.Itoa(http.StatusBadRequest), err.Error(), nil, 0, 0, 0))
		return
	}

	if err := h.usecase.CreateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(c, strconv.Itoa(http.StatusInternalServerError), err.Error(), nil, 0, 0, 0))
		return
	}
	c.JSON(http.StatusCreated, formatResponse(c, strconv.Itoa(http.StatusOK), "Item created successfully!", item, 1, 0, 1))
}

// GetItem handles HTTP GET requests to get a TodoItem by ID
func (h *TodoItemHandler) GetItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := h.usecase.GetItem(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, formatResponse(c, strconv.Itoa(http.StatusNotFound), errItemNotFound.Error(), nil, 0, 0, 0))
		} else {
			c.JSON(http.StatusInternalServerError, formatResponse(c, strconv.Itoa(http.StatusInternalServerError), err.Error(), nil, 0, 0, 0))
		}
		return
	}
	c.JSON(http.StatusOK, formatResponse(c, strconv.Itoa(http.StatusOK), "Item retrieved successfully!", item, 1, 0, 1))
}

// GetItems handles HTTP GET requests to get all TodoItems
func (h *TodoItemHandler) GetItems(c *gin.Context) {
	items, err := h.usecase.GetItems(func(db *gorm.DB) *gorm.DB {
		return db.Not("DeletedAt", nil)
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(c, strconv.Itoa(http.StatusBadRequest), err.Error(), nil, 0, 0, 0))
		return
	}
	c.JSON(http.StatusOK, formatResponse(c, strconv.Itoa(http.StatusOK), "Items retrieved successfully!", items, len(items), 0, len(items)))
}

// UpdateItem handles HTTP PUT requests to update a TodoItem by ID
func (h *TodoItemHandler) UpdateItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var item models.TodoItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(c, strconv.Itoa(http.StatusBadRequest), err.Error(), nil, 0, 0, 0))
		return
	}

	// Set the ID of the item to update
	item.ID = uint(id)

	// Check if the item exists in the database
	if _, err := h.usecase.GetItem(item.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, formatResponse(c, strconv.Itoa(http.StatusNotFound), errItemNotFound.Error(), nil, 0, 0, 0))
		} else {
			c.JSON(http.StatusInternalServerError, formatResponse(c, strconv.Itoa(http.StatusInternalServerError), err.Error(), nil, 0, 0, 0))
		}
		return
	}

	// Update the item in the usecase
	if err := h.usecase.UpdateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(c, strconv.Itoa(http.StatusInternalServerError), err.Error(), nil, 0, 0, 0))
		return
	}
	c.JSON(http.StatusOK, formatResponse(c, strconv.Itoa(http.StatusOK), "Item updated successfully!", item, 1, 0, 1))
}

// DeleteItem handles HTTP DELETE requests to delete a TodoItem by ID
func (h *TodoItemHandler) DeleteItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// Check if the item exists in the database
	_, err := h.usecase.GetItem(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, formatResponse(c, strconv.Itoa(http.StatusNotFound), errItemNotFound.Error(), nil, 0, 0, 0))
		} else {
			c.JSON(http.StatusInternalServerError, formatResponse(c, strconv.Itoa(http.StatusInternalServerError), err.Error(), nil, 0, 0, 0))
		}
		return
	}

	// Delete the item from the usecase
	if err := h.usecase.DeleteItem(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(c, strconv.Itoa(http.StatusInternalServerError), err.Error(), nil, 0, 0, 0))
		return
	}
	c.JSON(http.StatusOK, formatResponse(c, strconv.Itoa(http.StatusOK), "Item deleted successfully!", nil, 0, 0, 0))
}

func (h *TodoItemHandler) RestoreItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// Check if the item exists in the database
	_, err := h.usecase.GetItem(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, formatResponse(c, strconv.Itoa(http.StatusNotFound), errItemNotFound.Error(), nil, 0, 0, 0))
		} else {
			c.JSON(http.StatusInternalServerError, formatResponse(c, strconv.Itoa(http.StatusInternalServerError), err.Error(), nil, 0, 0, 0))
		}
		return
	}

	// Delete the item from the usecase
	if err := h.usecase.Restore(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(c, strconv.Itoa(http.StatusInternalServerError), err.Error(), nil, 0, 0, 0))
		return
	}
	c.JSON(http.StatusOK, formatResponse(c, strconv.Itoa(http.StatusOK), "Item restored successfully!", nil, 0, 0, 0))
}
