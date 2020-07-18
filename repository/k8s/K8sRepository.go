package k8s

import (
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	baseRep "github.com/aguncn/nezha/repository"
	"github.com/jinzhu/gorm"
)

type K8sRepository struct {
	Log  logger.ILogger         `inject:""`
	Base baseRep.BaseRepository `inject:"inline"`
}

func (k8s *K8sRepository) GetK8s(where interface{}) *models.K8s {
	var modK8s models.K8s
	other := map[string]string{
		"foreignkey": " Environment ",
	}
	if err := k8s.Base.First(where, &modK8s, other); err != nil {
		k8s.Log.Errorf("未找到相关K8s", err)
	}
	return &modK8s
}

func (k8s *K8sRepository) AddK8s(modK8s *models.K8s) bool {
	if err := k8s.Base.Create(modK8s); err != nil {
		k8s.Log.Errorf("添加K8s失败", err)
		return false
	}
	return true
}

func (k8s *K8sRepository) UpdateK8s(modK8s *models.K8s) bool {
	if err := k8s.Base.Save(modK8s); err != nil {
		k8s.Log.Errorf("更新K8s失败", err)
		return false
	}
	return true
}

func (k8s *K8sRepository) GetK8ss(PageNum uint, PageSize uint, total *uint64, where interface{}) *[]models.K8s {
	var modK8ss []models.K8s
	other := map[string]string{
		"order":      " ID  desc ",
		"foreignkey": " Environment ",
	}
	err := k8s.Base.GetPages(&models.K8s{}, &modK8ss, PageNum, PageSize, total, where, other)
	if err != nil {
		k8s.Log.Errorf("获取K8s信息失败", err)
	}
	return &modK8ss
}

func (k8s *K8sRepository) ExistK8sByName(where interface{}) bool {
	var modK8s models.K8s
	var whereK8s = models.K8s{Name: where.(*models.K8s).Name}
	other := map[string]string{}
	err := k8s.Base.First(whereK8s, &modK8s, other)
	//记录不存在错误(RecordNotFound)，返回false
	if gorm.IsRecordNotFoundError(err) {
		return false
	}
	//其他类型的错误，写下日志，返回false
	if err != nil {
		k8s.Log.Error(err)
		return false
	}
	return true
}

func (k8s *K8sRepository) DeleteK8s(id uint) bool {
	var modK8s models.K8s
	if err := k8s.Base.DeleteByID(modK8s, id); err != nil {
		k8s.Log.Errorf("删除K8s失败", err)
		return false
	}
	return true
}
