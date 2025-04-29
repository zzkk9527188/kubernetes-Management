package kubernetesApi

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

// DeploymentClient 创建实例
type DeploymentClient struct {
	//ns     string
	client kubernetes.Interface
}

// 初始化实例
func NewDeploymentClient() (*DeploymentClient, error) {
	clientSet, err := K8sClientSet()
	if err != nil {
		return nil, err
	}
	return &DeploymentClient{
		client: clientSet,
		//ns:     ns,
	}, nil

}

// ctx Go处理取消和超时的标准方式，创建deployment
func (d *DeploymentClient) Create(ctx context.Context, deployment *v1.Deployment, opts metav1.CreateOptions) (*v1.Deployment, error) {
	// 参数校验
	if deployment == nil {
		return nil, fmt.Errorf("deployment cannot be nil")
	}
	// 自动补全名称空间"default"
	if deployment.Namespace == "" {
		deployment.Namespace = metav1.NamespaceDefault
	}

	// 记录操作日志
	//klog.Infof("Create deployment in namespace: %s", d.ns)

	defer func() {
		klog.Infof("Create deployment operation completed for: %s/%s", deployment.Namespace, deployment.Name)
	}()
	// 确保标签存在
	if deployment.Labels == nil {
		deployment.Labels = make(map[string]string)
	}

	if _, exists := deployment.Labels[createorID]; !exists {
		deployment.Labels[createorID] = "cm-Visionary"
		deployment.Spec.Template.Labels[createorID] = "cm-Visionary"
	}
	// 调用Kubernetes API
	createdDeployment, err := d.client.AppsV1().Deployments("default").Create(ctx, deployment, opts)
	//错误处理
	if err != nil {
		if errors.IsAlreadyExists(err) {
			klog.Warningf("Deployment already exists: %s/%s", deployment.Namespace, deployment.Name)
			return nil, fmt.Errorf("deployment %s already exists in namespace %s: %w", deployment.Name, deployment.Namespace, err)
		}
		return nil, fmt.Errorf("failed to create deployment: %w", err)
	}
	return createdDeployment, nil
}

// 删除指定名称空间下的pod
func (d *DeploymentClient) DeleteNamespaced(ctx context.Context, namespace string, name string, opts *metav1.DeleteOptions) error {
	if name == "" || namespace == "" {
		return fmt.Errorf("deployment name or namespace cannot be empty")
	}
	// 记录操作日志
	klog.Infof("Delete deployment in namespace: %s/%s", namespace, name)
	defer func() {
		klog.Infof("Delete deployment operation completed for: %s/%s", namespace, name)
	}()
	//错误处理
	err := d.client.AppsV1().Deployments(namespace).Delete(ctx, name, *opts)
	if err != nil {
		if errors.IsNotFound(err) {
			klog.V(2).Infof("Deployment not found: %s/%s", namespace, name)
		}
		return fmt.Errorf("failed to delete deployment %s/%s: %w", namespace, name, err)
	}
	return nil
}

// replicas副本数为*int32类型，指针转化
func int32Ptr(i int32) *int32 {
	return &i
}
