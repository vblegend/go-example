package initialize

import "server/migration"

func InitMigration() {
	migration.DataBaseMigrate()
}
