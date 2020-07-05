package deploy

import "github.com/aguncn/nezha/models"

type IDeployRepository interface {
	GetDeploy(where interface{}) *models.Deploy

	AddDeploy(deploy *models.Deploy) bool

	UpdateDeploy(deploy *models.Deploy) bool

	UpdateDeployByCol(deploy *models.Deploy, where map[string]interface{}) bool

	GetDeploys(PageNum, PageSize uint, total *uint64, where interface{}) *[]models.Deploy

	GetOkErrorCount(where string) interface{}

	GetBetweenCount(where string, cycle, count int) interface{}

	ExistDeployByName(where interface{}) bool

	DeleteDeploy(id uint) bool
}
