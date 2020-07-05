package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//K8s K8s集群结构体
type K8s struct {
	gorm.Model
	CreatedBy     string      `json:"created_by"`
	UpdatedBy     string      `json:"updated_by"`
	Deleted       uint        `json:"deteled"`
	State         uint        `json:"state"`
	Name          string      `json:"name"`
	EnvironmentID uint        `json:"environment_id"`
	Conf          string      `json:"conf" gorm:"type:text"`
	Terminal      string      `json:"terminal"`
	UserID        uint        `json:"user_id"`
	Environment   Environment `gorm:"save_associations:false"`
	User          User        `gorm:"save_associations:false"`
}

//BeforeCreate CreatedOn赋值
func (k8s *K8s) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate ModifiedOn赋值
func (k8s *K8s) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
