package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hasona23/workis/api/models"
	"github.com/hasona23/workis/api/services"
)

func CreateQualification(c *gin.Context) {
	var qualification models.QualificationCreateRequest

	if err := c.ShouldBind(&qualification); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	err := services.CreateQualification(qualification)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Qualification created successfully"})
}

func DeleteQualification(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	err = services.DeleteQualification(id)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"msg": "qualification deleted successfuly"})
}

func UpdateQualification(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	var updateQualification models.Qualification
	if err = c.ShouldBind(&updateQualification); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	if id != updateQualification.ID {
		handleError(c, http.StatusBadRequest, fmt.Errorf("ids dont match"))
		return
	}

	err = services.UpdateQualification(updateQualification)
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"msg": "Worker updated successfuly"})
}

func AddQualificationHandlers(router *gin.Engine) {
	g := router.Group("/qualification")
	g.POST("/", CreateQualification)
	g.PUT("/:id", UpdateQualification)
	g.DELETE("/:id", DeleteQualification)
}
