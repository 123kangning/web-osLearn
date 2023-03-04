package models

import "sync"

type ProjectList []*Project
type Project struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Url         string `json:"url"`
	Description string `json:"description"`
}
type ProjectDAO struct {
}

var (
	projectDAO  *ProjectDAO
	projectOnce sync.Once
)

func NewProjectDAO() *ProjectDAO {
	projectOnce.Do(func() {
		projectDAO = &ProjectDAO{}
	})
	return projectDAO
}

func QueryAllProject() ProjectList {
	var projects ProjectList
	DB.Find(&projects)
	return projects
}
