package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/aguncn/nezha/common/datasource"
	"github.com/aguncn/nezha/common/setting"
	"github.com/aguncn/nezha/routers"
)

// 是否重新生成数据库
var migrate = flag.Bool("migrate", false, "if auto migrate all database?")

func main() {
	flag.Parse()
	datasource.Migrate = *migrate
	router := routers.InitRouter()
	conf := setting.Config.Server
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", conf.Port),
		Handler:        router,
		ReadTimeout:    conf.ReadTimeout * time.Second,
		WriteTimeout:   conf.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
