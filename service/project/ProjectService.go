package project

import (
	"github.com/aguncn/nezha/models"
	//pageModel "github.com/aguncn/nezha/page"
	appPro "github.com/aguncn/nezha/repository/project"
	"github.com/jinzhu/gorm"
)

// ArticleService 注入IArticleRepo
type ProjectService struct {
	Repository appPro.IProjectRepository `inject:""`
}

//GetProject 根据id获取Project
func (a *ProjectService) GetProject(id uint) *models.Project {
	where := models.Project{Model: gorm.Model{ID: id}}
	return a.Repository.GetProject(&where)
}

//AddProject 新增Project
func (a *ProjectService) AddProject(Project *models.Project) bool {

	return a.Repository.AddProject(Project)

}

//UpdateProject 更新Project
func (a *ProjectService) UpdateProject(Project *models.Project) bool {
	if a.Repository.ExistProjectByName(Project) {
		return false
	} else {
		return a.Repository.UpdateProject(Project)
	}
}

//GetProjects 获取Project信息
func (a *ProjectService) GetProjects(page, pagesize uint, maps interface{}) interface{} {
	res := make(map[string]interface{}, 2)
	var total uint64
	Projects := a.Repository.GetProjects(page, pagesize, &total, maps)
	res["list"] = &Projects
	res["total"] = total
	return &res
}

//DeleteProject 删除Project
func (a *ProjectService) DeleteProject(id uint) bool {
	return a.Repository.DeleteProject(id)
}
