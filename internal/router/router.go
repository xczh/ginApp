package router

import (
	"app/internal/config"
	"app/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// GinEngine 返回gin实例，它包含所有已注册的路由
func GinEngine() *gin.Engine {
	if log.IsDebug() {
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.Default()

	// 405 Handler
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, &ResponseMessage{"Method not allowed"})
	})

	// 404 Handler
	engine.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, &ResponseMessage{"Not found"})
	})

	router := engine.Group(config.GetURLPrefix())
	{
		// 注册pprof路由
		RegistPprofRoute(router)

		router.Use(TrafficControl(), APIv1Header())

		router.GET("/ping", Ping())
	}

	return engine
}
