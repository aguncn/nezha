package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Pod Pod在线列表
type Pod struct {
	gorm.Model
	CreatedBy     string      `json:"created_by"`
	UpdatedBy     string      `json:"updated_by"`
	Deleted       uint        `json:"deteled"`
	State         uint        `json:"state"`
	PodName       string      `json:"podname"`
	NameSpace     string      `json:"namespace"`
	Container     string      `json:"container"`
	Image         string      `json:"image"`
	Version       string      `json:"version"`
	EnvironmentID uint        `json:"environment_id"`
	K8sID         uint        `json:"k8s_id"`
	UserID        uint        `json:"user_id"`
	ApplicationID uint        `json:"application_id"`
	User          User        `gorm:"save_associations:false"`
	Application   Application `gorm:"save_associations:false"`
	Environment   Environment `gorm:"save_associations:false"`
	K8s           K8s         `gorm:"save_associations:false"`
}

//BeforeCreate CreatedOn赋值
func (pod *Pod) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate ModifiedOn赋值
func (pod *Pod) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
