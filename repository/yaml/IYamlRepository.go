package yaml

import "github.com/aguncn/nezha/models"

type IYamlRepository interface {
	GetYaml(where interface{}) *models.Yaml

	AddYaml(yaml *models.Yaml) bool

	UpdateYaml(yaml *models.Yaml) bool

	GetYamls(PageNum, PageSize uint, total *uint64, where interface{}) *[]models.Yaml

	ExistYamlByName(where interface{}) bool

	DeleteYaml(id uint) bool
}
