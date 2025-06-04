package api

import (
	"context"
	"net/http"
	"time"

	"web3-search/internal/ethereum"

	"github.com/gin-gonic/gin"
)

// Handler API处理器结构
type Handler struct {
	ethClient *ethereum.Client
}

// NewHandler 创建新的API处理器
func NewHandler(ethClient *ethereum.Client) *Handler {
	return &Handler{
		ethClient: ethClient,
	}
}

// Response 通用响应结构
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// GetTokenInfo 获取ERC20代币信息
func (h *Handler) GetTokenInfo(c *gin.Context) {
	tokenAddress := c.Param("address")
	if tokenAddress == "" {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Token address is required",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tokenInfo, err := h.ethClient.GetTokenInfo(ctx, tokenAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    tokenInfo,
	})
}

// GetTokenBalance 获取代币余额
func (h *Handler) GetTokenBalance(c *gin.Context) {
	tokenAddress := c.Param("tokenAddress")
	walletAddress := c.Param("walletAddress")

	if tokenAddress == "" || walletAddress == "" {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Both token address and wallet address are required",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	balanceInfo, err := h.ethClient.GetTokenBalance(ctx, tokenAddress, walletAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    balanceInfo,
	})
}

// GetETHBalance 获取ETH余额
func (h *Handler) GetETHBalance(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Address is required",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	balance, err := h.ethClient.GetETHBalance(ctx, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data: map[string]interface{}{
			"address": address,
			"balance": balance.String(),
			"unit":    "wei",
		},
	})
}

// HealthCheck 健康检查
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data: map[string]interface{}{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
			"service":   "web3-search",
		},
	})
}
