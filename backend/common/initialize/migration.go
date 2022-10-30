package initialize

import "backend/migration"

func InitMigration() {
	migration.DataBaseMigrate()
}
