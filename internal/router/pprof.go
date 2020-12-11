package router

import (
	"net/http"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

// RegistPprofRoute 向路由组中注册pprof路由
func RegistPprofRoute(r gin.IRouter) {
	p := r.Group("/pprof")
	p.GET("/", ginHandler(pprof.Index))
	p.GET("/cmdline", ginHandler(pprof.Cmdline))
	p.GET("/profile", ginHandler(pprof.Profile))
	p.POST("/symbol", ginHandler(pprof.Symbol))
	p.GET("/symbol", ginHandler(pprof.Symbol))
	p.GET("/trace", ginHandler(pprof.Trace))
	p.GET("/allocs", ginHandler(pprof.Handler("allocs").ServeHTTP))
	p.GET("/block", ginHandler(pprof.Handler("block").ServeHTTP))
	p.GET("/goroutine", ginHandler(pprof.Handler("goroutine").ServeHTTP))
	p.GET("/heap", ginHandler(pprof.Handler("heap").ServeHTTP))
	p.GET("/mutex", ginHandler(pprof.Handler("mutex").ServeHTTP))
	p.GET("/threadcreate", ginHandler(pprof.Handler("threadcreate").ServeHTTP))
}

func ginHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
