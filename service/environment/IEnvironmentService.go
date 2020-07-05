package environment

import (
	"github.com/aguncn/nezha/models"
)

//IEnvironmentService Environment接口定义
type IEnvironmentService interface {
	GetEnvironments(page, pagesize uint, maps interface{}) interface{}
	GetEnvironment(id uint) *models.Environment
	AddEnvironment(Environment *models.Environment) bool
	UpdateEnvironment(Environment *models.Environment) bool
	DeleteEnvironment(id uint) bool
}
