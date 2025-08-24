package main

import (
	"fmt"

	gin "github.com/gin-gonic/gin"
	"github.com/hasona23/workis/api/auth"
	"github.com/hasona23/workis/api/handlers"
	"github.com/hasona23/workis/api/models"
)

func main() {
	fmt.Println("APP BEGIN")
	models.InitDB()
	auth.CreateUserDB()
	router := gin.Default()

	router.MaxMultipartMemory = 32 * 1024 * 1024
	//router.Use(CORSMiddleware())
	//router.MaxMultipartMemory = 8 << 2
	handlers.AddWorkerHandler(router)
	handlers.AddQualificationHandlers(router)
	router.StaticFile("/", "./../web/index.html")
	router.Static("/web/imgs/", "./../web/imgs")
	router.Run("localhost:8080")
	fmt.Println("APP END")
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
