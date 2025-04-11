package interfaces

import (
	"context"
	"k8s.io/client-go/kubernetes"
)

type ClusterInfoService interface {
	GetClusterInfo(ctx context.Context, clientSet *kubernetes.Clientset) (*[]ComponentsStatus, error)
	GetNodeInfo(ctx context.Context, clientSet *kubernetes.Clientset) (*[]NodeInfo, error)
}

// 定义 ComponentCondition 结构体
type ComponentCondition struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

// 定义 ComponentsStatus 结构体
type ComponentsStatus struct {
	Name       string               `json:"name"`
	Conditions []ComponentCondition `json:"conditions"`
}

type NodeInfo struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	Roles      string `json:"roles"`
	Age        string `json:"age"`
	Version    string `json:"version"`
	InternalIP string `json:"internalIP"`
	ExternalIP string `json:"externalIP"`
	OsImage    string `json:"osImage"`
	Kernel     string `json:"kernel"`
	Container  string `json:"container"`
}
