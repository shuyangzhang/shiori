package dal

import (
	"github.com/shuyangzhang/shiori/configs"
	"github.com/shuyangzhang/shiori/dal/query"
)

var Query = query.Use(configs.EnvConfigs.DB)
