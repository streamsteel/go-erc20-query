# Web3 Search Service Makefile

.PHONY: help build run test clean docker-build docker-run docker-stop deps fmt lint

# 默认目标
help:
	@echo "Available commands:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run with Docker Compose"
	@echo "  docker-stop  - Stop Docker containers"
	@echo "  deps         - Download dependencies"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"

# 构建应用
build:
	@echo "Building application..."
	go build -o main .

# 运行应用
run:
	@echo "Running application..."
	go run main.go

# 运行测试
test:
	@echo "Running tests..."
	go test -v ./...

# 清理构建产物
clean:
	@echo "Cleaning..."
	rm -f main
	go clean

# 构建Docker镜像
docker-build:
	@echo "Building Docker image..."
	docker-compose build

# 使用Docker运行
docker-run:
	@echo "Starting with Docker Compose..."
	docker-compose up -d

# 停止Docker容器
docker-stop:
	@echo "Stopping Docker containers..."
	docker-compose down

# 下载依赖
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# 格式化代码
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 运行linter
lint:
	@echo "Running linter..."
	golangci-lint run

# 开发环境设置
dev-setup: deps
	@echo "Setting up development environment..."
	cp config.env.example .env
	@echo "Please edit .env file with your configuration"

# 快速启动（开发模式）
dev: dev-setup run

# 生产环境构建
prod-build:
	@echo "Building for production..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -s' -o main .

# 显示项目状态
status:
	@echo "Project Status:"
	@echo "Go version: $(shell go version)"
	@echo "Git branch: $(shell git branch --show-current 2>/dev/null || echo 'N/A')"
	@echo "Git commit: $(shell git rev-parse --short HEAD 2>/dev/null || echo 'N/A')" 