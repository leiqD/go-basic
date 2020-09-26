package app

import (
	"github.com/leiqD/go-socket5/app/socket5/launcher"
	"github.com/leiqD/go-socket5/infra/conf"
	"github.com/leiqD/go-socket5/infra/logger"
	"gorm.io/gorm"
)

type Program struct {
	conf *conf.Configs
	log  logger.LoggerInterface
	db   *gorm.DB
}

func (p *Program) Init() error {
	p.conf = launcher.InitializeConfig("")
	p.log = launcher.InitializeLog(p.conf)
	db, err := launcher.InitialDataStore(p.conf)
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

func (p *Program) Start() error {
	logger.Info("Service Start")
	return nil
}

func (p *Program) Stop() error {
	return nil
}

func (p *Program) ReloadConfig() error {
	p.conf.ReloadViper()
	return nil
}

func (p *Program) OneLoop() error {
	return nil
}
