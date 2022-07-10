package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v",
		EnvConfigs.PostgresUser,
		EnvConfigs.PostgresPassword,
		EnvConfigs.PostgresHost,
		EnvConfigs.PostgresPort,
		EnvConfigs.PostgresDB,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	EnvConfigs.DB = db
}
