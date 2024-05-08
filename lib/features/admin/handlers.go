package admin

import (
	"net/http"
	"standup-api/lib/utils/http_response"
	"standup-api/lib/utils/validator"

	"github.com/gin-gonic/gin"
)

type AdminHandlers struct {
	AdminService AdminService
}

func SetupAdminHandlers(r *gin.RouterGroup, adminService AdminService) {
	var handler=AdminHandlers{AdminService: adminService}
	adminRouter := r.Group("/admin")
	adminRouter.POST("/login", handler.Login)
}

func (h *AdminHandlers) Login(c *gin.Context) {

	body, err := validator.DecodeAndValidateRequestBody[AdminLoginInputDto](c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, http_response.NewHttpResponseWithError(err.Error()))
		return
	}

	result, err := h.AdminService.Login(body)

	if err != nil {
		c.JSON(http.StatusUnauthorized, http_response.NewHttpResponseWithError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}
