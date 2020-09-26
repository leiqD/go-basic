//+build wireinject

package launcher

import (
	"github.com/google/wire"
	"github.com/leiqD/go-socket5/infra/conf"
	"github.com/leiqD/go-socket5/infra/datastore"
	"github.com/leiqD/go-socket5/infra/logger"
	"gorm.io/gorm"
)

func InitializeConfig(cfgPath string) *conf.Configs {
	wire.Build(conf.NewConfig, conf.NewViper)
	return nil
}

func InitializeLog(cfg logger.LoggerConfig) *logger.Zap {
	wire.Build(logger.NewLogger)
	return nil
}

func InitialDataStore(cfg datastore.MYSQLconfig) (*gorm.DB, error) {
	wire.Build(datastore.NewMYSQLDB)
	return nil, nil
}
