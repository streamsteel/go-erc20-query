package api

import (
	"web3-search/internal/ethereum"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有API路由
func RegisterRoutes(router *gin.Engine, ethClient *ethereum.Client) {
	handler := NewHandler(ethClient)

	// API版本分组
	v1 := router.Group("/api/v1")
	{
		// 健康检查
		v1.GET("/health", handler.HealthCheck)

		// ERC20代币相关路由
		token := v1.Group("/token")
		{
			// 获取代币信息
			token.GET("/:address", handler.GetTokenInfo)

			// 获取代币余额
			token.GET("/:tokenAddress/balance/:walletAddress", handler.GetTokenBalance)
		}

		// 以太坊相关路由
		eth := v1.Group("/eth")
		{
			// 获取ETH余额
			eth.GET("/balance/:address", handler.GetETHBalance)
		}
	}

	// 根路径重定向到健康检查
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/api/v1/health")
	})
}
