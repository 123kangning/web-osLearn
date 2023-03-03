package router

import "github.com/gin-gonic/gin"

func initRouter() *gin.Engine {
	r := gin.Default()
	baseGroup := r.Group("/os")
	baseGroup.GET("/project/")
	return r
}
