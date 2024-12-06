package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"modern-stack/internal/config"
	"modern-stack/internal/models"
)

var DB *gorm.DB

func InitDB(cfg config.DatabaseConfig) *gorm.DB {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Database initialization failed: %v", r)
		}
	}()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// 添加连接池监控
	go func() {
		ticker := time.NewTicker(time.Minute)
		for range ticker.C {
			stats := sqlDB.Stats()
			log.Printf("DB Stats - Open: %d, Idle: %d, InUse: %d",
				stats.OpenConnections,
				stats.Idle,
				stats.InUse,
			)
		}
	}()

	// 自动迁移
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	return db
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized")
	}
	return DB
} 