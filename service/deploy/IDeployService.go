package deploy

import (
	"github.com/aguncn/nezha/models"
)

type IDeployService interface {
	GetDeploy(id uint) *models.Deploy

	AddDeploy(Deploy *models.Deploy) bool

	UpdateDeploy(Deploy *models.Deploy) bool

	GetDeploys(page, pagesize uint, maps interface{}) interface{}

	DeleteDeploy(id uint) bool

	SubmitDeploy(Deploy *models.Deploy) (bool, string, string)

	StatusDeploy(Deploy *models.Deploy) (bool, string, string)

	GetDockerTag(Harbor string) (bool, []models.DockerTag)

	GetOkErrorCount(maps string) interface{}

	GetBetweenCount(maps string, cycle, count int) interface{}
}
