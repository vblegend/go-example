package jwtauth

import (
	"server/sugar/env"

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

	LoginResponse(c *gin.Context, code int, token string, expire time.Time)
	RefreshResponse(c *gin.Context, code int, token string, expire time.Time)
}

// AuthInit jwt验证new
func NewJWT(jwtAuth IJWTAuth, _timeout int64, secret string) (*GinJWTMiddleware, error) {
	timeout := time.Hour
	if env.ModeIs(env.Develop) {
		timeout = time.Duration(876010) * time.Hour
	} else {
		if _timeout != 0 {
			timeout = time.Duration(_timeout) * time.Second
		}
	}
	middleware := GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte(secret),
		Timeout:         timeout,
		MaxRefresh:      time.Hour,
		PayloadFunc:     jwtAuth.PayloadFunc,
		IdentityHandler: jwtAuth.IdentityHandler,
		Authenticator:   jwtAuth.Authenticator,
		Authorizator:    jwtAuth.Authorizator,
		Unauthorized:    jwtAuth.Unauthorized,
		LoginResponse:   jwtAuth.LoginResponse,
		RefreshResponse: jwtAuth.RefreshResponse,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	}
	err := middleware.MiddlewareInit()
	return &middleware, err
}
