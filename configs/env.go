package configs

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var EnvConfigs *envConfigs

func initEnvConfigs() {
	EnvConfigs = loadEnvVariables()
	EnvConfigs.AllPrefixes = append(EnvConfigs.PrefixAliases, EnvConfigs.Prefix)
}

type envConfigs struct {
	Token string `mapstructure:"TOKEN"`

	Prefix        string   `mapstructure:"PREFIX"`
	PrefixAliases []string `mapstructure:"PREFIX_ALIASES"`
	AllPrefixes   []string

	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DB               *gorm.DB
}

func loadEnvVariables() (config *envConfigs) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
