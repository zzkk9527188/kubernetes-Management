package httpServer

import (
	"cm_platform/kubernetesApi"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"log"
	"net/http"
	"time"
)

// GetClusterInfoHandler 处理集群信息查询请求
func GetClusterInfoHandler(c *gin.Context) {
	//创建服务实例(kubernetesApi包下)
	clusterInfoService := kubernetesApi.ClusterInfoServiceImpl{}

	clientSet, err := kubernetesApi.K8sClientSet()
	if err != nil {
		log.Println("clientSet创建失败", err)
	}
	//执行查询
	clusterInfo, err := clusterInfoService.GetClusterInfo(c, clientSet)
	if err != nil {
		log.Println("获取集群信息失败: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"information": clusterInfo,
	})
}

func GetNodeInfoHandler(c *gin.Context) {
	//创建服务实例(kubernetesApi包下)
	nodeInfoServiceImpl := kubernetesApi.NodeInfoServiceImpl{}
	//获取clientSet
	clientSet, err := kubernetesApi.K8sClientSet()
	if err != nil {
		log.Println("clientSet创建失败", err)
	}
	//获取节点信息
	nodeInfo, err := nodeInfoServiceImpl.GetNodeInfo(c, clientSet)
	if err != nil {
		log.Println("获取节点信息失败: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get node info"})
	}
	//获取节点成功
	c.JSON(http.StatusOK, gin.H{
		"nodeInfo": nodeInfo,
	})
}

// 创建deployment Handler
func CreateDeploymentHandler(c *gin.Context) {
	var req CreateDeploymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("请求绑定失败...")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
	}

	log.Println(req)
	//构建deployment对象
	obj := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
			Labels:    req.Labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &req.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: req.Labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: req.Labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  req.Name,
							Image: req.Image,
						},
					},
				},
			},
		},
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 15*time.Second)
	defer cancel()
	//初始化DeploymentClient，调用create方法
	clientSet, err := kubernetesApi.NewDeploymentClient()
	if err != nil {
		log.Printf(err.Error())
	}

	d, err := clientSet.Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		klog.V(1).ErrorS(err, "Create failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	c.JSON(http.StatusCreated, gin.H{
		"deployment": fmt.Sprintf("%s/%s", d.Namespace, d.Name),
		"status":     "success",
	})
}

func DeleteDeploymentHandler(c *gin.Context) {
	var req CreateDeploymentRequest
	deploymentClient, err := kubernetesApi.NewDeploymentClient()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 15*time.Second)
	defer cancel()
	err = deploymentClient.DeleteNamespaced(ctx, req.Namespace, req.Name, &metav1.DeleteOptions{})
	if err != nil {
		klog.V(1).ErrorS(err, "Delete failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete deployment success"})
}

func HandlerCreateDeployment(c *gin.Context) {
	var deployment appsv1.Deployment
	//解析请求
	if err := c.ShouldBindJSON(&deployment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid Json format",
			"details": err.Error(),
		})
	}
	//设置上下文选项
	ctx := context.Background()

	client, err := kubernetesApi.NewDeploymentClient()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	d, err := client.Create(ctx, &deployment, metav1.CreateOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create deployment", "details": err.Error()})
		klog.V(1).ErrorS(err, "Deployment create failed")
		return
	}

	//创建成功响应
	c.JSON(http.StatusCreated, gin.H{"deployment": fmt.Sprintf("%s/%s", d.Namespace, d.Name), "status": "success"})
}
