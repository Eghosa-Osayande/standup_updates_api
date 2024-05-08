package main


import (
	"standup-api/lib/features/admin"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	// admin
	adminRepo := admin.NewAdminRepo()
	adminservice:= admin.NewAdminService(adminRepo)
	admin.SetupAdminHandlers(v1,adminservice)

	

	router.Run(":8080")
}
