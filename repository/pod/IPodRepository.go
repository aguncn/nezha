package pod

import "github.com/aguncn/nezha/models"

type IPodRepository interface {
	GetPod(where interface{}) *models.Pod

	AddPod(pod *models.Pod) bool

	UpdatePod(pod *models.Pod) bool

	GetPods(PageNum, PageSize uint, total *uint64, where interface{}) *[]models.Pod

	ExistPodByName(where interface{}) bool

	DeletePod(id uint) bool

	DeletePodByCol(pod *models.Pod) bool
}
