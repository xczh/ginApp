package app

import (
	"app/internal/config"
	"app/internal/log"
	"app/internal/model"
	"app/internal/redis"
	"app/internal/router"
	"context"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/spf13/pflag"
)

func Run(fs *pflag.FlagSet) {
	defer log.Terminate()

	// 初始化config
	cfgFilePath, _ := fs.GetString("config")
	if err := config.Initialize(cfgFilePath); err != nil {
		log.Logger().Fatal("Config initialize failed")
	}

	// 初始化mysql
	if err := model.Initialize(); err != nil {
		log.Logger().Fatal("MySQL initialize failed: ", err)
	}
	defer model.Terminate()

	// 初始化redis
	if err := redis.Initialize(); err != nil {
		log.Logger().Fatal("Redis initialize failed: ", err)
	}
	defer redis.Terminate()

	// 测试配置文件
	if testConfigOnly, _ := fs.GetBool("test"); testConfigOnly {
		log.Logger().Info("Test config ok.")
		return
	}

	// 启动HTTP服务
	server := &http.Server{
		Addr:    config.GetAddress(),
		Handler: router.GinEngine(),
	}
	go server.ListenAndServe()

	// 信号处理
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	runtime.GC()

	for {
		switch sig := <-sigChan; sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Logger().Infof("Receive signal '%s', shutting down server...", sig.String())
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				log.Logger().Error("Server shutdown error: ", err)
			} else {
				log.Logger().Info("Server shutdown.")
			}
			return
		}
	}
}
