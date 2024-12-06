package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	LogDir     = "logs"
	MaxLogSize = 10 * 1024 * 1024 // 10MB
)

func SetupLogger() *os.File {
	// 创建日志目录
	if err := os.MkdirAll(LogDir, 0755); err != nil {
		log.Fatal(err)
	}

	// 生成日志文件名
	timestamp := time.Now().Format("2006-01-02")
	logPath := filepath.Join(LogDir, fmt.Sprintf("server-%s.log", timestamp))

	// 创建或打开日志文件
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 设置日志输出
	log.SetOutput(logFile)
	gin.DefaultWriter = logFile

	return logFile
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		log.Printf("| %3d | %13v | %15s | %s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			method,
			path,
			errorMessage,
		)
	}
} 