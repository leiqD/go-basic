//+build wireinject

package launcher

import (
	"github.com/google/wire"
	"github.com/leiqD/go-socket5/infra/conf"
)

func InitializeConfig(cfgPath string) *conf.Configs {
	wire.Build(conf.NewConfig, conf.NewViper)

	return nil
}
