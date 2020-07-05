package project

import (
	"github.com/aguncn/nezha/models"
)

//IProjectService Project接口定义
type IProjectService interface {
	//GetProject 根据id获取Project
	GetProject(id uint) *models.Project

	//AddProject 新增Project
	AddProject(Project *models.Project) bool
	//UpdateProject 更新Project
	UpdateProject(Project *models.Project) bool
	//GetProjects 获取文章信息
	GetProjects(page, pagesize uint, maps interface{}) interface{}
	//DeleteProject 删除Project
	DeleteProject(id uint) bool
}
