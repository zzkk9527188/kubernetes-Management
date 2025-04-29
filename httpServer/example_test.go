package httpServer

import (
	"bytes"
	"cm_platform/kubernetesApi"
	"cm_platform/pkg/tools"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDeploymentHandler(t *testing.T) {
	// 注册路由
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// 包装 CreateDeploymentHandler 为符合 Gin 路由处理函数签名的形式
	r.POST("/api/v1/k8s/deployment/create1", CreateDeploymentHandler)

	// 构建请求体
	requestBody := map[string]interface{}{
		"namespace": "default",
		"name":      "test-4",
		"replicas":  1,
		"image":     "10.10.10.49:5000/test/nginx:latest",
		"labels": map[string]string{
			"app":     "nginx",
			"version": "v1",
		},
	}

	reqBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// 创建测试请求
	req := httptest.NewRequest("POST", "/api/v1/k8s/deployment/create", bytes.NewBuffer(reqBytes))
	req.Header.Set("Content-Type", "application/json")

	// 创建响应记录器
	w := httptest.NewRecorder()

	// 执行请求
	r.ServeHTTP(w, req)

	// 验证响应状态码
	assert.Equal(t, http.StatusCreated, w.Code)

	// 检查 Deployment 是否被创建
	cli, err := kubernetesApi.K8sClientSet()
	if err != nil {
		t.Fatalf("Failed to create Kubernetes client: %v", err)
	}
	list, err := cli.AppsV1().Deployments("default").Get(context.Background(), "test-4", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get deployment: %v", err)
	}

	// 验证返回的 Deployment 对象是否符合预期
	assert.NotNil(t, list)
	assert.Equal(t, "test-4", list.Name)
	assert.Equal(t, int32(1), *list.Spec.Replicas)
	assert.Equal(t, "10.10.10.49:5000/test/nginx:latest", list.Spec.Template.Spec.Containers[0].Image)

	tools.ProcessBar()

	// 清理资源
	defer func() {
		err := cli.AppsV1().Deployments("default").Delete(context.Background(), "test-4", metav1.DeleteOptions{})
		if err != nil {
			t.Logf("Failed to delete deployment: %v", err)
		}
	}()
}

func TestProcessBar(t *testing.T) {
	tools.ProcessBar()
}
