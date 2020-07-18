package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Deploy 发布单结构体
type Deploy struct {
	gorm.Model
	CreatedBy     string      `json:"created_by"`
	UpdatedBy     string      `json:"updated_by"`
	Deleted       uint        `json:"deteled"`
	State         uint        `json:"state" gorm:"default:1"`
	Name          string      `json:"name"`
	Version       string      `json:"version"`
	Reason        string      `json:"reason"`
	Harbor        string      `json:"harbor"`
	Yaml          string      `json:"yaml" gorm:"type:text"`
	Step          uint        `json:"step"`
	Result        string      `json:"result"`
	Log           string      `json:"log" gorm:"type:text"`
	EnvironmentID uint        `json:"environment_id"`
	K8sID         uint        `json:"k8s_id"`
	UserID        uint        `json:"user_id"`
	ProjectID     uint        `json:"project_id"`
	ApplicationID uint        `json:"application_id"`
	User          User        `gorm:"save_associations:false"`
	Project       Project     `gorm:"save_associations:false"`
	Application   Application `gorm:"save_associations:false"`
	K8s           K8s         `gorm:"save_associations:false"`
	Environment   Environment `gorm:"save_associations:false"`
}

//BeforeCreate CreatedAt赋值
func (deploy *Deploy) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Unix())

	return nil
}

//BeforeUpdate UpdatedAt赋值
func (deploy *Deploy) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())

	return nil
}
