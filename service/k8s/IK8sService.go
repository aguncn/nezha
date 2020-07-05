package k8s

import (
	"github.com/aguncn/nezha/models"
)

type IK8sService interface {
	GetK8s(id uint) *models.K8s

	AddK8s(K8s *models.K8s) bool

	UpdateK8s(K8s *models.K8s) bool

	GetK8ss(page, pagesize uint, maps interface{}) interface{}

	DeleteK8s(id uint) bool
}
