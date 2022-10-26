package jwtauth

import (
	"backend/core/sdk/pkg"

	"backend/core/sdk/config"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ILogin interface {
	Verify(tx *gorm.DB) (MapClaims, error)
}

type IJWTAuth interface {
	Authenticator(c *gin.Context) (interface{}, error)
	PayloadFunc(data interface{}) MapClaims
	IdentityHandler(c *gin.Context) interface{}
	// LogOut(c *gin.Context)
	Authorizator(data interface{}, c *gin.Context) bool
	Unauthorized(c *gin.Context, code int, message string)
}

// AuthInit jwt验证new
func NewJWT(jwtAuth IJWTAuth) (*GinJWTMiddleware, error) {
	timeout := time.Hour
	if config.ApplicationConfig.Mode == pkg.Develop {
		timeout = time.Duration(876010) * time.Hour
	} else {
		if config.JwtConfig.Timeout != 0 {
			timeout = time.Duration(config.JwtConfig.Timeout) * time.Second
		}
	}
	middleware := GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte(config.JwtConfig.Secret),
		Timeout:         timeout,
		MaxRefresh:      time.Hour,
		PayloadFunc:     jwtAuth.PayloadFunc,
		IdentityHandler: jwtAuth.IdentityHandler,
		Authenticator:   jwtAuth.Authenticator,
		Authorizator:    jwtAuth.Authorizator,
		Unauthorized:    jwtAuth.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	}
	err := middleware.MiddlewareInit()
	return &middleware, err
}
