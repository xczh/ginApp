package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 流量控制
func TrafficControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		if false {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, &ResponseMessage{"Too many requests"})
		}
	}
}

// 放置中间件
func APIv1Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Request必须设置User-Agent
		if ver := c.Request.Header.Get("User-Agent"); ver == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, &ResponseMessage{"Please make sure your request has a User-Agent header"})
			return
		}

		// Accept: application/vnd.app[.version].param[+type]
		if ver := c.Request.Header.Get("Accept"); ver != "" {
			if !strings.HasPrefix(ver, "application/vnd.app.v1") {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, &ResponseMessage{"Invalid Accept header"})
				return
			}
			ver = ver[len("application/vnd.app.v1"):]
			if l := len(ver); l > 0 {
				// parse type
				mark := strings.IndexByte(ver, '+')
				if mark >= 0 && mark+1 < l {
					switch ver[mark+1:] {
					case "json":
						c.Set("req.accept.type", "json")
					case "raw":
						c.Set("req.accept.type", "raw")
					default:
						c.AbortWithStatusJSON(http.StatusNotAcceptable, &ResponseMessage{"Invalid Accept header"})
						return
					}
				} else {
					// default type
					c.Set("req.accept.type", "json")
				}

				// 媒体类型参数
				if mark >= 0 {
					ver = ver[:mark]
				}
				// TODO: 暂时没有媒体类型参数
			}
		}
		c.Header("Connection", "keep-alive")
		c.Header("X-App-Media-Type", "app.v1")
		c.Header("Cache-Control", "max-age=0, private, must-revalidate")
		c.Header("X-Content-Type-Options", "nosniff")
		// CORS
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, If-Match, If-Modified-Since, If-None-Match, If-Unmodified-Since, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "ETag, Link, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval")
		c.Header("Access-Control-Max-Age", "86400")
	}
}

func Authentication(use404 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &ResponseMessage{"Bad credentials"})

		if use404 {
			c.AbortWithStatusJSON(http.StatusNotFound, &ResponseMessage{"Not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, &ResponseMessage{"Forbidden"})
		}
	}
}
