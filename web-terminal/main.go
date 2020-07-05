package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	clientset *kubernetes.Clientset
	kconfig   *rest.Config
	//apiUrl    *string
	err error
	//userToken string
)

func main() {
	staticDir := "./public"
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fs)

	cfgpath := flag.String("kubecfg", filepath.Join(homeDir(), ".kube", "config"), "kubeconfig绝对路径")
	httpPort := flag.String("port", "8000", "监听端口号")
	//apiUrl = flag.String("api", "http://127.0.0.1:8000/api.html", "权限api地址")

	flag.Parse()

	log.Println("加载kubeconfig的路径:", *cfgpath)
	log.Println("端口号:", *httpPort)
	//log.Println("权限认证api:", *apiUrl)

	if kconfig, err = clientcmd.BuildConfigFromFlags("", *cfgpath); err != nil {
		log.Fatalf("error creating k8s config: %v", err)
	}

	if clientset, err = kubernetes.NewForConfig(kconfig); err != nil {
		log.Fatalf("error creating clientset: %v", err)
	}

	http.HandleFunc("/terminal", wsHandler)

	log.Println("server running...")
	http.ListenAndServe(":"+*httpPort, nil)
	log.Println("server stopped...")

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}

	return os.Getenv("USERPROFILE") // windows
}
