package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hasona23/workis/api/models"
	"github.com/hasona23/workis/api/repositories"
)

func DeleteQualification(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("InValid Id %v", err)})
		return
	}
	if repositories.QualificationIdExist(id) {
		err := repositories.DeleteQualification(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("error deleting qualification %v", err)})
			return
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{"msg": fmt.Sprintf("successfuly delete qualification of id %v", id)})
			return
		}
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": fmt.Sprintf("not found qualification with id  %v", id)})
		return
	}

}
func CreateQualificationHandler(ctx *gin.Context) {
	var qualification models.Qualification
	if err := ctx.ShouldBind(&qualification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("error getting qualification  %v", err)})
		return
	}
	err := repositories.CreateQualification(qualification.WorkerId, qualification.CertName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("error creating qualification  %v", err)})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": "qualification created successfuly"})
}

func UpdateQualificationHandler(ctx *gin.Context) {
	var newQualification models.Qualification
	if err := ctx.ShouldBind(&newQualification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("failed to parse qualification data %v", err)})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("InValid Id %v", err)})
		return
	}
	if id != newQualification.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Ids dont match %v and %v", id, newQualification.ID)})
		return
	}
	qualification, err := repositories.GetQualificationByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("couldnt find worker with id %v", newQualification.ID)})
		return
	}
	if newQualification.CertName != "" {
		qualification.CertName = newQualification.CertName
	}
	err = repositories.UpdateQualification(id, *qualification)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("error updating qualification with id %v : %v", newQualification.ID, err)})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"msg": fmt.Sprintf("Updated qualification with id updating id %v successfuly", newQualification.ID)})

}

func AddQualificationEndPoints(router *gin.Engine) {
	router.DELETE("/qualifications/:id", DeleteQualification)
	router.POST("/qualifications/", CreateQualificationHandler)
	router.PUT("/qualifications/:id", UpdateQualificationHandler)
}
