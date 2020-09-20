package infra

import (
	. "github.com/spf13/viper"
	"sync"
)

type Configs struct {
	viper *Viper
}

var C *Configs
var cOnce sync.Once
var AppPath string

func NewConfig(vc *Viper) *Configs {
	cOnce.Do(func() {
		C = &Configs{viper: vc}
	})
	return C
}
