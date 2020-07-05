package datasource

import (
	"fmt"
	"log"

	"github.com/aguncn/nezha/common/setting"
	"github.com/aguncn/nezha/models"
	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Migrate bool
)

//Db gormDB
type Db struct {
	Conn *gorm.DB
}

//Connect 初始化数据库配置
func (d *Db) Connect() error {
	var (
		dbType, dbName, user, pwd, host string
	)

	conf := setting.Config.Database
	dbType = conf.Type
	dbName = conf.Name
	user = conf.User
	pwd = conf.Password
	host = conf.Host

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, dbName))
	if err != nil {
		log.Fatal("connecting mysql error: ", err)
		return err
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}
	db.LogMode(true) //打印SQL语句
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	d.Conn = db

	log.Println("Connect Mysql Success.")

	if Migrate {
		d.DropAndCreateTb()
		log.Println("Drop And Create Table...")
		if err := CreateTableData(db); err != nil {
			log.Fatal("CreateTableData error: ", err)
			return err
		}
		log.Println("Create Table...")
	}

	return nil
}

//DB 返回DB
func (d *Db) DB() *gorm.DB {
	return d.Conn
}

func (d *Db) DropAndCreateTb() {
	d.DropAndCreateSingleTb(&models.Application{})
	d.DropAndCreateSingleTb(&models.Project{})
	d.DropAndCreateSingleTb(&models.User{})
	d.DropAndCreateSingleTb(&models.K8s{})
	d.DropAndCreateSingleTb(&models.Yaml{})
	d.DropAndCreateSingleTb(&models.Deploy{})
	d.DropAndCreateSingleTb(&models.Release{})
	d.DropAndCreateSingleTb(&models.Pod{})
	d.DropAndCreateSingleTb(&models.Environment{})
}

func (d *Db) DropAndCreateSingleTb(Tb interface{}) {
	d.Conn.DropTableIfExists(Tb)
	d.Conn.CreateTable(Tb)

}
