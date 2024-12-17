# BookPavilion (图书阅读馆)

基于Vue.js和Go的在线图书存储与阅读服务。

## 项目概述

BookPavilion是一个在线图书存储与阅读服务平台，为用户提供图书上传、存储、在线阅读等功能。

## 开发计划

### 第一阶段：基础框架搭建
- [x] 项目结构初始化
- [ ] 数据库设计
- [ ] 基础API框架搭建

### 第二阶段：用户模块（进行中）
- [ ] 用户注册功能
  - 基本信息注册（用户名、邮箱）
  - 数据验证
  - 单元测试
- [ ] 用户信息管理
  - 查看个人信息
  - 修改个人信息

### 第三阶段：图书管理模块
- [ ] 图书上传功能
  - 支持多种格式（PDF、EPUB、TXT、MOBI）
  - 文件验证
  - 上传进度显示
- [ ] 图书元数据提取
  - 自动提取书籍信息
  - 支持手动编辑
- [ ] 图书分类管理
  - 自定义分类
  - 标签管理

### 第四阶段：阅读功能模块
- [ ] 在线阅读器
  - PDF阅读支持
  - EPUB阅读支持
  - 阅读进度保存
- [ ] 阅读工具
  - 书签功能
  - 笔记功能
  - 文字高亮

### 第五阶段：存储管理模块
- [ ] 文件存储系统
  - 本地存储
  - 云存储支持
- [ ] 存储空间管理
  - 空间使用统计
  - 配额管理

## 项目结构

```
bookpavilion/
├── backend/                 # Go后端项目
│   ├── main.go             # 主程序入口
│   ├── config/             # 配置文件
│   ├── models/             # 数据模型
│   ├── controllers/        # 控制器
│   ├── services/          # 业务逻辑
│   └── tests/             # 测试文件
│
├── frontend/               # Vue前端项目
│   ├── src/
│   │   ├── components/    # 组件
│   │   ├── views/        # 页面
│   │   ├── router/       # 路由配置
│   │   ├── store/        # 状态管理
│   │   └── assets/       # 静态资源
│   └── tests/            # 测试文件
│
└── docs/                  # 项目文档
```

## 技术栈

### 后端
- 语言：Go
- Web框架：Gin
- 数据库：MySQL
- ORM：GORM

### 前端
- 框架：Vue 3
- 状态管理：Vuex
- 路由：Vue Router
- UI组件库：Element Plus
- HTTP客户端：Axios

## 数据库设计

### 用户表 (users)
```sql
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### 图书表 (books)
```sql
CREATE TABLE books (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(200) NOT NULL,
    author VARCHAR(100),
    format VARCHAR(10),
    file_path VARCHAR(500),
    user_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

### 阅读进度表 (reading_progress)
```sql
CREATE TABLE reading_progress (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    book_id INT,
    progress VARCHAR(50),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (book_id) REFERENCES books(id)
);
```

## API设计

### 用户模块
- POST /api/users - 创建新用户
- GET /api/users/:id - 获取用户信息
- PUT /api/users/:id - 更新用户信息

### 图书模块
- POST /api/books - 上传图书
- GET /api/books - 获取图书列表
- GET /api/books/:id - 获取图书详情
- PUT /api/books/:id - 更新图书信息
- DELETE /api/books/:id - 删除图书

### 阅读进度模块
- POST /api/progress - 保存阅读进度
- GET /api/progress/:bookId - 获取阅读进度

## 开发环境搭建

### 后端环境要求
- Go >= 1.21
- MySQL >= 8.0

### 前端环境要求
- Node.js >= 16
- npm >= 8

### 开发环境配置步骤
1. 克隆项目
```bash
git clone https://github.com/yourusername/bookpavilion.git
cd bookpavilion
```

2. 配置后端
```bash
cd backend
go mod tidy
```

3. 配置前端
```bash
cd frontend
npm install
```

4. 配置数据库
```bash
# 创建数据库
mysql -u root -p
CREATE DATABASE bookpavilion;
```

## 测试

### 后端测试
```bash
cd backend
go test ./...
```

### 前端测试
```bash
cd frontend
npm run test
```

## 开发规范

### 代码规范
- 使用gofmt格式化Go代码
- 使用ESLint规范JavaScript代码
- 所有代码必须包含适当的注释
- 所有新功能必须包含单元测试

### Git提交规范
- feat: 新功能
- fix: 修复bug
- docs: 文档更新
- style: 代码格式修改
- refactor: 代码重构
- test: 测试用例修改
- chore: 其他修改

## 部署

### 开发环境
- 前端: `npm run serve`
- 后端: `go run main.go`

### 生产环境
- 前端: `npm run build`
- 后端: `go build`

## 贡献指南
1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 版本历史
- 0.1.0 (开发中)
  - 项目初始化
  - 基础框架搭建
  - 用户模块开发

## 作者
- [作者名称]

## 许可证
MIT License
