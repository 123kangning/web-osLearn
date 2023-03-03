package project

import (
	"github.com/gin-gonic/gin"
	"web-osLearn/models"
)

type feedResponse struct {
	models.CommonResponse
	models.ProjectList
}

func feedProjectList(c *gin.Context) {

}
func getProjectList()models.ProjectList{
	projectDAO:=models.NewProjectDAO()
	projectDAO.
}
