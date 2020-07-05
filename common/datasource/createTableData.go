package datasource

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var mockData = make(map[string]string)

func CreateTableData(db *gorm.DB) error {
	initData()
	for _, record := range mockData {
		db.Exec(record)
	}
	return nil
}

func initData() {
	pw, _ := bcrypt.GenerateFromPassword([]byte("111111"), bcrypt.DefaultCost)
	mockData["adminUser"] = fmt.Sprintf("INSERT INTO nezha_user(id, username, password, user_type, state, created_at) VALUES "+
		" (1, 'admin', '%s', 1, 1, '2020-04-06 13:01:09');",
		pw)
	mockData["devUser"] = fmt.Sprintf("INSERT INTO nezha_user(id, username, password, user_type, state, created_at) VALUES "+
		" (2, 'dev', '%s', 2, 1, '2020-04-06 13:01:09');",
		pw)
	mockData["chengang"] = fmt.Sprintf("INSERT INTO nezha_user(id, username, password, user_type, state, created_at) VALUES "+
		" (3, 'chengang', '%s', 2, 1, '2020-04-06 13:01:09');",
		pw)

	mockData["k8sDev"] = fmt.Sprintf("INSERT INTO nezha_k8s(id, name, environment_id, user_id, conf, state, created_at) VALUES " +
		" (1, 'k8sDev', '1', '1',  'kubeconf k8s Dev', 1, '2020-04-06 13:01:09');")
	mockData["k8sPrd"] = fmt.Sprintf("INSERT INTO nezha_k8s(id, name, environment_id, user_id, conf, state, created_at) VALUES " +
		" (2, 'k8sPrd', '2', '2', 'kubeconf k8s Prd', 1, '2020-04-06 13:01:09');")
	mockData["k8sStg"] = fmt.Sprintf("INSERT INTO nezha_k8s(id, name, environment_id, user_id, conf, state, created_at) VALUES " +
		" (3, 'k8sStg', '1', '1', 'kubeconf k8s Prd', 1, '2020-04-06 13:01:09');")

	mockData["YamlAppK8s"] = fmt.Sprintf("INSERT INTO nezha_yaml(id, name, description, application_id, k8s_id, user_id, yaml, state, created_at) VALUES " +
		" (1, 'YamlAppK8s', 'desc', '1', '1', '1', 'app k8s yaml', 1, '2020-04-06 13:01:09');")

	mockData["project1"] = "INSERT INTO nezha_project(id, user_id, name, cn_name, description) VALUES (1, 1, 'AI-PROJECT', '算法项目', '用于算法')"
	mockData["project2"] = "INSERT INTO nezha_project(id, user_id, name, cn_name, description) VALUES (2, 2, 'BI-PROJECT', '商务智能项目', '用于BI智能')"
	mockData["project3"] = "INSERT INTO nezha_project(id, user_id, name, cn_name, description) VALUES (3, 1, 'CI-PROJECT', '持续集成项目', '用于JKINS')"

	mockData["env1"] = "INSERT INTO nezha_environment(id, name, description) VALUES (1, 'DEV', '用于开发测试的环境')"
	mockData["env2"] = "INSERT INTO nezha_environment(id, name, description) VALUES (2, 'PRD', '用于线上运行的环境')"

	for i := 1; i < 31; i++ {
		key := "application " + string(i)
		value := fmt.Sprintf("INSERT INTO nezha_application(id, name, cn_name,  description, harbor, project_id, user_id, state) VALUES (%d,  'application-%s' , 'app-cn-name-%s',  'description', '192.168.1.111:8089/tmp/spring-helloworld', 1, 1, 0);", i, strconv.Itoa(i), strconv.Itoa(i))
		mockData[key] = value
	}

	for i := 3; i < 3; i++ {
		key := "user- " + string(i)
		value := fmt.Sprintf("INSERT INTO nezha_user(id, username, password, user_type, state, created_at) VALUES (%d, 'user-%s', '%s', 2, 1, '2020-04-06 13:01:09');", i, strconv.Itoa(i), pw)
		mockData[key] = value
	}

}
