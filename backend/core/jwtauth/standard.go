package jwtauth

import (
	"backend/core/sdk/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Standard struct {
}

// 根据登录信息执行用户身份验证的回调函数。
// 必须返回用户数据作为用户标识符，它将存储在索赔数组中。
// 必需的。检查错误(e)以确定适当的错误消息。
func (h *Standard) Authenticator(c *gin.Context) (interface{}, error) {
	login := LoginModel{}
	if err := c.ShouldBind(&login); err != nil {
		return nil, ErrMissingLoginValues
	}
	db, err := pkg.GetOrm(c)
	if err != nil {
		return nil, ErrFailedAuthentication
	}
	claims, err := login.Verify(db)
	if err != nil {
		return nil, ErrFailedAuthentication
	}
	return map[string]interface{}{"user": claims}, nil
}

// 在登录时调用的回调函数。使用这个函数可以向webtoken添加额外的有效负载数据。
// 然后，通过c.Get("JWT_PAYLOAD")在请求期间可用数据。注意，有效负载没有加密。
// 在jwt中提到的属性。IO不能用作映射的键。可选，默认情况下不设置额外数据。
// 验证通过 data = user
func (h *Standard) PayloadFunc(data interface{}) MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, ok := v["user"].(MapClaims)
		if ok {
			return u
		}
	}
	return MapClaims{}
}

// 验证通过 身份
func (h *Standard) IdentityHandler(c *gin.Context) interface{} {
	claims := ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims[IdentityKey],
		"UserName":    claims[NiceKey],
		"UserId":      claims[IdentityKey],
	}
}

// 它应该执行对已验证用户的授权。
// Token验证
// 认证成功后才调用。成功时必须返回true，失败时返回false。
// 可选，默认为success
func (h *Standard) Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(map[string]interface{}); ok {
		c.Set("userId", v["UserId"])
		c.Set("userName", v["UserName"])
		return true
	}
	return false
}

func (h *Standard) Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}
