package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeHandlers struct {
	employeeService EmployeeService
}

func SetupEmployeeHandlers(r *gin.RouterGroup, employeeService EmployeeService) {
	var handler = EmployeeHandlers{employeeService: employeeService}
	SprintRouter := r.Group("/updates")
	SprintRouter.POST("/login", handler.Login)
}

func (h *EmployeeHandlers) Login(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{})
}
