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
		adminEmployeesRouter.GET("/all", handler.FindAllEmployees)
		
	}

	employeesRouter.POST("/login", handler.login)
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


func (h *EmployeeHandlers) login(c *gin.Context) {

	var input, err = validator.DecodeAndValidateRequestBody[EmployeeLoginInputDto](c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(err.Code,err.ToResponse())
		return
	}

	result, loginErr := h.employeeService.Login(input)

	if loginErr != nil {
		c.AbortWithStatusJSON(loginErr.Code, loginErr.ToResponse())
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *EmployeeHandlers) FindAllEmployees(c *gin.Context) {

	input,err:=validator.DecodeAndValidateRequestBody[FindAllEmployeesInputDto](c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(err.Code,err.ToResponse())
		return
	}

	result, err := h.employeeService.FindAllEmployees(input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, http_response.NewHttpResponseWithError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}