package yaml

import (
	"github.com/aguncn/nezha/common/datasource"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	baseRep "github.com/aguncn/nezha/repository"
	"github.com/jinzhu/gorm"
)

type YamlRepository struct {
	Log    logger.ILogger         `inject:""`
	Base   baseRep.BaseRepository `inject:"inline"`
	Source datasource.IDb         `inject:""`
}

func (yaml *YamlRepository) GetYaml(where interface{}) *models.Yaml {
	var modYaml models.Yaml
	other := map[string]string{
		"foreignkey": " Environment ",
	}
	if err := yaml.Base.First(where, &modYaml, other); err != nil {
		yaml.Log.Errorf("未找到相关Yaml", err)
	}
	return &modYaml
}

func (yaml *YamlRepository) AddYaml(modYaml *models.Yaml) bool {
	if err := yaml.Base.Create(modYaml); err != nil {
		yaml.Log.Errorf("添加Yaml失败", err)
		return false
	}
	return true
}

func (yaml *YamlRepository) UpdateYaml(modYaml *models.Yaml) bool {
	if err := yaml.Base.Save(modYaml); err != nil {
		yaml.Log.Errorf("更新Yaml失败", err)
		return false
	}
	return true
}

func (yaml *YamlRepository) GetYamls(PageNum uint, PageSize uint, total *uint64, where interface{}) *[]models.Yaml {
	var yamls []models.Yaml
	db := yaml.Source.DB().Model(&yamls).Where(&yamls)
	db = db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,username")
	}).Preload("Application", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name")
	}).Preload("K8s", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name,environment_id")
	}).Preload("K8s.Environment", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name")
	})

	db = db.Order("updated_at desc")
	db = db.Where(where)
	err := db.Offset((PageNum - 1) * PageSize).Limit(PageSize).Find(&yamls).Error
	if err != nil {
		yaml.Log.Errorf("获取Yaml失败", err)
	}
	err = db.Count(total).Error
	if err != nil {
		yaml.Log.Errorf("查询总数出错", err)
		// return err
	}
	return &yamls

}

func (yaml *YamlRepository) ExistYamlByName(where interface{}) bool {
	var modYaml models.Yaml
	var whereYaml = models.Yaml{Name: where.(*models.Yaml).Name}
	other := map[string]string{}
	err := yaml.Base.First(whereYaml, &modYaml, other)
	//记录不存在错误(RecordNotFound)，返回false
	if gorm.IsRecordNotFoundError(err) {
		return false
	}
	//其他类型的错误，写下日志，返回false
	if err != nil {
		yaml.Log.Error(err)
		return false
	}
	return true
}

func (yaml *YamlRepository) DeleteYaml(id uint) bool {
	var modYaml models.Yaml
	if err := yaml.Base.DeleteByID(modYaml, id); err != nil {
		yaml.Log.Errorf("删除Yaml失败", err)
		return false
	}
	return true
}
