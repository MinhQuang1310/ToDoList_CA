package handlers

import (
	models "cleanAchitech/entities"
	"cleanAchitech/infrastucture/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TodoItemHandler handles todo item related HTTP requests
type TodoItemHandler struct {
	usecase *usecases.TodoItemUsecase
}

// NewTodoItemHandler creates a new TodoItemHandler instance
func NewTodoItemHandler(u *usecases.TodoItemUsecase) *TodoItemHandler {
	return &TodoItemHandler{usecase: u}
}

// CreateItem handles HTTP POST requests to create a new todo item
func (h *TodoItemHandler) CreateItem(c *gin.Context) {
	var item models.TodoItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.CreateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Item created successfully",
		"data":    item,
	})
}

// GetItem handles HTTP GET requests to get a todo item by ID
func (h *TodoItemHandler) GetItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := h.usecase.GetItem(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

// GetItems handles HTTP GET requests to get all todo items
func (h *TodoItemHandler) GetItems(c *gin.Context) {
	items, err := h.usecase.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

// UpdateItem handles HTTP PUT requests to update a todo item by ID
func (h *TodoItemHandler) UpdateItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var item models.TodoItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = uint(id)
	if err := h.usecase.UpdateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Item updated successfully",
		"data":    item,
	})
}

// DeleteItem handles HTTP DELETE requests to delete a todo item by ID
func (h *TodoItemHandler) DeleteItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.DeleteItem(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Item deleted successfully",
	})
}
