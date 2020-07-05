package environment

import (
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	baseRep "github.com/aguncn/nezha/repository"
)

//EnvironmentRepository 注入IDb
type EnvironmentRepository struct {
	Log  logger.ILogger         `inject:""`
	Base baseRep.BaseRepository `inject:"inline"`
}

//GetEnvironments 获取环境列表
func (environment *EnvironmentRepository) GetEnvironments(PageNum uint, PageSize uint, total *uint64, where interface{}) *[]models.Environment {
	var Environments []models.Environment
	other := map[string]string{
		"order": "ID desc",
	}
	err := environment.Base.GetPages(&models.Environment{}, &Environments, PageNum, PageSize, total, where, other)
	if err != nil {
		environment.Log.Errorf("获取环境失败", err)
	}
	return &Environments
}

func (environment *EnvironmentRepository) GetEnvironment(where interface{}) *models.Environment {
	var modEnvironment models.Environment
	other := map[string]string{}
	if err := environment.Base.First(where, &modEnvironment, other); err != nil {
		environment.Log.Errorf("未找到相关环境", err)
	}
	return &modEnvironment
}

func (environment *EnvironmentRepository) AddEnvironment(Environment *models.Environment) bool {
	if err := environment.Base.Create(Environment); err != nil {
		environment.Log.Errorf("添加环境失败", err)
		return false
	}
	return true
}

func (environment *EnvironmentRepository) UpdateEnvironment(Environment *models.Environment) bool {
	if err := environment.Base.Save(Environment); err != nil {
		environment.Log.Errorf("更新环境失败", err)
		return false
	}
	return true
}

func (environment *EnvironmentRepository) DeleteEnvironment(id uint) bool {
	var modEnvironment models.Environment
	if err := environment.Base.DeleteByID(modEnvironment, id); err != nil {
		environment.Log.Errorf("删除环境失败", err)
		return false
	}
	return true
}
