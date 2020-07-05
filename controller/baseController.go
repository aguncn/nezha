package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aguncn/nezha/common/codes"
	"github.com/aguncn/nezha/common/setting"
)

//ResponseData 数据返回结构体
type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//ResponseSuccess 返回成功结构体
type ResponseSuccess struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//ResponseFail 返回成功结构体
type ResponseFail struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Detail string `json:"detail"`
}

//RespData 数据返回
func RespData(c *gin.Context, httpCode, code int, data interface{}) {
	resp := ResponseData{
		Code: code,
		Msg:  codes.GetMsg(code),
		Data: data,
	}
	RespJSON(c, httpCode, resp)
}

//RespOk 返回操作成功
func RespOk(c *gin.Context, httpCode, code int) {
	resp := ResponseSuccess{
		Code: code,
		Msg:  codes.GetMsg(code),
	}
	RespJSON(c, httpCode, resp)
}

//RespFail 返回操作失败
func RespFail(c *gin.Context, httpCode, code int, detail string) {
	resp := ResponseFail{
		Code:   code,
		Msg:    codes.GetMsg(code),
		Detail: detail,
	}
	RespJSON(c, httpCode, resp)
}

//RespJSON 返回JSON数据
func RespJSON(c *gin.Context, httpCode int, resp interface{}) {
	c.JSON(httpCode, resp)
	c.Abort()
}

//GetPage 获取每页数量
func GetPage(c *gin.Context) (page, pagesize uint) {
	pageInt, _ := strconv.Atoi(c.Query("page"))
	pagesizeInt, _ := strconv.Atoi(c.Query("limit"))
	page = uint(pageInt)
	pagesize = uint(pagesizeInt)
	if pagesize == 0 {
		pagesize = uint(setting.Config.APP.Pagesize)
	}
	if page == 0 {
		page = 1
	}
	return
}
