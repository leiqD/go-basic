//+build wireinject

package launcher

import (
	"github.com/google/wire"
	"github.com/leiqD/go-socket5/infra/conf"
	"github.com/leiqD/go-socket5/infra/logger"
)

func InitializeConfig(cfgPath string) *conf.Configs {
	wire.Build(conf.NewConfig, conf.NewViper)
	return nil
}

func InitializeLog(cfg *conf.Configs) *logger.Zap {
	wire.Build(logger.NewLogger)
	return nil
}
