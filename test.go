package main

import "web-osLearn/models"

func main() {
	models.InitDB()
	models.NewProjectDAO().AddProject(models.Project{Name: "Code", Author: "123kangning", Url: "https://github.com/123kangning/Code", Description: "练习代码"})
}
