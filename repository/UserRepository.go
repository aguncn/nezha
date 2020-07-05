package repository

import (
	"fmt"

	"github.com/aguncn/nezha/common/datasource"

	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//UserRepository 注入IDb
type UserRepository struct {
	Log    logger.ILogger `inject:""`
	Base   BaseRepository `inject:"inline"`
	Source datasource.IDb `inject:""`
}

//CheckUser 身份验证
func (a *UserRepository) CheckUser(where interface{}) bool {
	var user models.User
	var whereUser = models.User{Username: where.(models.User).Username}
	other := map[string]string{}
	if err := a.Base.First(whereUser, &user, other); err != nil {
		a.Log.Errorf("用户名错误", err)
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(where.(models.User).Password))
	if err != nil {
		a.Log.Errorf("密码错误", err)
		return false
	}
	return true
}

//GetUserID 获取用户ID
func (a *UserRepository) GetUserID(sel *string, where interface{}) uint {
	var user models.User
	other := map[string]string{
		"selects": *sel,
	}
	if err := a.Base.First(where, &user, other); err != nil {
		a.Log.Errorf("获取用户ID失败", err)
		return 0
	}
	return user.ID
}

//GetUserType 获取用户Type
func (a *UserRepository) GetUserType(sel *string, where interface{}) uint {
	var user models.User
	other := map[string]string{
		"selects": *sel,
	}
	if err := a.Base.First(where, &user, other); err != nil {
		a.Log.Errorf("获取用户ID失败", err)
		return 0
	}
	return user.UserType
}

//GetUsers 获取用户信息
func (a *UserRepository) GetUsers(PageNum uint, PageSize uint, total *uint64, where interface{}) *[]models.User {
	var users []models.User
	other := map[string]string{}
	if err := a.Base.GetPages(&models.User{}, &users, PageNum, PageSize, total, where, other); err != nil {
		a.Log.Errorf("获取用户信息失败", err)
	}
	return &users
}

//AddUser 新建用户
func (a *UserRepository) AddUser(user *models.User) bool {
	if err := a.Base.Create(user); err != nil {
		a.Log.Errorf("新建用户失败", err)
		return false
	}
	return true
}

//ExistUserByName 判断用户名是否已存在
func (a *UserRepository) ExistUserByName(where interface{}) bool {
	var user models.User
	other := map[string]string{
		"selects": "id",
	}
	err := a.Base.First(where, &user, other)
	//记录不存在错误(RecordNotFound)，返回false
	if gorm.IsRecordNotFoundError(err) {
		return false
	}
	//其他类型的错误，写下日志，返回false
	if err != nil {
		a.Log.Error(err)
		return false
	}
	return true
}

//UpdateUser 更新用户
func (a *UserRepository) UpdateUser(user *models.User, userOpType string) bool {
	//使用事务更新用户数据
	fmt.Println(userOpType, user.Password, user.UserType, "===db==============")
	db := a.Source.DB().Model(&user)
	if userOpType == "adminResetPwd" {
		db.Update("password", user.Password)
		return true
	} else if userOpType == "adminChangePermission" {
		db.UpdateColumn("user_type", user.UserType)
		return true
	} else if userOpType == "userResetPwd" {
		db.Update("password", user.Password)
		return true
	} else {
		return false
	}
	return false
}

//DeleteUser 删除用户同时删除用户的角色
func (a *UserRepository) DeleteUser(id uint) bool {
	//采用事务同时删除用户和相应的用户角色
	var (
		userWhere = models.User{Model: gorm.Model{ID: id}}
		user      models.User
	)
	tx := a.Base.GetTransaction()
	if err := tx.Where(&userWhere).Delete(&user).Error; err != nil {
		a.Log.Errorf("删除用户失败", err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//GetUserByID 获取用户
func (a *UserRepository) GetUserByID(id uint) *models.User {
	var user models.User
	if err := a.Base.FirstByID(&user, id); err != nil {
		a.Log.Error(err)
	}
	return &user
}
