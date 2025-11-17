package todo

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	todos   = make(map[string]Todo)
	todosMu sync.RWMutex
)

func RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/todos")

	r.GET("", listTodos)
	r.POST("", createTodo)
	r.GET("/:id", getTodo)
	r.PUT("/:id", updateTodo)
	r.DELETE("/:id", deleteTodo)
}

func listTodos(c *gin.Context) {
	todosMu.RLock()
	defer todosMu.RUnlock()

	list := make([]Todo, 0, len(todos))
	for _, t := range todos {
		list = append(list, t)
	}

	c.JSON(http.StatusOK, list)
}

func createTodo(c *gin.Context) {
	var input struct {
		Title string `json:"title" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.NewString()
	t := Todo{
		ID:        id,
		Title:     input.Title,
		Completed: false,
	}

	todosMu.Lock()
	todos[id] = t
	todosMu.Unlock()

	c.JSON(http.StatusCreated, t)
}

func getTodo(c *gin.Context) {
	id := c.Param("id")

	todosMu.RLock()
	t, ok := todos[id]
	todosMu.RUnlock()

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	c.JSON(http.StatusOK, t)
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todosMu.Lock()
	defer todosMu.Unlock()

	t, ok := todos[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	if input.Title != nil {
		t.Title = *input.Title
	}
	if input.Completed != nil {
		t.Completed = *input.Completed
	}

	todos[id] = t
	c.JSON(http.StatusOK, t)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	todosMu.Lock()
	defer todosMu.Unlock()

	if _, ok := todos[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	delete(todos, id)
	c.Status(http.StatusNoContent)
}
