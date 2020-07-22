package application

import (
	"github.com/aguncn/nezha/models"
	//pageModel "github.com/aguncn/nezha/page"
	appRep "github.com/aguncn/nezha/repository/application"
	"github.com/jinzhu/gorm"
)

// ApplicationService 注入IApplicationRepository
type ApplicationService struct {
	Repository appRep.IApplicationRepository `inject:""`
}

//GetApplication 根据id获取Application
func (a *ApplicationService) GetApplication(id uint) *models.Application {
	where := models.Application{Model: gorm.Model{ID: id}}
	return a.Repository.GetApplication(&where)
}

//AddApplication 新增Application
func (a *ApplicationService) AddApplication(application *models.Application) bool {
	if a.Repository.ExistApplicationByName(application) {
		return false
	} else {
		return a.Repository.AddApplication(application)
	}
}

//UpdateApplication 更新Application
func (a *ApplicationService) UpdateApplication(application *models.Application) bool {
	/*
		此处需要改进，先注释。因为如果用户不改名称时，也会和已有名称冲突。要先判断，即ID不是自己的，但名称一样。
		if a.Repository.ExistApplicationByName(application) {
			return false
		} else {
			return a.Repository.UpdateApplication(application)
		}
	*/
	return a.Repository.UpdateApplication(application)
}

//GetApplications 获取Applications信息
func (a *ApplicationService) GetApplications(page, pagesize uint, maps interface{}) interface{} {
	res := make(map[string]interface{}, 2)
	var total uint64
	applications := a.Repository.GetApplications(page, pagesize, &total, maps)
	res["list"] = &applications
	res["total"] = total
	return &res
}

//RecentApplications 获取最近Application信息
func (a *ApplicationService) RecentApplications(maps interface{}, limit uint) interface{} {
	//res := make(map[string]interface{})
	applications := a.Repository.RecentApplications(maps, limit)
	//res["data"] = &applications
	return applications
}

//DeleteApplication 删除Application
func (a *ApplicationService) DeleteApplication(id uint) bool {
	return a.Repository.DeleteApplication(id)
}
