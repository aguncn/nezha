package service

import "github.com/aguncn/nezha/models"

//IUserService UserService接口定义
type IUserService interface {
	//CheckUser 身份验证
	CheckUser(username string, password string) bool
	//GetUserID 获取用户ID
	GetUserID(username string) uint
	//GetUserType 获取用户角色
	GetUserType(username string) uint
	//GetUsers 获取用户信息
	GetUsers(page, pagesize uint, maps interface{}) interface{}
	//AddUser 新建用户
	AddUser(user *models.User) bool
	//ExistUserByName 判断用户名是否已存在
	ExistUserByName(username string) bool
	//UpdateUser 更新用户
	UpdateUser(user *models.User, userOpType string) bool
	//DeleteUser 删除用户
	DeleteUser(id uint) bool
}
