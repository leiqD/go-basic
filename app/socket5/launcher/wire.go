//+build wireinject

package launcher

import (
	"github.com/google/wire"
	"github.com/leiqD/go-socket5/infra"
)

func InitializeConfig(cfgPath string) *infra.Configs {
	wire.Build(infra.NewConfig, infra.NewViper)

	return nil
}
