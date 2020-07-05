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
	k8sSvc "github.com/aguncn/nezha/service/k8s"
)

type K8s struct {
	Log     logger.ILogger     `inject:""`
	Service k8sSvc.IK8sService `inject:""`
}

func (k8s *K8s) GetK8s(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	var data *models.K8s
	code := codes.InvalidParams
	if !valid.HasErrors() {
		data = k8s.Service.GetK8s(uint(id))
		code = codes.SUCCESS
	} else {
		for _, err := range valid.Errors {
			k8s.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	RespData(c, http.StatusOK, code, data)
}

func (k8s *K8s) AddK8s(c *gin.Context) {
	modK8s := models.K8s{}
	code := codes.InvalidParams
	err := c.Bind(&modK8s)
	fmt.Printf("modK8s: %v", modK8s)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(modK8s.Name, "name").Message("名称不能为空")
		valid.Required(modK8s.CreatedBy, "created_by").Message("创建人不能为空")
		if !valid.HasErrors() {
			if k8s.Service.AddK8s(&modK8s) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				k8s.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

func (k8s *K8s) UpdateK8s(c *gin.Context) {
	modK8s := models.K8s{}
	code := codes.InvalidParams
	err := c.Bind(&modK8s)
	fmt.Printf("modK8s: %v", modK8s)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(modK8s.Name, "name").Message("Name不能为空")
		if !valid.HasErrors() {
			if k8s.Service.UpdateK8s(&modK8s) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				k8s.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

func (k8s *K8s) GetK8ss(c *gin.Context) {

	code := codes.SUCCESS
	name := c.Query("name")
	environment_id := c.Query("environment_id")

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
	if environment_id != "" {
		maps += " AND environment_id = " + environment_id
	}

	if isAdmin == false {
		maps += " AND user_id = " + strconv.Itoa(userID)
	}

	if len(maps) > 0 {
		maps = maps[4:len(maps)]
	}

	page, pagesize := GetPage(c)
	data := k8s.Service.GetK8ss(page, pagesize, maps)
	RespData(c, http.StatusOK, code, data)
}

func (k8s *K8s) DeleteK8s(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := codes.SUCCESS
	if !k8s.Service.DeleteK8s(uint(id)) {
		code = codes.ERROR
		RespFail(c, http.StatusOK, code, "删除出错，请联系管理员!")
	} else {
		RespOk(c, http.StatusOK, code)
	}
}
