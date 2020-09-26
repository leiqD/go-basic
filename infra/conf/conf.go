package conf

import (
	"fmt"
	"github.com/leiqD/go-socket5/pkg/fileutil"

	"os"
	"sync"
)

type Configs struct {
	Viper *Viper
}

/*
type ConfigsInterface interface {
	ReloadViper()
	Log() *LogInfo
	DataBase() *DataStoreInfo
}
*/
var C *Configs
var cOnce sync.Once
var AppPath string

func init() {
	var err error
	AppPath, err = fileutil.GetBinParentPath()
	if err != nil {
		fmt.Printf("Query Bindir error:%s", err.Error())
		os.Exit(-1)
	}
	AppPath = AppPath + "/"
}

func NewConfig(vc *Viper) *Configs {
	cOnce.Do(func() {
		C = &Configs{Viper: vc}
	})
	return C
}

func (c *Configs) ReloadViper() {
	c.Viper = ReloadViper()
}

func (c *Configs) Log() *LogInfo {
	log := c.Viper.Log
	return &log
}

func (c *Configs) DataBase() *DataStoreInfo {
	sql := c.Viper.DataStore
	return &sql
}
