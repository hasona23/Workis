package main

import (
	"fmt"

	gin "github.com/gin-gonic/gin"
	"github.com/hasona23/workis/api/handlers"
	"github.com/hasona23/workis/api/models"
)

func main() {

	models.InitDB()
	defer models.DB.Close()

	models.DeleteWorkerTable()
	models.CreateWorkerTable()
	models.SeedWorkers()

	router := gin.Default()
	handlers.AddWorkerEndPoints(router)
	handlers.AddQualificationEndPoints(router)
	router.Run("localhost:8080")

	fmt.Println("Server closed")
}
