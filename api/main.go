package main

import (
	"fmt"

	gin "github.com/gin-gonic/gin"
	"github.com/hasona23/workis/api/model"
)

func main() {
	fmt.Println("APP BEGIN")
	model.InitDB()
	model.SeedData()

	fmt.Println(model.Db == nil)
	var workers []model.Worker
	model.Db.Preload("Qualifications").Find(&workers)
	fmt.Println("Workers Count: ", len(workers))
	fmt.Println("Workers: ")
	for i := 0; i < len(workers); i++ {
		fmt.Println(workers[i])
	}

	fmt.Println("APP END")
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
