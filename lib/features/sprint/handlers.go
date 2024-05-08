package sprint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SprintHandlers struct {
	sprintService SprintService
}

func SetupSprintHandlers(r *gin.RouterGroup, sprintService SprintService) {
	var handler = SprintHandlers{sprintService: sprintService}
	SprintRouter := r.Group("/sprint")
	SprintRouter.POST("/login", handler.Login)
}

func (h *SprintHandlers) Login(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{})
}
