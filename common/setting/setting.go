package setting

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

//Config 定义配置
var (
	Config *Conf
)

//Conf 配置结构体
type Conf struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Harbor   Harbor   `yaml:"harbor"`
	APP      APP      `yaml:"app"`
}

//Server HTTP服务配置结构体
type Server struct {
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read-timeout"`
	WriteTimeout time.Duration `yaml:"write-timeout"`
}

//APP 程序配置
type APP struct {
	RunMode     string `yaml:"run-mode"`
	Pagesize    int    `yaml:"pagesize"`
	IdentityKey string `yaml:"identity-key"`
	LogPath     string `yaml:"log-path"`
}

//Database 数据库配置结构体
type Database struct {
	Type        string `yaml:"type"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"table-prefix"`
}

//Harbor配置结构体
type Harbor struct {
	ApiAddr  string `yaml:"api-addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

//init 初始化函数
func init() {
	Config = getConf()
	log.Println("[Setting] Config init success")
}

//getConf 读取配置文件
func getConf() *Conf {
	var c *Conf
	file, err := ioutil.ReadFile("../Config/config.yml")
	if err != nil {
		log.Println("[Setting] config error: ", err)
	}
	err = yaml.UnmarshalStrict(file, &c)
	if err != nil {
		log.Println("[Setting] yaml unmarshal error: ", err)
	}
	return c
}
