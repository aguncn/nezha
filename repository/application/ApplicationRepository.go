package application

import (
	"github.com/aguncn/nezha/common/datasource"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	baseRep "github.com/aguncn/nezha/repository"
	"github.com/jinzhu/gorm"
)

//ApplicationRepository 注入IDb
type ApplicationRepository struct {
	Log    logger.ILogger         `inject:""`
	Base   baseRep.BaseRepository `inject:"inline"`
	Source datasource.IDb         `inject:""`
}

//GetApplication 根据id获取Application
func (a *ApplicationRepository) GetApplication(where interface{}) *models.Application {
	var application models.Application

	db := a.Source.DB().Where(where)
	db = db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,username")
	}).Preload("Project", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name,cn_name, description")
	})
	if err := db.First(&application).Error; err != nil {
		a.Log.Errorf("未找到相关Application", err)
	}
	return &application
}

//AddApplication 新增Application
func (a *ApplicationRepository) AddApplication(application *models.Application) bool {
	if err := a.Base.Create(application); err != nil {
		a.Log.Errorf("添加文章失败", err)
		return false
	}
	return true
}

//UpdateApplication 更新Application
func (a *ApplicationRepository) UpdateApplication(application *models.Application) bool {
	if err := a.Base.Save(application); err != nil {
		a.Log.Errorf("更新应用失败", err)
		return false
	}
	return true
}

//GetApplications 获取Applications
func (a *ApplicationRepository) GetApplications(PageNum uint, PageSize uint, total *uint64, where interface{}) *[]models.Application {
	var applications []models.Application
	db := a.Source.DB().Model(&applications).Where(&applications)
	db = db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,username")
	}).Preload("Project", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name,cn_name, description")
	})
	db = db.Order("updated_at  desc")
	db = db.Where(where)
	err := db.Offset((PageNum - 1) * PageSize).Limit(PageSize).Find(&applications).Error
	if err != nil {
		a.Log.Errorf("获取Applications失败", err)
	}
	err = db.Count(total).Error
	if err != nil {
		a.Log.Errorf("查询总数出错", err)
		// return err
	}
	return &applications
}

//RecentApplications 获取Applications
func (a *ApplicationRepository) RecentApplications(where interface{}, limit uint) *[]models.Application {
	var applications []models.Application
	db := a.Source.DB().Model(&applications).Where(&applications)
	db = db.Order("updated_at  desc")
	db = db.Where(where)
	err := db.Limit(limit).Find(&applications).Error
	if err != nil {
		a.Log.Errorf("获取Applications失败", err)
	}
	return &applications
}

//ExistApplicationByName 判断用户名是否已存在
func (a *ApplicationRepository) ExistApplicationByName(where interface{}) bool {
	var application models.Application
	var whereApplication = models.Application{Name: where.(*models.Application).Name}
	other := map[string]string{}
	err := a.Base.First(whereApplication, &application, other)
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

//DeleteApplication 删除应用
func (a *ApplicationRepository) DeleteApplication(id uint) bool {
	var application models.Application
	if err := a.Base.DeleteByID(application, id); err != nil {
		a.Log.Errorf("删除应用失败", err)
		return false
	}
	return true
}
