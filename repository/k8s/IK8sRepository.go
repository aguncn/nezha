package k8s

import "github.com/aguncn/nezha/models"

type IK8sRepository interface {
	GetK8s(where interface{}) *models.K8s

	AddK8s(k8s *models.K8s) bool

	UpdateK8s(k8s *models.K8s) bool

	GetK8ss(PageNum, PageSize uint, total *uint64, where interface{}) *[]models.K8s

	ExistK8sByName(where interface{}) bool

	DeleteK8s(id uint) bool
}
