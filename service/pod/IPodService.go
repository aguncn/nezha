package pod

import (
	"github.com/aguncn/nezha/models"
)

type IPodService interface {
	GetPod(id uint) *models.Pod

	AddPod(Pod *models.Pod) bool

	UpdatePod(Pod *models.Pod) bool

	GetPods(page, pagesize uint, maps interface{}) interface{}

	DeletePod(id uint) bool

	DeletePodByCol(Pod *models.Pod) bool
}
