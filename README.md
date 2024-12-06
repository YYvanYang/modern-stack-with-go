# Modern Stack with Go

一个使用最新技术栈的全栈应用模板，采用 Go 后端和 Next.js 前端。

## 技术栈

### 后端
- Go 1.22
- Gin Web 框架
- GORM
- PostgreSQL
- JWT 认证

### 前端
- Next.js 15 (App Router)
- React 19
  - Server Components
  - Server Actions
  - Suspense
- TypeScript 5
- TailwindCSS 4
- Shadcn UI

## 特性
- 🚀 基于 Next.js 15 App Router 架构
- ⚡️ React 19 全新特性支持
  - Server Components
  - Server Actions
  - Suspense
  - Use Hook
- 🔒 JWT 认证和授权
- 🎨 基于 TailwindCSS 的响应式设计
- 📱 移动优先的 UI 设计
- 🔍 SEO 优化
- 🌐 API 路由优化
- 📦 Docker 容器化部署

## 项目结构

```
.
├── backend/                # Go 后端项目
│   ├── cmd/               # 主程序入口
│   ├── internal/          # 内部包
│   │   ├── api/          # API 路由和处理器
│   │   ├── config/       # 配置
│   │   ├── database/     # 数据库连接和模型
│   │   ├── middleware/   # 中间件
│   │   ├── models/       # 数据模型
│   │   ├── services/     # 业务逻辑
│   │   └── utils/        # 工具函数
│   └── tests/            # 测试
├── frontend/             # Next.js 前端项目
│   ├── app/             # App Router 页面
│   ├── components/      # React 组件
│   ├── lib/            # 工具函数
│   └── public/         # 静态资源
├── docker/             # Docker 配置
│   ├── Dockerfile.backend
│   └── Dockerfile.frontend
└── compose.yaml       # Docker Compose 配置
```

## 快速开始

### 使用 Docker Compose

1. 克隆项目并进入目录：
```bash
git clone https://github.com/your-username/modern-stack
cd modern-stack
```

2. 启动所有服务：
```bash
docker compose up -d
```

3. 访问应用：
- 前端: http://localhost:3000
- 后端: http://localhost:8080
- API 文档: http://localhost:8080/swagger/index.html

### 本地开发

#### 后端开发
1. 确保已安装 Go 1.22
2. 设置环境变量：
```bash
cp backend/.env.example backend/.env
```

3. 启动后端服务：
```bash
cd backend
go mod download
go run cmd/server/main.go
```

#### 前端开发
1. 确保已安装 Node.js 20
2. 安装依赖并启动：
```bash
cd frontend
npm install
npm run dev
```

## API 路由

### 认证相关
- POST `/api/v1/auth/register` - 用户注册
- POST `/api/v1/auth/login` - 用户登录

### 用户相关
- GET `/api/v1/user/profile` - 获取用户信息
- PUT `/api/v1/user/profile` - 更新用户信息

## 环境变量

### 后端环境变量
```env
DB_HOST=postgres
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=modern_stack
DB_PORT=5432
JWT_SECRET=your_jwt_secret_key_here
GIN_MODE=debug
```

### 前端环境变量
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## 部署

项目使用 Docker Compose 进行容器化部署，包含以下服务：
- `frontend`: Next.js 前端应用
- `backend`: Go 后端服务
- `postgres`: PostgreSQL 数据库

### 生产环境部署
```bash
docker compose -f compose.yaml up -d
```

## 开发规范

1. Git 提交信息格式：
```
<type>(<scope>): <subject>

<body>
```

2. 分支命名规范：
- 功能分支: `feature/功能名称`
- 修复分支: `fix/问题描述`
- 优化分支: `optimize/优化内容`

## 贡献指南

1. Fork 本仓库
2. 创建功能分支
3. 提交代码
4. 创建 Pull Request

## 许可证

MIT
