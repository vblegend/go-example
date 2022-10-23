package apis

import (
	"backend/core/sdk/api"
	"backend/core/sdk/config"
	"backend/core/sdk/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SysDBConfig struct {
	api.Api
}

func (e SysDBConfig) GetMYSQLConfigure(c *gin.Context) {
	e.MakeContext(c)
	cfg := config.Mysql{Host: config.MysqlConfig.Host, Port: config.MysqlConfig.Port, User: config.MysqlConfig.User, Password: config.MysqlConfig.Password}
	result, err := pkg.AESEncryptString(config.MysqlConfig.Password, pkg.DefaultAesKey)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	cfg.Password = result
	e.OK(cfg, "查询成功")
}

func (e SysDBConfig) SaveMYSQLConfigure(c *gin.Context) {
	req := config.Mysql{}
	e.MakeContext(c).MakeOrm()
	err := c.ShouldBind(&req)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	result, err := pkg.AESDecryptString(req.Password, pkg.DefaultAesKey)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	// 对比MySQL数据库配置是否发生修改
	err = compareMySQLConfig(c, &req)
	if err != nil {
		e.Logger.Error(err)
	}

	config.MysqlConfig.Host = req.Host
	config.MysqlConfig.Port = req.Port
	config.MysqlConfig.User = req.User
	config.MysqlConfig.Password = result
	config.Save()
	e.OK(gin.H{}, "保存成功")
}

func (e SysDBConfig) GetRedisConfigure(c *gin.Context) {
	e.MakeContext(c)
	cfg := config.Redis{Host: config.RedisConfig.Host, Port: config.RedisConfig.Port, Password: config.RedisConfig.Password, DB: config.RedisConfig.DB}
	result, err := pkg.AESEncryptString(config.RedisConfig.Password, pkg.DefaultAesKey)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	cfg.Password = result
	e.OK(cfg, "查询成功")
}

func (e SysDBConfig) SaveRedisConfigure(c *gin.Context) {
	req := config.Redis{}
	e.MakeContext(c).MakeOrm()
	err := c.ShouldBind(&req)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	result, err := pkg.AESDecryptString(req.Password, pkg.DefaultAesKey)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	// 对比Redis数据库配置是否发生修改
	err = compareRedisConfig(c, &req)
	if err != nil {
		e.Logger.Error(err)
	}

	config.RedisConfig.Host = req.Host
	config.RedisConfig.Port = req.Port
	config.RedisConfig.DB = req.DB
	config.RedisConfig.Password = result
	config.Save()
	e.OK(gin.H{}, "保存成功")
}

// compareMySQLConfig 通过MySQL新旧数据库配置对比,在系统操作日志中增加数据库修改配置
func compareMySQLConfig(c *gin.Context, newMysqlConfig *config.Mysql) error {
	if newMysqlConfig == nil {
		return nil
	}
	// 读取配置获取数据库连接对象
	oldMysqlConfig := config.Mysql{Host: config.MysqlConfig.Host, Port: config.MysqlConfig.Port, User: config.MysqlConfig.User, Password: config.MysqlConfig.Password}
	password, err := pkg.AESEncryptString(config.MysqlConfig.Password, pkg.DefaultAesKey)
	if err != nil {
		return err
	}
	oldMysqlConfig.Password = password
	return err
}

// compareMySQLConfig 通过Redis新旧数据库配置对比,在系统操作日志中增加数据库修改配置
func compareRedisConfig(c *gin.Context, newRedisConfig *config.Redis) error {
	if newRedisConfig == nil {
		return nil
	}
	// 读取配置获取数据库连接对象
	oldRedisConfig := config.Redis{Host: config.RedisConfig.Host, Port: config.RedisConfig.Port, Password: config.RedisConfig.Password, DB: config.RedisConfig.DB}
	password, err := pkg.AESEncryptString(config.RedisConfig.Password, pkg.DefaultAesKey)
	if err != nil {
		return err
	}
	oldRedisConfig.Password = password
	return err
}
