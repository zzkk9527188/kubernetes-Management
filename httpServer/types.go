package httpServer

// 制定又cm平台创建
type creator string

const CreatorByCmVisionary creator = "cm.platform.visionary/creator"

type CreateDeploymentRequest struct {
	Namespace string            `json:"namespace" binding:"required"` // 命名空间
	Name      string            `json:"name" binding:"required"`      // Deployment 名称
	Image     string            `json:"image"`                        // 容器镜像
	Replicas  int32             `json:"replicas,omitempty"`           // 副本数
	Labels    map[string]string `json:"labels,omitempty"`             // 标签
}
