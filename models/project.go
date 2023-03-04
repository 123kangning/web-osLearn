package models

import "sync"

type ProjectList []*Project
type Project struct {
	Id          int    `json:"id" gorm:"primary"`
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

func (*ProjectDAO) QueryAllProject() (ProjectList, error) {
	var projects ProjectList
	result := DB.Find(&projects)
	return projects, result.Error
}

func (*ProjectDAO) AddProject(project Project) error {
	result := DB.Create(&project)
	return result.Error
}
