package project

import (
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	baseRep "github.com/aguncn/nezha/repository"
	"github.com/jinzhu/gorm"
)

//ProjectRepository 注入IDb
type ProjectRepository struct {
	Log  logger.ILogger         `inject:""`
	Base baseRep.BaseRepository `inject:"inline"`
}

//GetProject 根据id获取Project
func (a *ProjectRepository) GetProject(where interface{}) *models.Project {
	var Project models.Project
	other := map[string]string{}
	if err := a.Base.First(where, &Project, other); err != nil {
		a.Log.Errorf("未找到相关文章", err)
	}
	return &Project
}

//AddProject 新增Project
func (a *ProjectRepository) AddProject(Project *models.Project) bool {
	if err := a.Base.Create(Project); err != nil {
		a.Log.Errorf("添加文章失败", err)
		return false
	}
	return true
}

//UpdateProject 更新Project
func (a *ProjectRepository) UpdateProject(Project *models.Project) bool {
	if err := a.Base.Save(Project); err != nil {
		a.Log.Errorf("更新应用失败", err)
		return false
	}
	return true
}

//GetProjects 获取项目列表
func (a *ProjectRepository) GetProjects(PageNum uint, PageSize uint, total *uint64, where interface{}) *[]models.Project {
	var Projects []models.Project
	other := map[string]string{
		"order": "ID desc",
	}
	err := a.Base.GetPages(&models.Project{}, &Projects, PageNum, PageSize, total, where, other)
	if err != nil {
		a.Log.Errorf("获取文章信息失败", err)
	}
	return &Projects
}

//ExistProjectByName 判断用户名是否已存在
func (a *ProjectRepository) ExistProjectByName(where interface{}) bool {
	var Project models.Project
	var whereProject = models.Project{Name: where.(*models.Project).Name}
	other := map[string]string{}
	err := a.Base.First(whereProject, &Project, other)
	//记录不存在错误(RecordNotFound)，返回false
	if gorm.IsRecordNotFoundError(err) {
		return false
	}
	//其他类型的错误，写下日志，返回false
	if err != nil {
		a.Log.Error(err)
		return false
	}
	return true
}

//DeleteProject 删除应用
func (a *ProjectRepository) DeleteProject(id uint) bool {
	var Project models.Project
	if err := a.Base.DeleteByID(Project, id); err != nil {
		a.Log.Errorf("删除应用失败", err)
		return false
	}
	return true
}
