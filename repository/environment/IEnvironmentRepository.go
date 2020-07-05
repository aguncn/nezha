package environment

import "github.com/aguncn/nezha/models"

//IEnvironmentRepository Environment接口定义
type IEnvironmentRepository interface {
	GetEnvironments(PageNum, PageSize uint, total *uint64, where interface{}) *[]models.Environment

	GetEnvironment(where interface{}) *models.Environment

	AddEnvironment(Environment *models.Environment) bool

	UpdateEnvironment(Environment *models.Environment) bool

	DeleteEnvironment(id uint) bool
}
