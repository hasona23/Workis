package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hasona23/workis/api/models"
	"github.com/hasona23/workis/api/repositories"
)

func CreateWorkerHandler(ctx *gin.Context) {
	var worker models.CreateWorkerDto
	if err := ctx.ShouldBind(&worker); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Couldn't parse worker"})
		return
	}
	err := repositories.CreateWorker(worker.ToWorker())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("Error Creating worker %s", err)})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": "Created worker successfuly"})
}

func GetAllWorkersHandler(ctx *gin.Context) {
	workers, err := repositories.GetAllWorkers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("Error Getting workers %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"workers": workers})
}
func GetWorkerByIdHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Error parsing id %v", err)})
		return
	}
	worker, err := repositories.GetWorkerByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Error getting worker %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"worker": worker})
}

func UpdateWorkerHandler(ctx *gin.Context) {
	var newWorker models.UpdateWorkerDto
	if err := ctx.ShouldBind(&newWorker); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("failed to parse worker data %v", err)})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("InValid Id %v", err)})
		return
	}
	if id != newWorker.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Ids dont match %v and %v", id, newWorker.ID)})
		return
	}
	worker, err := repositories.GetWorkerByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("couldnt find worker with id %v", newWorker.ID)})
		return
	}
	newWorker.CopyValuesToWorker(worker)
	err = repositories.UpdateWorker(id, *worker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("error updating worker with id %v : %v", newWorker.ID, err)})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"msg": fmt.Sprintf("Updated worker with id updating id %v successfuly", newWorker.ID)})

}

func DeleteWorkerHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("InValid Id %v", err)})
		return
	}
	worker, err := repositories.GetWorkerByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("error getting worker with id %v %v", id, err)})
		return
	}
	worker.IsActive = false
	err = repositories.UpdateWorker(worker.ID, *worker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("error soft deleting worker %v", err)})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"msg": fmt.Sprintf("successfully soft deleted worker with id %v", worker.ID)})
}

func EnableWorkerHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("InValid Id %v", err)})
		return
	}
	worker, err := repositories.GetWorkerByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("error getting worker with id %v %v", id, err)})
		return
	}
	worker.IsActive = true
	err = repositories.UpdateWorker(worker.ID, *worker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("error enabling worker %v", err)})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"msg": fmt.Sprintf("successfully enabled worker with id %v", worker.ID)})
}

func AddWorkerEndPoints(router *gin.Engine) {
	router.POST("/workers", CreateWorkerHandler)
	router.GET("/workers", GetAllWorkersHandler)
	router.GET("/workers/:id", GetWorkerByIdHandler)
	router.PUT("/workers/:id", UpdateWorkerHandler)
	router.PUT("/workers/enable/id", EnableWorkerHandler)
	router.DELETE("/workers/:id", DeleteWorkerHandler)
}
