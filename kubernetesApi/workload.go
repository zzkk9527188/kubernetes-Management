package kubernetesApi

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const createorID = "field.cm.io/creatorId"

type DeploymentInterface interface {
	// 基础CRUD操作
	Create(ctx context.Context, deployment *v1.Deployment, opts metav1.CreateOptions) (*v1.Deployment, error)
	Get(ctx context.Context, deployment *v1.Deployment, opts metav1.GetOptions) (*v1.Deployment, error)
	Update(deployment *v1.Deployment) (*v1.Deployment, error)
	Delete(ctx context.Context, name string, opts *metav1.DeleteOptions) error
	// 命名空间相关操作
	GetNamespaced(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*v1.Deployment, error)
	DeleteNamespaced(ctx context.Context, namespace string, name string, opts *metav1.DeleteOptions) error
	// 列表操作
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DeploymentList, error)
	ListNamespaced(ctx context.Context, namespace string, opts metav1.ListOptions) (*v1.DeploymentList, error)
}
