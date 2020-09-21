package app

import (
	"github.com/leiqD/go-socket5/app/socket5/launcher"
	"github.com/leiqD/go-socket5/infra/conf"
)

type Program struct {
	conf *conf.Configs
}

func (p *Program) Init() error {
	p.conf = launcher.InitializeConfig("")
	return nil
}

func (p *Program) Start() error {
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
