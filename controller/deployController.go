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
	deploySvc "github.com/aguncn/nezha/service/deploy"
)

type Deploy struct {
	Log     logger.ILogger           `inject:""`
	Service deploySvc.IDeployService `inject:""`
}

func (deploy *Deploy) GetDeploy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	var data *models.Deploy
	code := codes.InvalidParams
	if !valid.HasErrors() {
		data = deploy.Service.GetDeploy(uint(id))
		code = codes.SUCCESS
	} else {
		for _, err := range valid.Errors {
			deploy.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	RespData(c, http.StatusOK, code, data)
}

func (deploy *Deploy) AddDeploy(c *gin.Context) {
	modDeploy := models.Deploy{}
	code := codes.InvalidParams
	err := c.Bind(&modDeploy)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(modDeploy.Name, "name").Message("名称不能为空")
		valid.Required(modDeploy.CreatedBy, "created_by").Message("创建人不能为空")
		if !valid.HasErrors() {
			if deploy.Service.AddDeploy(&modDeploy) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				deploy.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

func (deploy *Deploy) UpdateDeploy(c *gin.Context) {
	modDeploy := models.Deploy{}
	code := codes.InvalidParams
	err := c.Bind(&modDeploy)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(modDeploy.Name, "name").Message("Name不能为空")
		if !valid.HasErrors() {
			if deploy.Service.UpdateDeploy(&modDeploy) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				deploy.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

func (deploy *Deploy) GetDeploys(c *gin.Context) {
	code := codes.SUCCESS
	name := c.Query("name")
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
	data := deploy.Service.GetDeploys(page, pagesize, maps)
	RespData(c, http.StatusOK, code, data)
}

func (deploy *Deploy) DeleteDeploy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := codes.SUCCESS
	if !deploy.Service.DeleteDeploy(uint(id)) {
		code = codes.ERROR
		RespFail(c, http.StatusOK, code, "删除出错，请联系管理员!")
	} else {
		RespOk(c, http.StatusOK, code)
	}
}

func (deploy *Deploy) SubmitDeploy(c *gin.Context) {
	var msg string
	modDeploy := models.Deploy{}
	code := codes.InvalidParams
	err := c.Bind(&modDeploy)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(modDeploy.Name, "name").Message("批次不能为空")
		valid.Required(modDeploy.Yaml, "yaml").Message("Yaml不能为空")
		if !valid.HasErrors() {
			ok, stdout, stderr := deploy.Service.SubmitDeploy(&modDeploy)

			if ok {
				msg = stdout
				code = codes.SUCCESS
			} else {
				msg = stderr
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				fmt.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespData(c, http.StatusOK, code, msg)
}

func (deploy *Deploy) StatusDeploy(c *gin.Context) {
	var msg string
	modDeploy := models.Deploy{}
	code := codes.InvalidParams
	err := c.Bind(&modDeploy)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(modDeploy.Name, "name").Message("批次不能为空")
		valid.Required(modDeploy.Yaml, "yaml").Message("Yaml不能为空")
		if !valid.HasErrors() {
			fmt.Println(modDeploy.UserID, "@@@@@@@@@@@@@")
			ok, stdout, stderr := deploy.Service.StatusDeploy(&modDeploy)

			if ok {
				msg = stdout
				code = codes.SUCCESS
			} else {
				msg = stderr
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				fmt.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespData(c, http.StatusOK, code, msg)
}

func (deploy *Deploy) GetDockerTag(c *gin.Context) {

	harbor := c.Query("harbor")
	code := codes.InvalidParams
	ok, data := deploy.Service.GetDockerTag(harbor)

	if ok {
		code = codes.SUCCESS
	} else {
		code = codes.ERROR
	}

	RespData(c, http.StatusOK, code, data)
}

//返回dashborad的饼状图
func (deploy *Deploy) GetOkErrorCount(c *gin.Context) {

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
		maps += "user_id = " + string(userID)
	}
	data := deploy.Service.GetOkErrorCount(maps)
	code := codes.SUCCESS
	RespData(c, http.StatusOK, code, data)
}

//返回dashboard的柱状图
func (deploy *Deploy) GetBetweenCount(c *gin.Context) {
	cycle, _ := strconv.Atoi(c.Query("cycle"))
	count, _ := strconv.Atoi(c.Query("count"))

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
		maps += "user_id = " + string(userID)
	}
	data := deploy.Service.GetBetweenCount(maps, cycle, count)
	code := codes.SUCCESS
	RespData(c, http.StatusOK, code, data)
}
