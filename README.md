# Web3 ERC-20链上数据查询服务

一个基于Go语言开发的Web3服务，用于查询以太坊网络上的ERC-20代币信息和余额。

## 🚀 功能特性

- **ERC-20代币信息查询**: 获取代币名称、符号、小数位数、总供应量
- **余额查询**: 查询指定地址的ERC-20代币余额和ETH余额
- **实时数据**: 直接从以太坊网络获取最新数据
- **RESTful API**: 提供简洁的HTTP API接口
- **Docker支持**: 容器化部署，开箱即用
- **健康检查**: 内置服务健康监控

## 🛠 技术栈

- **后端**: Go 1.21 + Gin框架
- **区块链**: go-ethereum SDK
- **容器化**: Docker + Docker Compose
- **智能合约**: Solidity (测试合约)

## 📦 快速开始

### 环境要求

- Go 1.21+
- Docker & Docker Compose (可选)
- 以太坊RPC节点访问权限 (Infura/Alchemy等)

### 1. 克隆项目

\`\`\`bash
git clone <repository-url>
cd web3-search
\`\`\`

### 2. 配置环境变量

复制配置文件并修改：

\`\`\`bash
cp config.env.example .env
\`\`\`

编辑 \`.env\` 文件，设置你的以太坊RPC URL：

\`\`\`env
ETH_RPC_URL=https://sepolia.infura.io/v3/YOUR_PROJECT_ID
PORT=8080
GIN_MODE=release
\`\`\`

### 3. 运行方式

#### 方式一：直接运行

\`\`\`bash
# 安装依赖
go mod tidy

# 运行服务
go run main.go
\`\`\`

#### 方式二：Docker运行

\`\`\`bash
# 构建并启动
docker-compose up --build

# 后台运行
docker-compose up -d --build
\`\`\`

服务启动后，访问 http://localhost:8080 查看健康状态。

## 📚 API文档

### 基础信息

- **Base URL**: \`http://localhost:8080/api/v1\`
- **Content-Type**: \`application/json\`

### 接口列表

#### 1. 健康检查

\`\`\`
GET /api/v1/health
\`\`\`

**响应示例**:
\`\`\`json
{
  "success": true,
  "data": {
    "status": "healthy",
    "timestamp": 1703123456,
    "service": "web3-search"
  }
}
\`\`\`

#### 2. 获取ERC-20代币信息

\`\`\`
GET /api/v1/token/{tokenAddress}
\`\`\`

**参数**:
- \`tokenAddress\`: ERC-20代币合约地址

**响应示例**:
\`\`\`json
{
  "success": true,
  "data": {
    "name": "USD Coin",
    "symbol": "USDC",
    "decimals": 6,
    "totalSupply": "1000000000000000",
    "address": "0xA0b86a33E6441b8435b662da0C0C8C2A8d3d2074"
  }
}
\`\`\`

#### 3. 获取ERC-20代币余额

\`\`\`
GET /api/v1/token/{tokenAddress}/balance/{walletAddress}
\`\`\`

**参数**:
- \`tokenAddress\`: ERC-20代币合约地址
- \`walletAddress\`: 钱包地址

**响应示例**:
\`\`\`json
{
  "success": true,
  "data": {
    "address": "0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6",
    "tokenAddress": "0xA0b86a33E6441b8435b662da0C0C8C2A8d3d2074",
    "balance": "1000000000",
    "decimals": 6
  }
}
\`\`\`

#### 4. 获取ETH余额

\`\`\`
GET /api/v1/eth/balance/{address}
\`\`\`

**参数**:
- \`address\`: 钱包地址

**响应示例**:
\`\`\`json
{
  "success": true,
  "data": {
    "address": "0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6",
    "balance": "1000000000000000000",
    "unit": "wei"
  }
}
\`\`\`

## 🧪 测试

### 使用测试代币

项目包含一个测试用的ERC-20代币合约 (\`contracts/TestToken.sol\`)，你可以部署到测试网进行测试。

### API测试示例

\`\`\`bash
# 健康检查
curl http://localhost:8080/api/v1/health

# 查询USDC代币信息 (Sepolia测试网)
curl http://localhost:8080/api/v1/token/0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238

# 查询代币余额
curl http://localhost:8080/api/v1/token/0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238/balance/0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6

# 查询ETH余额
curl http://localhost:8080/api/v1/eth/balance/0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6
\`\`\`

## 🏗 项目结构

\`\`\`
web3-search/
├── main.go                 # 程序入口
├── go.mod                  # Go模块定义
├── go.sum                  # 依赖校验
├── Dockerfile              # Docker镜像构建
├── docker-compose.yml      # Docker Compose配置
├── config.env.example      # 环境变量示例
├── README.md              # 项目文档
├── contracts/             # 智能合约
│   └── TestToken.sol      # 测试ERC-20合约
└── internal/              # 内部包
    ├── api/               # API层
    │   ├── handlers.go    # 请求处理器
    │   └── routes.go      # 路由定义
    ├── config/            # 配置管理
    │   └── config.go      # 配置加载
    └── ethereum/          # 以太坊客户端
        └── client.go      # 区块链交互
\`\`\`

## 🔧 开发指南

### 添加新的API接口

1. 在 \`internal/ethereum/client.go\` 中添加区块链交互方法
2. 在 \`internal/api/handlers.go\` 中添加HTTP处理器
3. 在 \`internal/api/routes.go\` 中注册新路由

### 支持新的代币标准

可以扩展 \`ethereum/client.go\` 中的ABI定义，支持ERC-721、ERC-1155等其他标准。

## 🚨 注意事项

1. **RPC限制**: 免费的RPC服务通常有请求频率限制
2. **网络选择**: 确保RPC URL对应正确的网络（主网/测试网）
3. **地址格式**: 所有地址参数必须是有效的以太坊地址格式
4. **错误处理**: API会返回详细的错误信息，便于调试

## 📄 许可证

MIT License

## 🤝 贡献

欢迎提交Issue和Pull Request！

## 📞 联系方式

如有问题，请通过GitHub Issues联系。 