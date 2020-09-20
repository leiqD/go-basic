package svcutil

// Service interface
type Program interface {
	Init() error
	ReloadConfig() error
	OneLoop() error
	Start() error
	Stop() error
}
