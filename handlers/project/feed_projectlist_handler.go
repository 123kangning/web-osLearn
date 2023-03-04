package project

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-osLearn/models"
)

type feedResponse struct {
	models.CommonResponse
	models.ProjectList
}

func FeedProjectList(c *gin.Context) {
	c.JSON(http.StatusOK, prepareData(getProjectList()))
}
func getProjectList() (models.ProjectList, error) {
	projectDAO := models.NewProjectDAO()
	return projectDAO.QueryAllProject()
}
func prepareData(pl models.ProjectList, err error) feedResponse {

	res := feedResponse{
		models.CommonResponse{StatusMsg: err.Error()},
		pl,
	}
	if err != nil {
		res.StatusCode = 1
	}
	return res
}
