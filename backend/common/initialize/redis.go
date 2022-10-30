package initialize

import (
	"backend/common/config"
	"backend/core/echo"
	"backend/core/log"
	"backend/core/sdk"
	"fmt"

	"github.com/go-redis/redis/v7"
)

// 初始化redis连接
func InitRedisDB() {
	client := sdk.Runtime.GetRedisClient()
	if client != nil {
		client.Close()
	}
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Error(echo.Red("Redis connect fail..."))
		return
	}
	sdk.Runtime.SetRedisClient(client)
	log.Info(echo.Green("Redis connect sucess..."))
}
