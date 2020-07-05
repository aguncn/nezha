package yaml

import (
	"github.com/aguncn/nezha/models"
	//pageModel "github.com/aguncn/nezha/page"
	yamlRep "github.com/aguncn/nezha/repository/yaml"
	"github.com/jinzhu/gorm"
)

type YamlService struct {
	Repository yamlRep.IYamlRepository `inject:""`
}

func (yaml *YamlService) GetYaml(id uint) *models.Yaml {
	where := models.Yaml{Model: gorm.Model{ID: id}}
	return yaml.Repository.GetYaml(&where)
}

func (yaml *YamlService) AddYaml(modYaml *models.Yaml) bool {
	if yaml.Repository.ExistYamlByName(modYaml) {
		return false
	} else {
		return yaml.Repository.AddYaml(modYaml)
	}
}

func (yaml *YamlService) UpdateYaml(modYaml *models.Yaml) bool {

	return yaml.Repository.UpdateYaml(modYaml)

}

func (yaml *YamlService) GetYamls(page, pagesize uint, maps interface{}) interface{} {
	res := make(map[string]interface{}, 2)
	var total uint64
	Yamls := yaml.Repository.GetYamls(page, pagesize, &total, maps)
	res["list"] = &Yamls
	res["total"] = total
	return &res
}

func (yaml *YamlService) DeleteYaml(id uint) bool {
	return yaml.Repository.DeleteYaml(id)
}
