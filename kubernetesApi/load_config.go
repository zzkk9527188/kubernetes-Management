package kubernetesApi

import (
	"cm_platform/internal/config"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func K8sClientSet() (*kubernetes.Clientset, error) {
	kubeConfig, err := LoadKubeConfig()
	if err != nil {
		log.Printf("配置加载失败: %v", err)
		return nil, err
	}
	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	return clientSet, nil
}

func LoadKubeConfig() (*rest.Config, error) {
	//获取配置文件目录
	//currentDir, err := os.Getwd()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//platform := fmt.Sprintf("%s/configPath/cm_platform.yaml", currentDir)
	platform := "E:\\StudyCode\\go_code\\cm_platform\\cmd\\configPath\\cm_platform.yaml"
	cm, err := config.ViperLoadConfig(platform)
	if cm == nil {
		return nil, fmt.Errorf("配置文件信息获取失败")
	}
	//fmt.Println(cm.KubeVisionary.CmPlatform)
	if err != nil {
		log.Printf("err: %v", err)
	}

	kubeConfig := cm.KubeVisionary.KubeConfig
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		log.Printf("err: %v", err)
		return nil, err
	}

	return restConfig, nil
}
