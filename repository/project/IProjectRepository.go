package project

import "github.com/aguncn/nezha/models"

//IProjectRepository Project接口定义
type IProjectRepository interface {
	//GetProject 根据id获取Project
	GetProject(where interface{}) *models.Project
	//AddProject 新增Project
	AddProject(Project *models.Project) bool
	//UpdateProject 更新Project
	UpdateProject(Project *models.Project) bool
	//GetProjects 获取Project
	GetProjects(PageNum, PageSize uint, total *uint64, where interface{}) *[]models.Project
	//ExistProjectByName 是否存在已有应用
	ExistProjectByName(where interface{}) bool
	//DeleteProject 删除已有应用
	DeleteProject(id uint) bool
}
