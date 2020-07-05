package controller

import (
	"fmt"

	"net/http"
	"strconv"

	"github.com/aguncn/nezha/common/codes"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	"github.com/aguncn/nezha/page"
	"github.com/aguncn/nezha/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//### 如果是使用Go Module,gin-jwt模块应使用v2
//下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
//导入应写成 import "github.com/appleboy/gin-jwt/v2"
//### 如果不是使用Go Module
//下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
//导入import "github.com/appleboy/gin-jwt"

//User 注入IUserService
type User struct {
	Log     logger.ILogger       `inject:""`
	Service service.IUserService `inject:""`
}

//GetUserInfo 根据token获取用户信息
func (a *User) GetUserInfo(c *gin.Context) {
	userRole := jwt.ExtractClaims(c)
	userName := userRole["userName"].(string)
	userID := uint(userRole["userID"].(float64))
	userType := uint(userRole["userType"].(float64))
	code := codes.SUCCESS
	data := page.User{UserType: userType, Introduction: "", UserID: userID, UserName: userName}
	RespData(c, http.StatusOK, code, &data)
}

//Logout 退出登录
func (a *User) Logout(c *gin.Context) {
	RespOk(c, http.StatusOK, codes.SUCCESS)
}

//GetUsers 获取用户信息
func (a *User) GetUsers(c *gin.Context) {
	var maps string
	code := codes.SUCCESS
	name := c.Query("name")
	if name != "" {
		maps = "username LIKE '%" + name + "%'"
	}
	page, pagesize := GetPage(c)
	data := a.Service.GetUsers(page, pagesize, maps)
	RespData(c, http.StatusOK, code, data)
}

//RegisterUser 注册用户
func (a *User) RegisterUser(c *gin.Context) {
	user := models.User{}
	code := codes.InvalidParams
	err := c.Bind(&user)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(user.Username, "username").Message("用户名不能为空")
		valid.Required(user.Password, "password").Message("密码不能为空")
		if !valid.HasErrors() {
			user.CreatedBy = user.Username
			// 1- noral; 2 - forbidden
			user.State = 1
			// 1 - admin; 2 - dev
			user.UserType = 2
			if !a.Service.ExistUserByName(user.Username) {
				if a.Service.AddUser(&user) {
					code = codes.SUCCESS
				} else {
					code = codes.ERROR
				}
			} else {
				code = codes.ErrExistUser
			}
		} else {
			for _, err := range valid.Errors {
				a.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}

	RespOk(c, http.StatusOK, code)
}

//AddUser 新建用户
func (a *User) AddUser(c *gin.Context) {
	user := models.User{}
	code := codes.InvalidParams
	err := c.Bind(&user)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(user.Username, "username").Message("用户名不能为空")
		valid.Required(user.Password, "password").Message("密码不能为空")
		valid.Required(user.UserType, "user_type").Message("用户类型不能为空")
		if !valid.HasErrors() {
			userRole := jwt.ExtractClaims(c)
			userName := userRole["userName"].(string)
			createdBy := userName
			user.CreatedBy = createdBy
			user.State = 1
			if !a.Service.ExistUserByName(user.Username) {
				if a.Service.AddUser(&user) {
					code = codes.SUCCESS
				} else {
					code = codes.ERROR
				}
			} else {
				code = codes.ErrExistUser
			}
		} else {
			for _, err := range valid.Errors {
				a.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}

	RespOk(c, http.StatusOK, code)
}

//UpdateUser 修改用户
func (a *User) UpdateUser(c *gin.Context) {
	code := codes.InvalidParams
	type UserOpType struct {
		OpType string `json:"op_type"`
	}
	userOpType := UserOpType{}

	if errA := c.ShouldBindBodyWith(&userOpType, binding.JSON); errA != nil {
		fmt.Println("ShouldBindBodyWith userOpType error")
	}
	user := models.User{}
	if errB := c.ShouldBindBodyWith(&user, binding.JSON); errB != nil {
		fmt.Println("ShouldBindBodyWith user error")
	}

	valid := validation.Validation{}
	valid.Required(user.Username, "username").Message("用户名不能为空")
	if !valid.HasErrors() {
		roles := jwt.ExtractClaims(c)
		UpdatedBy := roles["userName"].(string)
		user.UpdatedBy = UpdatedBy
		if a.Service.UpdateUser(&user, userOpType.OpType) {
			code = codes.SUCCESS
		} else {
			code = codes.ERROR
		}
	} else {
		for _, err := range valid.Errors {
			a.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	RespOk(c, http.StatusOK, code)
}

//DeleteUser 删除用户
func (a *User) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := codes.SUCCESS
	if !a.Service.DeleteUser(uint(id)) {
		code = codes.ERROR
		RespFail(c, http.StatusOK, code, "不允许删除admin账号!")
	} else {
		RespOk(c, http.StatusOK, code)
	}
}
