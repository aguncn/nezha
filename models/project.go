package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Project 项目结构体
type Project struct {
	gorm.Model
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	Deleted     uint   `json:"deteled"`
	State       uint   `json:"state"`
	Name        string `json:"name"`
	CnName      string `json:"cn_name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

//BeforeCreate CreatedOn赋值
func (project *Project) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate ModifiedOn赋值
func (project *Project) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
