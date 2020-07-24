package deploy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/aguncn/nezha/common/setting"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apps_v1 "k8s.io/api/apps/v1"
	yaml_k8s "k8s.io/apimachinery/pkg/util/yaml"

	"github.com/aguncn/nezha/common"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	"github.com/aguncn/nezha/service/k8s"
	"github.com/aguncn/nezha/service/pod"

	//pageModel "github.com/aguncn/nezha/page"
	deployRep "github.com/aguncn/nezha/repository/deploy"
	"github.com/jinzhu/gorm"
)

type DeployService struct {
	Log        logger.ILogger              `inject:""`
	Repository deployRep.IDeployRepository `inject:""`
	K8sService k8s.IK8sService             `inject:""`
	PodService pod.IPodService             `inject:""`
}

func (deploy *DeployService) GetDeploy(id uint) *models.Deploy {
	where := models.Deploy{Model: gorm.Model{ID: id}}
	return deploy.Repository.GetDeploy(&where)
}

func (deploy *DeployService) GetDeployByName(name string) *models.Deploy {
	where := models.Deploy{Name: name}
	return deploy.Repository.GetDeploy(&where)
}

func (deploy *DeployService) AddDeploy(modDeploy *models.Deploy) bool {
	if deploy.Repository.ExistDeployByName(modDeploy) {
		return false
	} else {
		return deploy.Repository.AddDeploy(modDeploy)
	}
}

func (deploy *DeployService) UpdateDeploy(modDeploy *models.Deploy) bool {

	return deploy.Repository.UpdateDeploy(modDeploy)

}

func (deploy *DeployService) GetDeploys(page, pagesize uint, maps interface{}) interface{} {
	res := make(map[string]interface{}, 2)
	var total uint64
	Deploys := deploy.Repository.GetDeploys(page, pagesize, &total, maps)
	res["list"] = &Deploys
	res["total"] = total
	return &res
}

func (deploy *DeployService) GetOkErrorCount(maps string) interface{} {
	Deploys := deploy.Repository.GetOkErrorCount(maps)
	return &Deploys
}

func (deploy *DeployService) GetBetweenCount(maps string, cycle, count int) interface{} {
	Deploys := deploy.Repository.GetBetweenCount(maps, cycle, count)
	return &Deploys
}

func (deploy *DeployService) DeleteDeploy(id uint) bool {
	return deploy.Repository.DeleteDeploy(id)
}

func (deploy *DeployService) SubmitDeploy(modDeploy *models.Deploy) (bool, string, string) {
	fileName := modDeploy.Name
	modK8s := deploy.K8sService.GetK8s(modDeploy.K8sID)
	var cmdStr, yamlPath, k8sPath string
	switch os := runtime.GOOS; os {
	case "windows":
		yamlPath = fmt.Sprintf("D:\\tmp\\tmp\\%s", fileName)
		k8sPath = fmt.Sprintf("D:\\tmp\\tmp\\%s", fileName+"-k8s")
		cmdStr = fmt.Sprintf("D:\\docker\\kubectl.exe --kubeconfig %s apply -f %s", k8sPath, yamlPath)
	case "linux":
		yamlPath = fmt.Sprintf("/tmp/%s", fileName)
		k8sPath = fmt.Sprintf("/tmp/%s", fileName+"-k8s")
		cmdStr = fmt.Sprintf("/usr/local/bin/kubectl --kubeconfig %s apply -f %s", k8sPath, yamlPath)
	}
	//生成yaml文件
	yamlFile, err := os.Create(yamlPath)
	if err != nil {
		fmt.Println(err.Error())
		return false, "", err.Error()
	}
	yamlFile.WriteString(modDeploy.Yaml)
	defer yamlFile.Close()
	//生成kubeconfig文件
	k8sFile, err := os.Create(k8sPath)
	if err != nil {
		fmt.Println(err.Error())
		return false, "", err.Error()
	}
	k8sFile.WriteString(modK8s.Conf)
	defer k8sFile.Close()
	//调用Kubectl
	okResult, stdoutStr, stderrStr := exeCommand(cmdStr, 200000)
	if okResult == false {
		updateDeploy := deploy.GetDeployByName(modDeploy.Name)
		where := map[string]interface{}{"step": 2, "log": stderrStr, "result": "error"}
		if deploy.Repository.UpdateDeployByCol(updateDeploy, where) {
			fmt.Println("update log finish.")
		}

	}
	return okResult, stdoutStr, stderrStr

}

func (deploy *DeployService) StatusDeploy(modDeploy *models.Deploy) (bool, string, string) {
	yaml := modDeploy.Yaml
	modK8s := deploy.K8sService.GetK8s(modDeploy.K8sID)
	k8sConf := modK8s.Conf
	var ns string

	dataArr := strings.Split(yaml, "---")
	for _, value := range dataArr {
		// YAML转JSON
		deployJson, err := yaml_k8s.ToJSON([]byte(value))
		if err != nil {
			fmt.Println(err)
		}
		// JSON转struct
		var deployment apps_v1.Deployment
		err = json.Unmarshal(deployJson, &deployment)
		if err != nil {
			fmt.Println(err)
		}
		if deployment.Kind == "Deployment" {
			//dp = deployment.Name
			ns = deployment.Namespace
			//lb = deployment.Spec.Template.Labels
			if len(ns) == 0 {
				ns = "default"
			}
			clientset, err := common.InitClient(&k8sConf)
			if err != nil {
				fmt.Println(err)
			}
			success, reasons, errStr := GetDeploymentStatus(clientset, deployment, ns)

			updateDeploy := deploy.GetDeployByName(modDeploy.Name)
			// 如果部署成功，则更新在线POD的数据，方便使用web terminal进入。
			if success == true {
				where := map[string]interface{}{"step": 3, "log": reasons, "result": "success"}
				if deploy.Repository.UpdateDeployByCol(updateDeploy, where) {
					fmt.Println("update log finish.")
				}

				// 拼接selector
				labelSelector := ""
				for key, value := range deployment.Spec.Selector.MatchLabels {
					labelSelector = labelSelector + key + "=" + value + ","
				}
				labelSelector = strings.TrimRight(labelSelector, ",")

				// 获取pod状态
				podList, err := clientset.CoreV1().Pods(ns).List(meta_v1.ListOptions{LabelSelector: labelSelector})
				if err != nil {
					fmt.Println(err)
				}
				// 删除旧的，插入新的POD
				appID := modDeploy.ApplicationID
				envID := modDeploy.EnvironmentID
				k8sID := modDeploy.K8sID
				userID := modDeploy.UserID
				delPod := models.Pod{
					ApplicationID: appID,
					EnvironmentID: envID,
					K8sID:         k8sID,
					UserID:        userID,
				}
				if deploy.PodService.DeletePodByCol(&delPod) {
					fmt.Println("delete pod success.")
				} else {
					fmt.Println("delete pod error.")
				}
				for _, pod := range podList.Items {
					for index, container := range pod.Spec.Containers {
						podName := fmt.Sprintf("%s-%s", pod.Name, strconv.Itoa(index))
						modPod := models.Pod{
							PodName:       podName,
							NameSpace:     ns,
							Container:     container.Name,
							ApplicationID: appID,
							EnvironmentID: envID,
							K8sID:         k8sID,
							UserID:        userID,
						}
						if deploy.PodService.AddPod(&modPod) {
							fmt.Println("update pod success.")
						} else {
							fmt.Println("update pod error.")
						}

					}
				}
			} else {
				where := map[string]interface{}{"step": 3, "log": reasons, "result": "error"}
				if deploy.Repository.UpdateDeployByCol(updateDeploy, where) {
					fmt.Println("update log finish.")
				}

			}
			return success, reasons, errStr
		}

	}
	return false, "", "error"

}

func exeCommand(cmdStr string, killTime time.Duration) (okResult bool, stdout, stderr string) {
	fmt.Println("running bash command...")
	fmt.Println(cmdStr)
	var cmdExe, cmdOpt string
	switch os := runtime.GOOS; os {
	case "windows":
		cmdExe = "cmd"
		cmdOpt = "/C"
	case "linux":
		cmdExe = "sh"
		cmdOpt = "-c"
	}
	cmd := exec.Command(cmdExe, cmdOpt, cmdStr)

	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	okResult = true

	err := cmd.Start()
	log.Printf(time.Now().Format("2006/1/2 15:04:05"))
	log.Printf("Waiting for command to finish...")
	done := make(chan error, 1)
	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(killTime * time.Millisecond):
		if err := cmd.Process.Kill(); err != nil {
			log.Printf(time.Now().Format("2006/1/2 15:04:05"))
			log.Fatal("failed to kill: ", err)
			okResult = false
		}
		<-done // allow goroutine to exit
		// log.Println("process killed")
	case err := <-done:
		if err != nil {
			log.Printf(time.Now().Format("2006/1/2 15:04:05"))
			log.Printf("process done with error = %v", err)
			okResult = false
		}
		stdout = outbuf.String()
		stderr = errbuf.String()
	}
	if err != nil {
		log.Fatal(err)
		okResult = false
	}
	return
}

func (deploy *DeployService) GetDockerTag(harbor string) (bool, []models.DockerTag) {
	// 获取harbor地址，然后，认证成功之后，拉取指定docker的tag列表
	harborConf := setting.Config.Harbor
	harborApiAddr := harborConf.ApiAddr
	harborUser := harborConf.User
	harborPwd := harborConf.Password

	harborUrl := harborApiAddr + "/" + harbor + "/tags"

	client := &http.Client{}
	request, err := http.NewRequest("GET", harborUrl, nil)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	request.SetBasicAuth(harborUser, harborPwd)
	request.Header.Add("Accept", "application/json")
	response, err := client.Do(request)
	defer response.Body.Close()
	body := []models.DockerTag{}
	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
			os.Exit(0)
		}
		err = json.Unmarshal(bodyByte, &body)
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
			os.Exit(0)
		}
		return true, body
	} else {
		return false, body
	}

}
