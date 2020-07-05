package pod

import (
	"github.com/aguncn/nezha/models"
	//pageModel "github.com/aguncn/nezha/page"
	podRep "github.com/aguncn/nezha/repository/pod"
	"github.com/jinzhu/gorm"
)

type PodService struct {
	Repository podRep.IPodRepository `inject:""`
}

func (pod *PodService) GetPod(id uint) *models.Pod {
	where := models.Pod{Model: gorm.Model{ID: id}}
	return pod.Repository.GetPod(&where)
}

func (pod *PodService) AddPod(modPod *models.Pod) bool {
	if pod.Repository.ExistPodByName(modPod) {
		return false
	} else {
		return pod.Repository.AddPod(modPod)
	}
}

func (pod *PodService) UpdatePod(modPod *models.Pod) bool {

	return pod.Repository.UpdatePod(modPod)

}

func (pod *PodService) GetPods(page, pagesize uint, maps interface{}) interface{} {
	res := make(map[string]interface{}, 2)
	var total uint64
	Pods := pod.Repository.GetPods(page, pagesize, &total, maps)
	res["list"] = &Pods
	res["total"] = total
	return &res
}

func (pod *PodService) DeletePod(id uint) bool {
	return pod.Repository.DeletePod(id)
}

func (pod *PodService) DeletePodByCol(modPod *models.Pod) bool {
	return pod.Repository.DeletePodByCol(modPod)
}
