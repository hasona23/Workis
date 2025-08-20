package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hasona23/workis/api/models"
	"github.com/hasona23/workis/api/services"
)

func CreateWorker(c *gin.Context) {
	var workerCreateRequest models.WorkerCreateRequest

	if err := c.ShouldBind(&workerCreateRequest); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	err := services.CreateWorker(workerCreateRequest, services.FileRequest{}, services.FileRequest{})
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Worker created successfully"})
}

func GetAllWorkers(c *gin.Context) {
	workers, err := services.GetAllWorkers()
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, workers)
}
func GetWorkerWithID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	worker, err := services.GetWorkerByID(id)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, worker)
}
func SoftDeleteWorker(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	err = services.SoftDeleteWorker(id)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"msg": "worker deleted successfuly"})
}
func ReviveWorker(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	err = services.ReviveWorker(id)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"msg": "worker deleted successfuly"})
}
func UpdateWorker(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	var updateWorker models.WorkerUpdateRequest
	if err = c.ShouldBind(&updateWorker); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	if id != updateWorker.ID {
		handleError(c, http.StatusBadRequest, fmt.Errorf("ids dont match"))
		return
	}

	err = services.UpdateWorker(updateWorker)
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"msg": "Worker updated successfuly"})
}
func AddWorkerHandler(router *gin.Engine) {
	g := router.Group("/worker")
	g.POST("/", CreateWorker)
	g.GET("/", GetAllWorkers)
	g.GET("/:id", GetWorkerWithID)
	g.DELETE("/:id", SoftDeleteWorker)
	g.PATCH("/:id", ReviveWorker)
	g.PUT("/:id", UpdateWorker)
}

func handleError(c *gin.Context, errorCode int, err error) {
	c.JSON(errorCode, gin.H{"msg": fmt.Sprintf("Error: %v", err)})
}
