package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hasona23/workis/api/models"
	"github.com/hasona23/workis/api/repositories"
)

func RemoveQualification(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("error parsing qualification Id %v", err)})
		return
	}
	err = repositories.DeleteQualification(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("error removing qualification  %v", err)})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"msg": "qualification removed successfuly"})

}
func AddQualification(ctx *gin.Context) {
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

func AddQualificationEndPoints(router *gin.Engine) {
	router.DELETE("/qualifications/:id", RemoveQualification)
	router.POST("/qualifications/", AddQualification)
}
