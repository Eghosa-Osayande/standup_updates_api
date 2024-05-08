package employee

import (
	"net/http"
	"standup-api/lib/common/middleware"
	"standup-api/lib/utils/http_response"
	"standup-api/lib/utils/validator"

	"github.com/gin-gonic/gin"
)

type EmployeeHandlers struct {
	employeeService EmployeeService
}

func SetupEmployeeHandlers(r *gin.RouterGroup, employeeService EmployeeService) {
	var handler = EmployeeHandlers{employeeService: employeeService}
	employeesRouter := r.Group("/employees")

	adminEmployeesRouter := employeesRouter.Group("")
	{
		adminEmployeesRouter.Use(middleware.AdminAuthMiddleware())
		adminEmployeesRouter.POST("/create", handler.CreateEmployee)
	}
}

func (h *EmployeeHandlers) CreateEmployee(c *gin.Context) {

	input, err := validator.DecodeAndValidateRequestBody[CreateEmployeeInputDto](c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http_response.NewHttpResponseWithError(err.Error()))
		return
	}

	result, err := h.employeeService.CreateEmployee(input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, http_response.NewHttpResponseWithError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}
