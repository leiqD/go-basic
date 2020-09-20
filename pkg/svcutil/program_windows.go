package svcutil

import (
	"os"
	"os/signal"
	"syscall"
)

// Run 启动服务，并阻塞等待终端信号，接收到指定信号后退出，默认为（SIGINT, SIGTERM）
// windows环境下无SIGUSR2信号，不支持重载配置文件。
func Run(prg Program, sigs ...os.Signal) error {
	if err := prg.Init(); err != nil {
		return err
	}

	if len(sigs) == 0 {
		sigs = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, sigs...)

	if err := prg.Start(); err != nil {
		return err
	}

	for {
		select {
		case sig := <-sigChan:
			if sig == syscall.SIGINT || sig == syscall.SIGTERM {
				return prg.Stop()
			}
			//if sig == syscall.SIGUSR2 {
			//	prg.ReloadConfig()
			//}
		default:
		}

		prg.OneLoop()
	}

	return prg.Stop()
}
