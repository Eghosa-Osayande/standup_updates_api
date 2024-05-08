package main

import (
	"standup-api/lib/features/admin"
	"standup-api/lib/features/employee"
	"standup-api/lib/features/sprint"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	router := gin.Default()

	v1 := router.Group("/v1")

	{
		adminRepo := admin.NewAdminRepo()
		adminservice := admin.NewAdminService(adminRepo)
		admin.SetupAdminHandlers(v1, adminservice)
	}

	{
		employeeRepo := employee.NewEmployeeRepo()
		employeeService := employee.NewEmployeeService(employeeRepo)
		employee.SetupEmployeeHandlers(v1, employeeService)
	}

	{
		sprintRepo := sprint.NewSprintRepo()
		sprintService := sprint.NewSprintService(sprintRepo)
		sprint.SetupSprintHandlers(v1, sprintService)
	}

	

	router.Run(":8080")
}
