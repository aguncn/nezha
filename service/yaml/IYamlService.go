package yaml

import (
	"github.com/aguncn/nezha/models"
)

type IYamlService interface {
	GetYaml(id uint) *models.Yaml

	AddYaml(yaml *models.Yaml) bool

	UpdateYaml(yaml *models.Yaml) bool

	GetYamls(page, pagesize uint, maps interface{}) interface{}

	DeleteYaml(id uint) bool
}
