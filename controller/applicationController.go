package controller

import (
	"fmt"
	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/aguncn/nezha/common/codes"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	appSvc "github.com/aguncn/nezha/service/application"
)

// Application 注入IApplicationService
type Application struct {
	Log     logger.ILogger             `inject:""`
	Service appSvc.IApplicationService `inject:""`
}

//GetApplication 获取单个Application
func (a *Application) GetApplication(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	var data *models.Application
	code := codes.InvalidParams
	if !valid.HasErrors() {
		data = a.Service.GetApplication(uint(id))
		code = codes.SUCCESS
	} else {
		for _, err := range valid.Errors {
			a.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	RespData(c, http.StatusOK, code, data)
}

//AddApplication 新增Application
func (a *Application) AddApplication(c *gin.Context) {

	application := models.Application{}

	code := codes.InvalidParams
	err := c.Bind(&application)
	fmt.Printf("application: %v", application)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(application.Name, "name").Message("名称不能为空")
		valid.Required(application.CnName, "cn_name").Message("中文名称不能为空")
		valid.Required(application.Description, "description").Message("简述不能为空")
		valid.Required(application.Harbor, "harbor").Message("Harbor不能为空")
		valid.Required(application.CreatedBy, "created_by").Message("创建人不能为空")
		if !valid.HasErrors() {
			if a.Service.AddApplication(&application) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				a.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

//UpdateApplication 修改Application
func (a *Application) UpdateApplication(c *gin.Context) {
	applicaion := models.Application{}
	code := codes.InvalidParams
	err := c.Bind(&applicaion)
	fmt.Printf("application: %v", applicaion)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(applicaion.Harbor, "harbor").Message("Harbor不能为空")
		if !valid.HasErrors() {
			if a.Service.UpdateApplication(&applicaion) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				a.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

//GetApplications 获取Applications信息
func (a *Application) GetApplications(c *gin.Context) {

	code := codes.SUCCESS
	name := c.Query("name")
	var maps string
	isAdmin := false

	userRole := jwt.ExtractClaims(c)
	//userName := userRole["userName"].(string)
	userID := int(userRole["userID"].(float64))
	userType := int(userRole["userType"].(float64))
	if userType == 1 {
		isAdmin = true
	}

	if isAdmin {
		if name != "" {
			maps = "name LIKE '%" + name + "%'"
		}
	} else {
		if name != "" {
			maps = "name LIKE '%" + name + "%' AND user_id = " + strconv.Itoa(userID)
		} else {
			maps = "user_id = " + strconv.Itoa(userID)
		}
	}

	page, pagesize := GetPage(c)
	data := a.Service.GetApplications(page, pagesize, maps)
	RespData(c, http.StatusOK, code, data)
}

//DeleteApplication 删除Applications
func (a *Application) DeleteApplication(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := codes.SUCCESS
	if !a.Service.DeleteApplication(uint(id)) {
		code = codes.ERROR
		RespFail(c, http.StatusOK, code, "删除出错，请联系管理员!")
	} else {
		RespOk(c, http.StatusOK, code)
	}
}

//RecentApplications 获取最近有部署的Applications信息
func (a *Application) RecentApplications(c *gin.Context) {

	code := codes.SUCCESS
	tmpLimit, _ := strconv.ParseUint(c.Query("limit"), 10, 64)
	limit := uint(tmpLimit)
	var maps string
	isAdmin := false

	userRole := jwt.ExtractClaims(c)
	//userName := userRole["userName"].(string)
	userID := uint(userRole["userID"].(float64))
	userType := uint(userRole["userType"].(float64))
	if userType == 1 {
		isAdmin = true
	}

	if isAdmin == false {
		maps = "user_id = " + string(userID)
	}

	data := a.Service.RecentApplications(maps, limit)
	RespData(c, http.StatusOK, code, data)
}
