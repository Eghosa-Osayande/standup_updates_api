package updates

import (
	"net/http"
	"standup-api/lib/common/middleware"
	"standup-api/lib/utils/http_response"
	"standup-api/lib/utils/validator"

	"github.com/gin-gonic/gin"
)

type UpdateHandlers struct {
	updateService UpdateService
}

func SetupUpdateHandlers(r *gin.RouterGroup, taskService UpdateService) {
	var handler = UpdateHandlers{updateService: taskService}
	updatesRouter := r.Group("/updates")
	
	employeeProtectedRouter := updatesRouter.Group("")

	{
		employeeProtectedRouter.Use(middleware.EmployeeAuthMiddleware())
		employeeProtectedRouter.POST("/create", handler.createUpdate)
	}
}

func (h *UpdateHandlers) createUpdate(c *gin.Context) {
	input, err := validator.DecodeAndValidateRequestBody[CreateUpdateInputDto](c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(err.Code,err.ToResponse())
		return
	}

	result, err := h.updateService.CreateUpdate(input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, http_response.NewHttpResponseWithError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}
