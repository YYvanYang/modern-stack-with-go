package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"modern-stack/internal/api"
	"modern-stack/internal/config"
	"modern-stack/internal/database"
	"modern-stack/internal/middleware"
	"modern-stack/internal/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 加载配置
	cfg := config.Load()

	// 初始化数据库
	db := database.InitDB(cfg.Database)

	// 初始化服务
	authService := services.NewAuthService(db)
	userService := services.NewUserService(db)

	// 初始化路由
	router := api.NewRouter(cfg, authService, userService)

	// 创建 Gin 实例
	engine := gin.New()

	// 设置中间件
	m := middleware.NewMiddleware(cfg)
	m.Setup(engine)

	// 设置路由
	router.Setup(engine)

	// 启动服务器
	if err := engine.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
} 