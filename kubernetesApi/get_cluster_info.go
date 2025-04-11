package kubernetesApi

import (
	"cm_platform/interfaces"
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"strings"
	"time"
)

// 实现 ClusterInfoService 接口的结构体
type ClusterInfoServiceImpl struct{}

// NewClusterInfoService 创建服务实例
func NewClusterInfoService() *ClusterInfoServiceImpl {
	return &ClusterInfoServiceImpl{}
}

type NodeInfoServiceImpl struct {
}

func (s *ClusterInfoServiceImpl) GetClusterInfo(ctx context.Context, clientSet *kubernetes.Clientset) (*[]interfaces.ComponentsStatus, error) {
	list, err := clientSet.CoreV1().ComponentStatuses().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	// 转换 Kubernetes API 响应到自定义结构体

	var componentsStatuses []interfaces.ComponentsStatus

	for _, component := range list.Items {
		var conditions []interfaces.ComponentCondition
		for _, condition := range component.Conditions {
			conditions = append(conditions, interfaces.ComponentCondition{
				Type:    string(condition.Type),
				Status:  string(condition.Status),
				Message: condition.Message,
				Error:   condition.Error,
			})
		}
		componentsStatuses = append(componentsStatuses, interfaces.ComponentsStatus{
			Name:       component.Name,
			Conditions: conditions,
		})
	}
	log.Printf("Cluster component statuses: %+v", componentsStatuses)
	return &componentsStatuses, nil
}

func (n *NodeInfoServiceImpl) GetNodeInfo(ctx context.Context, clientSet *kubernetes.Clientset) (*[]interfaces.NodeInfo, error) {
	nodes, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("获取节点信息失败: %v", err)
		return &[]interfaces.NodeInfo{}, err
	}
	var nodeInfoList []interfaces.NodeInfo
	for _, node := range nodes.Items {
		//获取节点状态,没获取到默认为未知
		status := corev1.ConditionUnknown // 默认为未知，在源码中const定义
		for _, nodeStatus := range node.Status.Conditions {
			if nodeStatus.Type == corev1.NodeReady {
				if nodeStatus.Status == corev1.ConditionTrue {
					status = "Ready"
				} else {
					status = "NotReady"
				}
				break
			}
		}
		nodeRoles := []string{}
		if _, exists := node.Labels["node-role.kubernetes.io/master"]; exists {
			nodeRoles = append(nodeRoles, "master")
		}

		if _, exists := node.Labels["node-role.kubernetes.io/control-plane"]; exists {
			nodeRoles = append(nodeRoles, "control-plane")
		}
		if _, exists := node.Labels["node-role.kubernetes.io/worker"]; exists {
			nodeRoles = append(nodeRoles, "worker")
		}
		roles := strings.Join(nodeRoles, ",")

		//获取ip
		internalIP := ""
		externalIP := "None"
		for _, address := range node.Status.Addresses {
			if address.Type == corev1.NodeInternalIP {
				internalIP = address.Address
			} else if address.Type == corev1.NodeExternalIP {
				externalIP = address.Address
			}
		}
		nodeInfo := interfaces.NodeInfo{
			Name:       node.Name,
			Status:     string(status),
			Age:        fmt.Sprintf("%d days", int(time.Since(node.CreationTimestamp.Time).Hours()/24)),
			Roles:      roles,
			InternalIP: internalIP,
			ExternalIP: externalIP,
			Version:    node.Status.NodeInfo.KubeletVersion,
			Kernel:     node.Status.NodeInfo.KernelVersion,
			OsImage:    node.Status.NodeInfo.OSImage,
			Container:  node.Status.NodeInfo.ContainerRuntimeVersion,
		}
		nodeInfoList = append(nodeInfoList, nodeInfo)

	}
	return &nodeInfoList, nil
}
