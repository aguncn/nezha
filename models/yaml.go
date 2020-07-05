package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Yaml Yaml结构体
type Yaml struct {
	gorm.Model
	CreatedBy     string      `json:"created_by"`
	UpdatedBy     string      `json:"updated_by"`
	Deleted       uint        `json:"deteled"`
	State         uint        `json:"state"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Yaml          string      `json:"yaml" gorm:"type:text"`
	K8sID         uint        `json:"k8s_id"`
	ApplicationID uint        `json:"application_id"`
	UserID        uint        `json:"user_id"`
	K8s           K8s         `gorm:"save_associations:false"`
	Application   Application `gorm:"save_associations:false"`
	User          User        `gorm:"save_associations:false"`
}

//BeforeCreate CreatedOn赋值
func (yaml *Yaml) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate ModifiedOn赋值
func (yaml *Yaml) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
