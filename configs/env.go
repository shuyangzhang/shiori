package configs

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
	EnvConfigs.AllPrefixes = append(EnvConfigs.PrefixAliases, EnvConfigs.Prefix)
}

type envConfigs struct {
	Token string `mapstructure:"TOKEN"`

	Prefix        string   `mapstructure:"PREFIX"`
	PrefixAliases []string `mapstructure:"PREFIX_ALIASES"`
	AllPrefixes   []string
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
