package environment

import (
	"github.com/aguncn/nezha/models"
	envRep "github.com/aguncn/nezha/repository/environment"
	"github.com/jinzhu/gorm"
)

// ArticleService 注入IArticleRepo
type EnvironmentService struct {
	Repository envRep.IEnvironmentRepository `inject:""`
}

//GetEnvironments 获取Environment信息
func (environment *EnvironmentService) GetEnvironments(page, pagesize uint, maps interface{}) interface{} {
	res := make(map[string]interface{}, 2)
	var total uint64
	Environments := environment.Repository.GetEnvironments(page, pagesize, &total, maps)
	res["list"] = &Environments
	res["total"] = total
	return &res
}

func (environment *EnvironmentService) GetEnvironment(id uint) *models.Environment {
	where := models.Environment{Model: gorm.Model{ID: id}}
	return environment.Repository.GetEnvironment(&where)
}

func (environment *EnvironmentService) AddEnvironment(Environment *models.Environment) bool {

	return environment.Repository.AddEnvironment(Environment)

}

func (environment *EnvironmentService) UpdateEnvironment(Environment *models.Environment) bool {

	return environment.Repository.UpdateEnvironment(Environment)

}

func (environment *EnvironmentService) DeleteEnvironment(id uint) bool {
	return environment.Repository.DeleteEnvironment(id)
}
