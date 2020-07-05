package controller

import (
	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/aguncn/nezha/common/codes"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	yamlSvc "github.com/aguncn/nezha/service/yaml"
)

type Yaml struct {
	Log     logger.ILogger       `inject:""`
	Service yamlSvc.IYamlService `inject:""`
}

func (yaml *Yaml) GetYaml(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	var data *models.Yaml
	code := codes.InvalidParams
	if !valid.HasErrors() {
		data = yaml.Service.GetYaml(uint(id))
		code = codes.SUCCESS
	} else {
		for _, err := range valid.Errors {
			yaml.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	RespData(c, http.StatusOK, code, data)
}

func (yaml *Yaml) AddYaml(c *gin.Context) {
	modYaml := models.Yaml{}
	code := codes.InvalidParams
	err := c.Bind(&modYaml)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(modYaml.Name, "name").Message("名称不能为空")
		valid.Required(modYaml.CreatedBy, "created_by").Message("创建人不能为空")
		if !valid.HasErrors() {
			if yaml.Service.AddYaml(&modYaml) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				yaml.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

func (yaml *Yaml) UpdateYaml(c *gin.Context) {
	modYaml := models.Yaml{}
	code := codes.InvalidParams
	err := c.Bind(&modYaml)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(modYaml.Name, "name").Message("Name不能为空")
		if !valid.HasErrors() {
			if yaml.Service.UpdateYaml(&modYaml) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				yaml.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

func (yaml *Yaml) GetYamls(c *gin.Context) {

	code := codes.SUCCESS
	name := c.Query("name")
	k8s_id := c.Query("k8s_id")
	application_id := c.Query("application_id")

	var maps string
	isAdmin := false

	userRole := jwt.ExtractClaims(c)
	//userName := userRole["userName"].(string)
	userID := int(userRole["userID"].(float64))
	userType := int(userRole["userType"].(float64))
	if userType == 1 {
		isAdmin = true
	}

	if name != "" {
		maps += " AND name LIKE '%" + name + "%'"
	}
	if k8s_id != "" {
		maps += " AND k8s_id = " + k8s_id
	}
	if application_id != "" {
		maps += " AND application_id = " + application_id
	}

	if isAdmin == false {
		maps += " AND user_id = " + strconv.Itoa(userID)
	}

	if len(maps) > 0 {
		maps = maps[4:len(maps)]
	}

	page, pagesize := GetPage(c)
	data := yaml.Service.GetYamls(page, pagesize, maps)
	RespData(c, http.StatusOK, code, data)
}

func (yaml *Yaml) DeleteYaml(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := codes.SUCCESS
	if !yaml.Service.DeleteYaml(uint(id)) {
		code = codes.ERROR
		RespFail(c, http.StatusOK, code, "删除出错，请联系管理员!")
	} else {
		RespOk(c, http.StatusOK, code)
	}
}
