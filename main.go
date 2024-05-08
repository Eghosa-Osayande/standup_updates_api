package main

import (
	"standup-api/lib/common/database"
	"standup-api/lib/features/employee"
	"standup-api/lib/features/sprint"
	"standup-api/lib/features/updates"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading environment variables")
	}

	db, err := database.GetDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	router := gin.Default()

	v1 := router.Group("/v1")

	{
		employeeRepo := employee.NewEmployeeRepo(db)
		employeeService := employee.NewEmployeeService(employeeRepo)
		employee.SetupEmployeeHandlers(v1, employeeService)
	}

	{
		sprintRepo := sprint.NewSprintRepo(db)
		sprintService := sprint.NewSprintService(sprintRepo)
		sprint.SetupSprintHandlers(v1, sprintService)
	}
	{
		updatesRepo := updates.NewUpdatesRepo(db)
		updatesService := updates.NewUpdateService(updatesRepo)
		updates.SetupUpdateHandlers(v1, updatesService)
	}

	router.Run(":8080")
}
