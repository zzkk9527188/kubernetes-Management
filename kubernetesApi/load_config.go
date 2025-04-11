package kubernetesApi

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

func K8sClientSet(config string) *kubernetes.Clientset {
	kubeConfig, err := LoadKubeConfig(config)
	if err != nil {
		log.Printf("配置加载失败: %v", err)
		return nil
	}
	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	return clientSet
}

func LoadKubeConfig(config string) (*rest.Config, error) {
	// 当传入自定义kubeconfig时
	if config != "" {
		// 使用新变量名避免遮蔽，同时处理错误
		kubeConfig, err := clientcmd.BuildConfigFromFlags("", config)
		if err != nil {
			return nil, fmt.Errorf("failed to load custom kubeconfig: %w", err)
		}
		return kubeConfig, nil
	}

	// 使用默认kubeconfig路径
	homeConfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	// 同样使用新变量名避免类型冲突
	defaultConfig, err := clientcmd.BuildConfigFromFlags("", homeConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to load default kubeconfig: %w", err)
	}
	return defaultConfig, nil
}
