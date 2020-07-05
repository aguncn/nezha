package jwt

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/aguncn/nezha/common/codes"
	"github.com/aguncn/nezha/common/setting"
	"github.com/aguncn/nezha/models"
	"github.com/aguncn/nezha/service"
)

//### 如果是使用Go Module,gin-jwt模块应使用v2
//下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
//导入应写成 import "github.com/appleboy/gin-jwt/v2"
//### 如果不是使用Go Module
//下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
//导入import "github.com/appleboy/gin-jwt"

// JWT 注入IService
type JWT struct {
	UserService service.IUserService `inject:""`
}

//JwtAuthorizator 定义身份授权事件类型
type JwtAuthorizator func(data interface{}, c *gin.Context) bool

//app 程序配置
var app = setting.Config.APP

//GinJWTMiddlewareInit 初始化中间件
func (j *JWT) GinJWTMiddlewareInit(jwtAuthorizator JwtAuthorizator) (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Minute * 1440,
		MaxRefresh:  time.Hour,
		IdentityKey: app.IdentityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//handles the login logic. On success LoginResponse is called, on failure Unauthorized is called
			var loginVals models.User
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password
			if j.UserService.CheckUser(username, password) {
				userID := j.UserService.GetUserID(username)
				userType := j.UserService.GetUserType(username)
				return &models.UserRole{
					UserName: username,
					UserID:   userID,
					UserType: userType,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.UserRole); ok {
				//maps the claims in the JWT
				return jwt.MapClaims{
					"userID":   v.UserID,
					"userName": v.UserName,
					"userType": v.UserType,
				}
			}
			return jwt.MapClaims{}
		},
		//receives identity and handles authorization logic
		IdentityHandler: func(c *gin.Context) interface{} {
			userRole := jwt.ExtractClaims(c)
			userName := userRole["userName"].(string)
			userID := uint(userRole["userID"].(float64))
			userType := uint(userRole["userType"].(float64))
			return &models.UserRole{
				UserID:   userID,
				UserName: userName,
				UserType: userType,
			}
		},
		Authorizator: jwtAuthorizator,

		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return
}

//AdminAuthorizator role is admin can access
func AdminAuthorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.UserRole); ok {
		if v.UserType == 1 {
			return true
		}
	}

	return false
}

//DevAuthorizator username is test can access
func DevAuthorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.UserRole); ok {
		if v.UserType == 2 {
			return true
		}
	}
	return false
}

//AllUserAuthorizator all users access
func AllUserAuthorizator(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*models.UserRole); ok {
		return true
	}
	return false
}

//NoRouteHandler 404 handler
func NoRouteHandler(c *gin.Context) {
	code := codes.PageNotFound
	c.JSON(404, gin.H{"code": code, "message": codes.GetMsg(code)})
}
