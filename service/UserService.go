package service

import (
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	pageModel "github.com/aguncn/nezha/page"
	"github.com/aguncn/nezha/page/emun"
	"github.com/aguncn/nezha/repository"
)

// UserService 注入IUserRepository
type UserService struct {
	Repository repository.IUserRepository `inject:""`
	Log        logger.ILogger             `inject:""`
}

//CheckUser 身份验证
func (a *UserService) CheckUser(username string, password string) bool {
	where := models.User{Username: username, Password: password}
	return a.Repository.CheckUser(where)
}

//GetUserID 获取用户ID
func (a *UserService) GetUserID(username string) uint {
	userWhere := models.User{Username: username}
	userSel := "id"
	userID := a.Repository.GetUserID(&userSel, &userWhere)
	return userID
}

//GetUserType 获取用户Type(管理员，研发)
func (a *UserService) GetUserType(username string) uint {
	userWhere := models.User{Username: username}
	userSel := "user_type"
	UserType := a.Repository.GetUserType(&userSel, &userWhere)
	return UserType
}

//GetUsers 获取用户信息
func (a *UserService) GetUsers(page, pagesize uint, maps interface{}) interface{} {
	res := make(map[string]interface{}, 2)
	var total uint64
	users := a.Repository.GetUsers(page, pagesize, &total, maps)
	var pageUsers []pageModel.Users
	var pageUser pageModel.Users
	for _, user := range *users {
		pageUser.ID = user.ID
		pageUser.Name = user.Username
		pageUser.Password = user.Password
		pageUser.UserType = emun.GetUserType(user.UserType)
		pageUser.State = emun.GetStatus(user.State)
		pageUser.Deleted = emun.GetDeleted(user.Deleted)
		pageUser.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
		pageUsers = append(pageUsers, pageUser)
	}
	res["list"] = &pageUsers
	res["total"] = total
	return &res
}

//RegisterUser 新建用户，并不同时新建用户角色
func (a *UserService) RegisterUser(user *models.User) bool {

	isOK := a.Repository.AddUser(user)
	if !isOK {
		return false
	}
	return true
}

//AddUser 新建用户
func (a *UserService) AddUser(user *models.User) bool {

	isOK := a.Repository.AddUser(user)
	if !isOK {
		return false
	}
	return true
}

//ExistUserByName 判断用户名是否已存在
func (a *UserService) ExistUserByName(username string) bool {
	where := models.User{Username: username}
	return a.Repository.ExistUserByName(where)
}

//UpdateUser 更新用户
func (a *UserService) UpdateUser(modUser *models.User, userOpType string) bool {
	user := a.Repository.GetUserByID(modUser.ID)
	//不允许更新用户名
	// user.Username = modUser.Username
	user.Password = modUser.Password
	user.UpdatedBy = modUser.UpdatedBy
	user.UserType = modUser.UserType
	return a.Repository.UpdateUser(user, userOpType)
}

//DeleteUser 删除用户
func (a *UserService) DeleteUser(id uint) bool {
	user := a.Repository.GetUserByID(id)
	if user.Username == "admin" {
		a.Log.Errorf("删除用户失败:不能删除admin账号")
		return false
	}
	return a.Repository.DeleteUser(id)
}
