package main

import (
	"fmt"
	"go-web-app/logger"
	"go-web-app/settings"
	"go.uber.org/zap"
)

/**
go web 开发较通用的脚手架模板
*/
func main() {
	// 1. 加载配置(创建settings模块)
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	// 2. 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Println("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	// 3. 初始化MySQL连接
	// 4. 初始化redis连接
	// 5. 注册路由
	// 6. 启动服务(优雅关机)
}
