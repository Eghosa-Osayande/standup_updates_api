package sprint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TasksHandlers struct {
	taskService TaskService
}

func SetupTasksHandlers(r *gin.RouterGroup, taskService TaskService) {
	var handler = TasksHandlers{taskService: taskService}
	SprintRouter := r.Group("/tasks")
	SprintRouter.POST("/login", handler.Login)
}

func (h *TasksHandlers) Login(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{})
}
