package commands

import (
	"backend/core/logger"
	"backend/core/sdk/config"
	"backend/core/sdk/console"
	"backend/core/sdk/pkg"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	ResetConfig = &cobra.Command{
		Use:     "reset",
		Short:   "Reset the configuration file in the current directory.(The operation cannot be rolled back)",
		Example: "siteweb-manager config reset",
		Run: func(cmd *cobra.Command, args []string) {
			ymal := filepath.Join(pkg.AssemblyDir(), "/config/settings.yml")
			config.Setup(ymal)
			config.ApplicationConfig.Mode = pkg.Production
			config.ApplicationConfig.Host = "0.0.0.0"
			config.ApplicationConfig.Port = 8000
			config.ApplicationConfig.Name = "AppSidecar"
			// database
			config.DatabaseConfig.Driver = "sqlite3"
			config.DatabaseConfig.Source = "./sqlite.db"
			// mysql
			config.MysqlConfig.Host = "127.0.0.1"
			config.MysqlConfig.User = "root"
			config.MysqlConfig.Password = "Vertiv@086"
			config.MysqlConfig.Port = 3306
			// redis
			config.RedisConfig.DB = 0
			config.RedisConfig.Host = "127.0.0.1"
			config.RedisConfig.Password = "siteweb1!"
			config.RedisConfig.Port = 6379
			// logger
			config.LoggerConfig.Level = "error"
			config.LoggerConfig.Location = false
			config.LoggerConfig.Stdout = ""
			config.LoggerConfig.Path = "./temp/logs"
			config.LoggerConfig.Cap = 0

			config.Save()
			logger.Infof(console.Yellow(fmt.Sprintf("Configuration file “%s” has been reset...", ymal)))
			os.Exit(0)
		},
	}
)
