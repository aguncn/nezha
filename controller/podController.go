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
	podSvc "github.com/aguncn/nezha/service/pod"
)

type Pod struct {
	Log     logger.ILogger     `inject:""`
	Service podSvc.IPodService `inject:""`
}

func (pod *Pod) GetPod(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	var data *models.Pod
	code := codes.InvalidParams
	if !valid.HasErrors() {
		data = pod.Service.GetPod(uint(id))
		code = codes.SUCCESS
	} else {
		for _, err := range valid.Errors {
			pod.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	RespData(c, http.StatusOK, code, data)
}

func (pod *Pod) AddPod(c *gin.Context) {
	modPod := models.Pod{}
	code := codes.InvalidParams
	err := c.Bind(&modPod)
	fmt.Printf("modPod: %v", modPod)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(modPod.PodName, "podname").Message("名称不能为空")
		valid.Required(modPod.CreatedBy, "created_by").Message("创建人不能为空")
		if !valid.HasErrors() {
			if pod.Service.AddPod(&modPod) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				pod.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

func (pod *Pod) UpdatePod(c *gin.Context) {
	modPod := models.Pod{}
	code := codes.InvalidParams
	err := c.Bind(&modPod)
	fmt.Printf("modPod: %v", modPod)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(modPod.PodName, "podname").Message("Name不能为空")
		if !valid.HasErrors() {
			if pod.Service.UpdatePod(&modPod) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				pod.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

func (pod *Pod) GetPods(c *gin.Context) {
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
	data := pod.Service.GetPods(page, pagesize, maps)
	RespData(c, http.StatusOK, code, data)
}

func (pod *Pod) DeletePod(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := codes.SUCCESS
	if !pod.Service.DeletePod(uint(id)) {
		code = codes.ERROR
		RespFail(c, http.StatusOK, code, "删除出错，请联系管理员!")
	} else {
		RespOk(c, http.StatusOK, code)
	}
}
