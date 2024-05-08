package updates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateHandlers struct {
	taskService UpdateService
}

func SetupUpdateHandlers(r *gin.RouterGroup, taskService UpdateService) {
	var handler = UpdateHandlers{taskService: taskService}
	SprintRouter := r.Group("/updates")
	SprintRouter.POST("/login", handler.Login)
}

func (h *UpdateHandlers) Login(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{})
}
