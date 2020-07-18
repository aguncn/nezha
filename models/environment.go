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

//BeforeCreate CreatedAt赋值
func (environment *Environment) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Unix())

	return nil
}

//BeforeUpdate UpdateAt赋值
func (environment *Environment) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateAt", time.Now().Unix())

	return nil
}
