package deploy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	apps_v1 "k8s.io/api/apps/v1"
	core_v1 "k8s.io/api/core/v1"

	//"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

// 读取deployment yaml内容
func ReadDeploymentYaml(filename string) apps_v1.Deployment {
	// 读取YAML
	deployYaml, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return apps_v1.Deployment{}
	}

	// YAML转JSON
	deployJson, err := yaml.ToJSON(deployYaml)
	if err != nil {
		fmt.Println(err)
		return apps_v1.Deployment{}
	}

	// JSON转struct
	var deployment apps_v1.Deployment
	err = json.Unmarshal(deployJson, &deployment)
	if err != nil {
		fmt.Println(err)
		return apps_v1.Deployment{}
	}
	return deployment
}

func GetDeploymentStatus(clientset kubernetes.Clientset, deployment apps_v1.Deployment, namespace string) (bool, string, string) {

	// 获取deployment状态
	k8sDeployment, err := clientset.AppsV1().Deployments(namespace).Get(deployment.Name, meta_v1.GetOptions{})
	if err != nil {
		return false, "", "get deployments status error"
	}

	// 获取pod的状态
	labelSelector := ""
	for key, value := range deployment.Spec.Selector.MatchLabels {
		labelSelector = labelSelector + key + "=" + value + ","
	}
	labelSelector = strings.TrimRight(labelSelector, ",")

	podList, err := clientset.CoreV1().Pods(namespace).List(meta_v1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return false, "", "get pods status error"
	}

	var readyPod, unavailablePod int
	var waitingReasons string
	for _, pod := range podList.Items {
		// 记录等待原因
		for _, containerStatus := range pod.Status.ContainerStatuses {
			if containerStatus.State.Waiting != nil {
				reason := "pod " + pod.Name + ", container " + containerStatus.Name + ", waiting reason: " + containerStatus.State.Waiting.Reason
				waitingReasons += reason
			}
		}

		podScheduledCondition := GetPodCondition(pod.Status, core_v1.PodScheduled)
		initializedCondition := GetPodCondition(pod.Status, core_v1.PodInitialized)
		readyCondition := GetPodCondition(pod.Status, core_v1.PodReady)
		containersReadyCondition := GetPodCondition(pod.Status, core_v1.ContainersReady)

		if pod.Status.Phase == "Running" &&
			podScheduledCondition.Status == "True" &&
			initializedCondition.Status == "True" &&
			readyCondition.Status == "True" &&
			containersReadyCondition.Status == "True" {
			readyPod++
		} else {
			unavailablePod++
		}
	}

	// 根据container状态判定
	if len(waitingReasons) != 0 {
		return false, "", waitingReasons
	}

	// 根据pod状态判定
	if int32(readyPod) < *(k8sDeployment.Spec.Replicas) ||
		int32(unavailablePod) != 0 {
		return false, "", "pods not ready!"
	}

	// deployment进行状态判定
	availableCondition := GetDeploymentCondition(k8sDeployment.Status, apps_v1.DeploymentAvailable)
	progressingCondition := GetDeploymentCondition(k8sDeployment.Status, apps_v1.DeploymentProgressing)

	if k8sDeployment.Status.UpdatedReplicas != *(k8sDeployment.Spec.Replicas) ||
		k8sDeployment.Status.Replicas != *(k8sDeployment.Spec.Replicas) ||
		k8sDeployment.Status.AvailableReplicas != *(k8sDeployment.Spec.Replicas) ||
		availableCondition.Status != "True" ||
		progressingCondition.Status != "True" {
		return false, "", "deployments not ready!"
	}

	if k8sDeployment.Status.ObservedGeneration < k8sDeployment.Generation {
		return false, "", "observed generation less than generation!"
	}

	// 发布成功
	return true, "deployments is success", "success"
}

func PrintDeploymentStatus(clientset kubernetes.Clientset, deployment apps_v1.Deployment, namespace string) {

	// 拼接selector
	labelSelector := ""
	for key, value := range deployment.Spec.Selector.MatchLabels {
		labelSelector = labelSelector + key + "=" + value + ","
	}
	labelSelector = strings.TrimRight(labelSelector, ",")

	for {
		// 获取k8s中deployment的状态
		k8sDeployment, err := clientset.AppsV1().Deployments(namespace).Get(deployment.Name, meta_v1.GetOptions{})
		if err != nil {
			fmt.Println(err)
		}

		// 打印deployment状态
		fmt.Printf("-------------deployment status------------\n")
		fmt.Printf("deployment.name: %s\n", k8sDeployment.Name)
		fmt.Printf("deployment.generation: %d\n", k8sDeployment.Generation)
		fmt.Printf("deployment.status.observedGeneration: %d\n", k8sDeployment.Status.ObservedGeneration)
		fmt.Printf("deployment.spec.replicas: %d\n", (k8sDeployment.Spec.Replicas))
		fmt.Printf("deployment.status.replicas: %d\n", k8sDeployment.Status.Replicas)
		fmt.Printf("deployment.status.updatedReplicas: %d\n", k8sDeployment.Status.UpdatedReplicas)
		fmt.Printf("deployment.status.readyReplicas: %d\n", k8sDeployment.Status.ReadyReplicas)
		fmt.Printf("deployment.status.unavailableReplicas: %d\n", k8sDeployment.Status.UnavailableReplicas)
		for _, condition := range k8sDeployment.Status.Conditions {
			fmt.Printf("condition.type: %s, condition.status: %s, condition.reason: %s\n", condition.Type, condition.Status, condition.Reason)
		}

		// 获取pod状态
		podList, err := clientset.CoreV1().Pods(namespace).List(meta_v1.ListOptions{LabelSelector: labelSelector})
		if err != nil {
			fmt.Println(err)
			return
		}

		for index, pod := range podList.Items {
			// 打印pod的状态
			fmt.Printf("-------------pod %d status------------\n", index)
			fmt.Printf("pod.name: %s\n", pod.Name)
			for _, container := range pod.Spec.Containers {
				fmt.Printf("pod.container.name: %s\n", container.Name)
				fmt.Printf("pod.container.image: %s\n", container.Image)
			}

			fmt.Printf("pod.status.phase: %s\n", pod.Status.Phase)
			for _, condition := range pod.Status.Conditions {
				fmt.Printf("condition.type: %s, condition.status: %s, conditon.reason: %s\n", condition.Type, condition.Status, condition.Reason)
			}

			for _, containerStatus := range pod.Status.ContainerStatuses {
				if containerStatus.State.Waiting != nil {
					fmt.Printf("containerStatus.state.waiting.reason: %s\n", containerStatus.State.Waiting.Reason)
				}
				if containerStatus.State.Running != nil {
					fmt.Printf("containerStatus.state.running.startedAt: %s\n", containerStatus.State.Running.StartedAt)
				}
			}
		}

		availableCondition := GetDeploymentCondition(k8sDeployment.Status, apps_v1.DeploymentAvailable)
		progressingCondition := GetDeploymentCondition(k8sDeployment.Status, apps_v1.DeploymentProgressing)
		if k8sDeployment.Status.UpdatedReplicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.Replicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.AvailableReplicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.ObservedGeneration >= k8sDeployment.Generation &&
			availableCondition.Status == "True" &&
			progressingCondition.Status == "True" {
			fmt.Printf("-------------deploy status------------\n")
			fmt.Println("success!")
		} else {
			fmt.Printf("-------------deploy status------------\n")
			fmt.Println("waiting...")
		}

		time.Sleep(3 * time.Second)
	}
}

// GetDeploymentCondition returns the condition with the provided type.
func GetDeploymentCondition(status apps_v1.DeploymentStatus, condType apps_v1.DeploymentConditionType) *apps_v1.DeploymentCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}

func GetPodCondition(status core_v1.PodStatus, condType core_v1.PodConditionType) *core_v1.PodCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}
