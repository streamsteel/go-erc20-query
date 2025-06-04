package ethereum

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client 以太坊客户端封装
type Client struct {
	client   *ethclient.Client
	erc20ABI abi.ABI
}

// ERC20 代币信息结构
type TokenInfo struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimals    uint8  `json:"decimals"`
	TotalSupply string `json:"totalSupply"`
	Address     string `json:"address"`
}

// BalanceInfo 余额信息结构
type BalanceInfo struct {
	Address      string `json:"address"`
	TokenAddress string `json:"tokenAddress"`
	Balance      string `json:"balance"`
	Decimals     uint8  `json:"decimals"`
}

// ERC20 ABI (简化版，包含常用方法)
const erc20ABI = `[
	{
		"constant": true,
		"inputs": [],
		"name": "name",
		"outputs": [{"name": "", "type": "string"}],
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "symbol",
		"outputs": [{"name": "", "type": "string"}],
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "decimals",
		"outputs": [{"name": "", "type": "uint8"}],
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "totalSupply",
		"outputs": [{"name": "", "type": "uint256"}],
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [{"name": "_owner", "type": "address"}],
		"name": "balanceOf",
		"outputs": [{"name": "balance", "type": "uint256"}],
		"type": "function"
	}
]`

// NewClient 创建新的以太坊客户端
func NewClient(rpcURL string) (*Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	// 解析ERC20 ABI
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		return nil, err
	}

	return &Client{
		client:   client,
		erc20ABI: parsedABI,
	}, nil
}

// Close 关闭客户端连接
func (c *Client) Close() {
	c.client.Close()
}

// GetTokenInfo 获取ERC20代币信息
func (c *Client) GetTokenInfo(ctx context.Context, tokenAddress string) (*TokenInfo, error) {
	address := common.HexToAddress(tokenAddress)

	// 获取代币名称
	name, err := c.callContract(ctx, address, "name")
	if err != nil {
		return nil, err
	}

	// 获取代币符号
	symbol, err := c.callContract(ctx, address, "symbol")
	if err != nil {
		return nil, err
	}

	// 获取小数位数
	decimalsResult, err := c.callContract(ctx, address, "decimals")
	if err != nil {
		return nil, err
	}

	// 获取总供应量
	totalSupplyResult, err := c.callContract(ctx, address, "totalSupply")
	if err != nil {
		return nil, err
	}

	// 解析结果
	var tokenName string
	var tokenSymbol string
	var decimals uint8
	var totalSupply *big.Int

	if err := c.erc20ABI.UnpackIntoInterface(&tokenName, "name", name); err != nil {
		return nil, err
	}

	if err := c.erc20ABI.UnpackIntoInterface(&tokenSymbol, "symbol", symbol); err != nil {
		return nil, err
	}

	if err := c.erc20ABI.UnpackIntoInterface(&decimals, "decimals", decimalsResult); err != nil {
		return nil, err
	}

	if err := c.erc20ABI.UnpackIntoInterface(&totalSupply, "totalSupply", totalSupplyResult); err != nil {
		return nil, err
	}

	return &TokenInfo{
		Name:        tokenName,
		Symbol:      tokenSymbol,
		Decimals:    decimals,
		TotalSupply: totalSupply.String(),
		Address:     tokenAddress,
	}, nil
}

// GetTokenBalance 获取地址的代币余额
func (c *Client) GetTokenBalance(ctx context.Context, tokenAddress, walletAddress string) (*BalanceInfo, error) {
	tokenAddr := common.HexToAddress(tokenAddress)
	walletAddr := common.HexToAddress(walletAddress)

	// 准备balanceOf调用数据
	data, err := c.erc20ABI.Pack("balanceOf", walletAddr)
	if err != nil {
		return nil, err
	}

	// 调用合约
	result, err := c.client.CallContract(ctx, ethereum.CallMsg{
		To:   &tokenAddr,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	// 解析余额
	var balance *big.Int
	if err := c.erc20ABI.UnpackIntoInterface(&balance, "balanceOf", result); err != nil {
		return nil, err
	}

	// 获取小数位数
	decimalsResult, err := c.callContract(ctx, tokenAddr, "decimals")
	if err != nil {
		return nil, err
	}

	var decimals uint8
	if err := c.erc20ABI.UnpackIntoInterface(&decimals, "decimals", decimalsResult); err != nil {
		return nil, err
	}

	return &BalanceInfo{
		Address:      walletAddress,
		TokenAddress: tokenAddress,
		Balance:      balance.String(),
		Decimals:     decimals,
	}, nil
}

// GetETHBalance 获取ETH余额
func (c *Client) GetETHBalance(ctx context.Context, address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := c.client.BalanceAt(ctx, account, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// callContract 调用合约方法的辅助函数
func (c *Client) callContract(ctx context.Context, contractAddress common.Address, method string, args ...interface{}) ([]byte, error) {
	data, err := c.erc20ABI.Pack(method, args...)
	if err != nil {
		return nil, err
	}

	result, err := c.client.CallContract(ctx, ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}
