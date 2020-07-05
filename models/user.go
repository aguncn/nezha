package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User 用户授权信息
type User struct {
	gorm.Model
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	Deleted   uint   `json:"deteled"`
	State     uint   `json:"state" gorm:"default:1"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	UserType  uint   `json:"user_type"`

	Application *[]Application
}

//BeforeCreate CreatedOn赋值
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now())
	return nil
}

func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
	if pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err == nil {
		scope.SetColumn("Password", pw)
	}
	return nil
}

//BeforeUpdate ModifiedOn赋值
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now())
	return nil
}

// UserRole 用户身份结构体
type UserRole struct {
	UserID   uint
	UserName string
	UserType uint
}
