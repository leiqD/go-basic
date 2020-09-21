package conf

import (
	"fmt"
	"github.com/leiqD/go-socket5/pkg/fileutil"

	"os"
	"sync"
)

type Configs struct {
	viper *Viper
}

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
		C = &Configs{viper: vc}
	})
	return C
}

func (c *Configs) ReloadViper() {
	c.viper = ReloadViper()
}
