package application

import "github.com/aguncn/nezha/models"

//IApplicationRepository Application接口定义
type IApplicationRepository interface {
	//GetApplication 根据id获取Application
	GetApplication(where interface{}) *models.Application
	//AddApplication 新增Application
	AddApplication(application *models.Application) bool
	//UpdateApplication 更新Application
	UpdateApplication(application *models.Application) bool
	//GetApplications 获取Application
	GetApplications(PageNum, PageSize uint, total *uint64, where interface{}) *[]models.Application
	//RecentApplications 获取最近Application
	RecentApplications(where interface{}, limit uint) *[]models.Application
	//ExistApplicationByName 是否存在已有应用
	ExistApplicationByName(where interface{}) bool
	//DeleteApplication 删除已有应用
	DeleteApplication(id uint) bool
}
