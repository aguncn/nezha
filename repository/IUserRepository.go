package repository

import "github.com/aguncn/nezha/models"

//IUserRepository User接口定义
type IUserRepository interface {
	//CheckUser 身份验证
	CheckUser(where interface{}) bool
	//GetUserID 获取用户ID
	GetUserID(sel *string, where interface{}) uint
	//GetUserType 获取用户Type
	GetUserType(sel *string, where interface{}) uint
	//GetUsers 获取用户信息
	GetUsers(PageNum uint, PageSize uint, total *uint64, where interface{}) *[]models.User
	//AddUser 新建用户
	AddUser(user *models.User) bool
	//ExistUserByName 判断用户名是否已存在
	ExistUserByName(where interface{}) bool
	//UpdateUser 更新用户
	UpdateUser(user *models.User, userOpType string) bool
	//DeleteUser 更新用户
	DeleteUser(id uint) bool
	//GetUserByID 获取用户
	GetUserByID(id uint) *models.User
}
