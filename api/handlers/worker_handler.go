package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hasona23/workis/api/models"
	"github.com/hasona23/workis/api/services"
)

func CreateWorker(c *gin.Context) {
	var workerCreateRequest models.WorkerCreateRequest

	if err := json.Unmarshal([]byte(c.Request.FormValue("ModelData")), &workerCreateRequest); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	fmt.Println(workerCreateRequest)
	faceImg, faceImgHeader, err := c.Request.FormFile("faceImg")

	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(faceImgHeader.Filename, "  ", faceImgHeader.Size)
	}
	idImg, idImgHeader, err := c.Request.FormFile("idImg")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(idImgHeader.Filename, "  ", idImgHeader.Size)
	}
	faceImgFile := models.FileRequest{File: faceImg, Header: faceImgHeader}
	idImgFile := models.FileRequest{File: idImg, Header: idImgHeader}

	defer faceImg.Close()
	defer idImg.Close()

	err = services.CreateWorker(workerCreateRequest, faceImgFile, idImgFile)

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
func UpdateWorkerImg(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	isFaceImg, err := strconv.ParseBool(c.Query("isFaceImg"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	var key string
	if isFaceImg {
		key = "faceImg"
	} else {
		key = "idImg"
	}
	imgFile, err := getImageFile(key, c)
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	err = services.UpdateWokerImg(id, *imgFile, isFaceImg)
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"msg": "worker updated successfuly"})
}
func AddWorkerHandler(router *gin.Engine) {
	g := router.Group("/worker")

	g.POST("/", CreateWorker)
	g.GET("/", GetAllWorkers)
	g.GET("/:id", GetWorkerWithID)
	g.DELETE("/:id", SoftDeleteWorker)
	g.PATCH("/:id", ReviveWorker)
	g.PATCH("/img/:id", UpdateWorkerImg)
	g.PUT("/:id", UpdateWorker)

}

func handleError(c *gin.Context, errorCode int, err error) {
	c.JSON(errorCode, gin.H{"msg": fmt.Sprintf("Error: %v", err)})
}

func getImageFile(key string, c *gin.Context) (*models.FileRequest, error) {
	img, imgHeader, err := c.Request.FormFile(key)
	if err != nil {
		return nil, err
	}
	imgFile := models.FileRequest{File: img, Header: imgHeader}
	return &imgFile, nil
}
