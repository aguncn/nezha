package routers

import (
	"log"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"

	"github.com/aguncn/nezha/common/datasource"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/common/middleware/cors"
	"github.com/aguncn/nezha/common/middleware/jwt"
	"github.com/aguncn/nezha/common/setting"
	"github.com/aguncn/nezha/controller"
	"github.com/aguncn/nezha/repository"
	appRep "github.com/aguncn/nezha/repository/application"
	deployRep "github.com/aguncn/nezha/repository/deploy"
	envRep "github.com/aguncn/nezha/repository/environment"
	k8sRep "github.com/aguncn/nezha/repository/k8s"
	podRep "github.com/aguncn/nezha/repository/pod"
	proRep "github.com/aguncn/nezha/repository/project"
	yamlRep "github.com/aguncn/nezha/repository/yaml"
	"github.com/aguncn/nezha/service"
	appSvc "github.com/aguncn/nezha/service/application"
	deploySvc "github.com/aguncn/nezha/service/deploy"
	envSvc "github.com/aguncn/nezha/service/environment"
	k8sSvc "github.com/aguncn/nezha/service/k8s"
	podSvc "github.com/aguncn/nezha/service/pod"
	proSvc "github.com/aguncn/nezha/service/project"
	yamlSvc "github.com/aguncn/nezha/service/yaml"
)

//InitRouter 初始化Router
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())
	gin.SetMode(setting.Config.APP.RunMode)
	Configure(r)
	return r
}

//Configure 配置router
func Configure(r *gin.Engine) {
	//controller declare
	var user controller.User
	//var tag cv1.Tag
	var myjwt jwt.JWT
	//inject declare
	var application controller.Application
	var project controller.Project
	var env controller.Environment
	var k8s controller.K8s
	var pod controller.Pod
	var yaml controller.Yaml
	var deploy controller.Deploy
	db := datasource.Db{}
	zap := logger.Logger{}
	//Injection
	var injector inject.Graph
	if err := injector.Provide(
		&inject.Object{Value: &db},
		&inject.Object{Value: &zap},
		&inject.Object{Value: &myjwt},
		&inject.Object{Value: &application},
		&inject.Object{Value: &appRep.ApplicationRepository{}},
		&inject.Object{Value: &appSvc.ApplicationService{}},
		&inject.Object{Value: &project},
		&inject.Object{Value: &proRep.ProjectRepository{}},
		&inject.Object{Value: &proSvc.ProjectService{}},
		&inject.Object{Value: &env},
		&inject.Object{Value: &envRep.EnvironmentRepository{}},
		&inject.Object{Value: &envSvc.EnvironmentService{}},
		&inject.Object{Value: &k8s},
		&inject.Object{Value: &k8sRep.K8sRepository{}},
		&inject.Object{Value: &k8sSvc.K8sService{}},
		&inject.Object{Value: &pod},
		&inject.Object{Value: &podRep.PodRepository{}},
		&inject.Object{Value: &podSvc.PodService{}},
		&inject.Object{Value: &yaml},
		&inject.Object{Value: &yamlRep.YamlRepository{}},
		&inject.Object{Value: &yamlSvc.YamlService{}},
		&inject.Object{Value: &deploy},
		&inject.Object{Value: &deployRep.DeployRepository{}},
		&inject.Object{Value: &deploySvc.DeployService{}},
		&inject.Object{Value: &user},
		&inject.Object{Value: &repository.UserRepository{}},
		&inject.Object{Value: &service.UserService{}},
		&inject.Object{Value: &repository.BaseRepository{}},
	); err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("injector fatal: ", err)
	}
	//zap log init
	zap.Init()
	//database connect
	if err := db.Connect(); err != nil {
		log.Fatal("db fatal:", err)
	}
	r.POST("/register", user.RegisterUser)
	var authMiddleware = myjwt.GinJWTMiddlewareInit(jwt.AllUserAuthorizator)
	r.NoRoute(authMiddleware.MiddlewareFunc(), jwt.NoRouteHandler)
	r.POST("/login", authMiddleware.LoginHandler)

	allUserAPI := r.Group("/")
	allUserAPI.Use(authMiddleware.MiddlewareFunc())
	{
		allUserAPI.GET("/project/list", project.GetProjects)
		allUserAPI.GET("/environment/list", env.GetEnvironments)
		allUserAPI.GET("/k8s/list", k8s.GetK8ss)
		allUserAPI.GET("/k8s/detail/:id", k8s.GetK8s)
		allUserAPI.POST("/k8s/create", k8s.AddK8s)
		allUserAPI.PUT("/k8s/update", k8s.UpdateK8s)
		allUserAPI.DELETE("/k8s/delete/:id", k8s.DeleteK8s)
		allUserAPI.GET("/pod/list", pod.GetPods)
		allUserAPI.GET("/pod/detail/:id", pod.GetPod)
		allUserAPI.POST("/pod/create", pod.AddPod)
		allUserAPI.PUT("/pod/update", pod.UpdatePod)
		allUserAPI.DELETE("/pod/delete/:id", pod.DeletePod)
		allUserAPI.GET("/yaml/list", yaml.GetYamls)
		allUserAPI.GET("/yaml/detail/:id", yaml.GetYaml)
		allUserAPI.POST("/yaml/create", yaml.AddYaml)
		allUserAPI.PUT("/yaml/update", yaml.UpdateYaml)
		allUserAPI.DELETE("/yaml/delete/:id", yaml.DeleteYaml)

	}

	applicationAPI := r.Group("/application")
	applicationAPI.Use(authMiddleware.MiddlewareFunc())
	{
		applicationAPI.GET("/list", application.GetApplications)
		applicationAPI.GET("/recentList", application.RecentApplications)
		applicationAPI.GET("/detail/:id", application.GetApplication)
		applicationAPI.POST("/create", application.AddApplication)
		applicationAPI.PUT("/update", application.UpdateApplication)
		applicationAPI.DELETE("/delete/:id", application.DeleteApplication)

	}
	deployAPI := r.Group("/deploy")
	deployAPI.Use(authMiddleware.MiddlewareFunc())
	{
		deployAPI.GET("/list", deploy.GetDeploys)
		deployAPI.GET("/detail/:id", deploy.GetDeploy)
		deployAPI.POST("/create", deploy.AddDeploy)
		deployAPI.PUT("/update", deploy.UpdateDeploy)
		deployAPI.DELETE("/delete/:id", deploy.DeleteDeploy)
		deployAPI.POST("/submit", deploy.SubmitDeploy)
		deployAPI.POST("/status", deploy.StatusDeploy)
		deployAPI.GET("/dockerTag", deploy.GetDockerTag)
		deployAPI.GET("/okErrorCount", deploy.GetOkErrorCount)
		deployAPI.GET("/cycleCount", deploy.GetBetweenCount)

	}

	userAPI := r.Group("/user")
	{
		userAPI.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	userAPI.Use(authMiddleware.MiddlewareFunc())
	{
		userAPI.GET("/info", user.GetUserInfo)
		userAPI.PUT("/update", user.UpdateUser)
		userAPI.POST("/logout", user.Logout)
	}

	var adminMiddleware = myjwt.GinJWTMiddlewareInit(jwt.AdminAuthorizator)
	apiv1 := r.Group("/api/v1")
	//使用AdminAuthorizator中间件，只有admin权限的用户才能获取到接口
	apiv1.Use(adminMiddleware.MiddlewareFunc())
	{
		//vue获取table信息
		apiv1.GET("/user/list", user.GetUsers)
		apiv1.POST("/user", user.AddUser)
		apiv1.DELETE("/user/:id", user.DeleteUser)
		apiv1.GET("/project/detail/:id", project.GetProject)
		apiv1.POST("/project/create", project.AddProject)
		apiv1.PUT("/project/update", project.UpdateProject)
		apiv1.DELETE("/project/delete/:id", project.DeleteProject)

		apiv1.GET("/environment/detail/:id", env.GetEnvironment)
		apiv1.POST("/environment/create", env.AddEnvironment)
		apiv1.PUT("/environment/update", env.UpdateEnvironment)
		apiv1.DELETE("/environment/delete/:id", env.DeleteEnvironment)
	}
}
