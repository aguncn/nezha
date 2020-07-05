package common

import (
	//"io/ioutil"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// 初始化k8s客户端
func InitClient(k8sConf *string) (kubernetes.Clientset, error) {

	// 读kubeconfig文件
	//kubeconfig, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	return rest.Config{}, err
	//}
	// 生成rest client配置
	restConf, err := clientcmd.RESTConfigFromKubeConfig([]byte(*k8sConf))
	if err != nil {
		return kubernetes.Clientset{}, err
	}

	// 生成clientset配置
	clientset, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		return kubernetes.Clientset{}, err
	}

	return *clientset, nil
}

// 获取k8s restful client配置
func GetRestConf(k8sConf string) (rest.Config, error) {

	// 读kubeconfig文件
	//kubeconfig, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	return rest.Config{}, err
	//}
	// 生成rest client配置
	restConf, err := clientcmd.RESTConfigFromKubeConfig([]byte(k8sConf))
	if err != nil {
		return rest.Config{}, err
	}

	return *restConf, nil
}
