package plugs

import (
	"github.com/gin-gonic/gin"
)

// Cross is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.Cross domain
func Cross(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		// c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

// CorsByRules 按照配置处理跨域请求
func CrossByRules() gin.HandlerFunc {
	// 放行全部
	// if global.GVA_CONFIG.Cors.Mode == "allow-all" {
	// 	return Cross
	// }
	return func(c *gin.Context) {
		// whitelist := checkCors(c.GetHeader("origin"))

		// 通过检查, 添加请求头
		// if whitelist != nil {
		// 	c.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
		// 	c.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
		// 	c.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
		// 	c.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
		// 	if whitelist.AllowCredentials {
		// 		c.Header("Access-Control-Allow-Credentials", "true")
		// 	}
		// }

		// 严格白名单模式且未通过检查，直接拒绝处理请求
		// if whitelist == nil && global.GVA_CONFIG.Cors.Mode == "strict-whitelist" && !(c.Request.Method == "GET" && c.Request.URL.Path == "/health") {
		// 	c.AbortWithStatus(http.StatusForbidden)
		// } else {
		// 	// 非严格白名单模式，无论是否通过检查均放行所有 OPTIONS 方法
		// 	if c.Request.Method == "OPTIONS" {
		// 		c.AbortWithStatus(http.StatusNoContent)
		// 	}
		// }

		// 处理请求
		c.Next()
	}
}

// func checkCors(currentOrigin string) *config.CORSWhitelist {
// 	for _, whitelist := range global.GVA_CONFIG.Cors.Whitelist {
// 		// 遍历配置中的跨域头，寻找匹配项
// 		if currentOrigin == whitelist.AllowOrigin {
// 			return &whitelist
// 		}
// 	}
// 	return nil
// }
