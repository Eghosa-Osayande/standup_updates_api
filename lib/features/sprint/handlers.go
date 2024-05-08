package sprint

import (
	"net/http"
	"standup-api/lib/common/middleware"
	"standup-api/lib/utils/http_response"
	"standup-api/lib/utils/validator"
	"github.com/gin-gonic/gin"
)

type SprintHandlers struct {
	sprintService SprintService
}

func SetupSprintHandlers(r *gin.RouterGroup, sprintService SprintService) {
	var handler = SprintHandlers{sprintService: sprintService}
	SprintRouter := r.Group("/sprints")

	adminSprintRouter := SprintRouter.Group("")
	{
		adminSprintRouter.Use(middleware.AdminAuthMiddleware())
		adminSprintRouter.POST("/create", handler.createSprint)
		adminSprintRouter.POST("/start", handler.startSprint)
		// adminSprintRouter.POST("/end", handler.endSprint)
	}
}

func (h *SprintHandlers) createSprint(c *gin.Context) {
	input, err := validator.DecodeAndValidateRequestBody[CreateSprintInputDto](c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http_response.NewHttpResponseWithError(err.Error()))
		return
	}

	_, err = h.sprintService.CreateSprint(input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, http_response.NewHttpResponseWithError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, input)
}



func (h *SprintHandlers) startSprint(c *gin.Context) {
	

	_, err := h.sprintService.StartSprint()

	if err != nil {
		c.AbortWithStatusJSON(err.Code, err.ToResponse())
		return
	}

	c.JSON(http.StatusOK,"")
}