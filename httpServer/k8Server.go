package httpServer

import (
	"cm_platform/kubernetesApi"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var configPath = "./configPath/config"

// GetClusterInfoHandler 处理集群信息查询请求
func GetClusterInfoHandler(c *gin.Context) {
	//创建服务实例(kubernetesApi包下)
	clusterInfoService := kubernetesApi.ClusterInfoServiceImpl{}
	//执行查询
	clusterInfo, err := clusterInfoService.GetClusterInfo(c, kubernetesApi.K8sClientSet(configPath))
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
	//获取节点信息
	nodeInfo, err := nodeInfoServiceImpl.GetNodeInfo(c, kubernetesApi.K8sClientSet(configPath))
	if err != nil {
		log.Println("获取节点信息失败: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"nodeInfo": nodeInfo,
	})
}
