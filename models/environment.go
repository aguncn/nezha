package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Environment 环境列表
type Environment struct {
	gorm.Model
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	Deleted     uint   `json:"deteled"`
	State       uint   `json:"state"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

//BeforeCreate CreatedOn赋值
func (environment *Environment) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate ModifiedOn赋值
func (environment *Environment) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
