package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Release 版本单结构体
type Release struct {
	gorm.Model
	CreatedBy   string  `json:"created_by"`
	UpdatedBy   string  `json:"updated_by"`
	Deleted     uint    `json:"deteled"`
	State       uint    `json:"state" gorm:"default:1"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Branch      string  `json:"branch"`
	Yaml        string  `json:"yaml" gorm:"type:text"`
	UserID      uint    `json:"user_id"`
	ProjectID   uint    `json:"project_id"`
	User        User    `gorm:"save_associations:false"`
	Project     Project `gorm:"save_associations:false"`
}

//BeforeCreate CreatedOn赋值
func (release *Release) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Unix())

	return nil
}

//BeforeUpdate ModifiedOn赋值
func (release *Release) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())

	return nil
}
