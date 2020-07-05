package controller

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/aguncn/nezha/common/codes"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	proSvc "github.com/aguncn/nezha/service/project"
)

// Project 注入IProjectService
type Project struct {
	Log     logger.ILogger         `inject:""`
	Service proSvc.IProjectService `inject:""`
}

//GetProject 获取单个Project
func (p *Project) GetProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	var data *models.Project
	code := codes.InvalidParams
	if !valid.HasErrors() {
		data = p.Service.GetProject(uint(id))
		code = codes.SUCCESS
	} else {
		for _, err := range valid.Errors {
			p.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	RespData(c, http.StatusOK, code, data)
}

//AddProject 新增Project
func (p *Project) AddProject(c *gin.Context) {

	Project := models.Project{}

	code := codes.InvalidParams
	err := c.Bind(&Project)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(Project.Name, "name").Message("名称不能为空")
		valid.Required(Project.Description, "description").Message("简述不能为空")
		valid.Required(Project.CnName, "cn_name").Message("中文名不能为空")
		valid.Required(Project.CreatedBy, "created_by").Message("创建人不能为空")
		if !valid.HasErrors() {
			if p.Service.AddProject(&Project) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				p.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

//UpdateProject 修改Project
func (p *Project) UpdateProject(c *gin.Context) {
	project := models.Project{}
	code := codes.InvalidParams
	err := c.Bind(&project)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(project.Name, "name").Message("Name不能为空")
		valid.Required(project.CnName, "cn_name").Message("CnName不能为空")
		if !valid.HasErrors() {
			if p.Service.UpdateProject(&project) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				p.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

//GetProjects 获取项目信息
func (p *Project) GetProjects(c *gin.Context) {
	var maps string
	code := codes.SUCCESS
	name := c.Query("name")
	if name != "" {
		maps = "name LIKE '%" + name + "%'"
	}
	page, pagesize := GetPage(c)
	data := p.Service.GetProjects(page, pagesize, maps)
	RespData(c, http.StatusOK, code, data)
}

//DeleteProject 删除项目
func (p *Project) DeleteProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := codes.SUCCESS
	if !p.Service.DeleteProject(uint(id)) {
		code = codes.ERROR
		RespFail(c, http.StatusOK, code, "删除出错，请联系管理员!")
	} else {
		RespOk(c, http.StatusOK, code)
	}
}
