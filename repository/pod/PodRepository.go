package pod

import (
	"github.com/aguncn/nezha/common/datasource"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	baseRep "github.com/aguncn/nezha/repository"
	"github.com/jinzhu/gorm"
)

type PodRepository struct {
	Log    logger.ILogger         `inject:""`
	Base   baseRep.BaseRepository `inject:"inline"`
	Source datasource.IDb         `inject:""`
}

func (pod *PodRepository) GetPod(where interface{}) *models.Pod {
	var modPod models.Pod
	other := map[string]string{
		"foreignkey": " Environment ",
	}
	if err := pod.Base.First(where, &modPod, other); err != nil {
		pod.Log.Errorf("未找到相关文章", err)
	}
	return &modPod
}

func (pod *PodRepository) AddPod(modPod *models.Pod) bool {
	if err := pod.Base.Create(modPod); err != nil {
		pod.Log.Errorf("添加Pod失败", err)
		return false
	}
	return true
}

func (pod *PodRepository) UpdatePod(modPod *models.Pod) bool {
	if err := pod.Base.Save(modPod); err != nil {
		pod.Log.Errorf("更新应用失败", err)
		return false
	}
	return true
}

func (pod *PodRepository) GetPods(PageNum uint, PageSize uint, total *uint64, where interface{}) *[]models.Pod {
	var modPods []models.Pod
	other := map[string]string{
		"order":      " ID  desc ",
		"foreignkey": " Environment, K8s ",
	}
	err := pod.Base.GetPages(&models.Pod{}, &modPods, PageNum, PageSize, total, where, other)
	if err != nil {
		pod.Log.Errorf("获取文章信息失败", err)
	}
	return &modPods
}

func (pod *PodRepository) ExistPodByName(where interface{}) bool {
	var modPod models.Pod
	var wherePod = models.Pod{PodName: where.(*models.Pod).PodName}
	other := map[string]string{}
	err := pod.Base.First(wherePod, &modPod, other)
	//记录不存在错误(RecordNotFound)，返回false
	if gorm.IsRecordNotFoundError(err) {
		return false
	}
	//其他类型的错误，写下日志，返回false
	if err != nil {
		pod.Log.Error(err)
		return false
	}
	return true
}

func (pod *PodRepository) DeletePod(id uint) bool {
	var modPod models.Pod
	if err := pod.Base.DeleteByID(modPod, id); err != nil {
		pod.Log.Errorf("删除pod失败", err)
		return false
	}
	return true
}

func (pod *PodRepository) DeletePodByCol(modPod *models.Pod) bool {
	if err := pod.Source.DB().Where(modPod).Delete(models.Pod{}).Error; err != nil {
		pod.Log.Errorf("更新POD失败", err)
		return false
	}
	return true
}
