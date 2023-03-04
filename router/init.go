package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web-osLearn/handlers/project"
	"web-osLearn/models"
)

func InitRouter() *gin.Engine {
	models.InitDB()
	fmt.Println(4)
	r := gin.Default()
	baseGroup := r.Group("/os")
	baseGroup.GET("/project/", project.FeedProjectList)
	return r
}
