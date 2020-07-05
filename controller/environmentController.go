package controller

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/aguncn/nezha/common/codes"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"

	envSvc "github.com/aguncn/nezha/service/environment"
)

// Environment 注入IEnvironmentService
type Environment struct {
	Log     logger.ILogger             `inject:""`
	Service envSvc.IEnvironmentService `inject:""`
}

//GetEnvironments 获取项目信息
func (environment *Environment) GetEnvironments(c *gin.Context) {
	var maps string
	code := codes.SUCCESS
	name := c.Query("name")
	if name != "" {
		maps = "name LIKE '%" + name + "%'"
	}
	page, pagesize := GetPage(c)
	data := environment.Service.GetEnvironments(page, pagesize, maps)
	RespData(c, http.StatusOK, code, data)
}

//GetProject 获取单个Project
func (environment *Environment) GetEnvironment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	var data *models.Environment
	code := codes.InvalidParams
	if !valid.HasErrors() {
		data = environment.Service.GetEnvironment(uint(id))
		code = codes.SUCCESS
	} else {
		for _, err := range valid.Errors {
			environment.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	RespData(c, http.StatusOK, code, data)
}

//AddProject 新增Project
func (environment *Environment) AddEnvironment(c *gin.Context) {

	Environment := models.Environment{}

	code := codes.InvalidParams
	err := c.Bind(&Environment)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(Environment.Name, "name").Message("名称不能为空")
		if !valid.HasErrors() {
			if environment.Service.AddEnvironment(&Environment) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				environment.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

//UpdateProject 修改Project
func (environment *Environment) UpdateEnvironment(c *gin.Context) {
	Environment := models.Environment{}
	code := codes.InvalidParams
	err := c.Bind(&Environment)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(Environment.Name, "name").Message("Name不能为空")
		if !valid.HasErrors() {
			if environment.Service.UpdateEnvironment(&Environment) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				environment.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

//DeleteProject 删除项目
func (environment *Environment) DeleteEnvironment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := codes.SUCCESS
	if !environment.Service.DeleteEnvironment(uint(id)) {
		code = codes.ERROR
		RespFail(c, http.StatusOK, code, "删除出错，请联系管理员!")
	} else {
		RespOk(c, http.StatusOK, code)
	}
}
