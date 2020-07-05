package k8s

import (
	"github.com/aguncn/nezha/models"
	//pageModel "github.com/aguncn/nezha/page"
	k8sRep "github.com/aguncn/nezha/repository/k8s"
	"github.com/jinzhu/gorm"
)

type K8sService struct {
	Repository k8sRep.IK8sRepository `inject:""`
}

func (k8s *K8sService) GetK8s(id uint) *models.K8s {
	where := models.K8s{Model: gorm.Model{ID: id}}
	return k8s.Repository.GetK8s(&where)
}

func (k8s *K8sService) AddK8s(modK8s *models.K8s) bool {
	if k8s.Repository.ExistK8sByName(modK8s) {
		return false
	} else {
		return k8s.Repository.AddK8s(modK8s)
	}
}

func (k8s *K8sService) UpdateK8s(modK8s *models.K8s) bool {

	return k8s.Repository.UpdateK8s(modK8s)

}

func (k8s *K8sService) GetK8ss(page, pagesize uint, maps interface{}) interface{} {
	res := make(map[string]interface{}, 2)
	var total uint64
	K8ss := k8s.Repository.GetK8ss(page, pagesize, &total, maps)
	/*
		可用于定制输出
		var pageProjects []pageModel.Project
		var pageProject pageModel.Project
		for _, Project := range *Projects {
			pageProject.ID = Project.ID
			pageProject.Name = Project.Name
			pageProject.Git = Project.Git
			pageProject.Jenkins = Project.Jenkins
			pageProject.State = Project.State
			pageProject.CreatedAt = Project.CreatedAt.Format("2006-01-02 15:04:05")
			pageProjects = append(pageProjects, pageProject)
		}
	*/
	res["list"] = &K8ss
	res["total"] = total
	return &res
}

func (k8s *K8sService) DeleteK8s(id uint) bool {
	return k8s.Repository.DeleteK8s(id)
}
