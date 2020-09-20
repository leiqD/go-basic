package fileutil

import (
	"os"
	"path/filepath"
)

// OpenOrCreateReadNamedPipe 如果存在则打开有名管道，否则创建并打开
func OpenOrCreateReadNamedPipe(pipename string) (pipe *os.File, err error) {
	pipe, err = os.OpenFile(pipename, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		if err := os.MkdirAll(filepath.Dir(pipename), 0755); err != nil {
			return pipe, err
		}

		//syscall.Mkfifo(pipename, 0666)
		pipe, err = os.OpenFile(pipename, os.O_RDONLY|os.O_CREATE, os.ModeNamedPipe)
		if err != nil {
			return nil, err
		}
	}

	return pipe, err
}

// OpenOrCreateWriteNamedPipe 如果存在则打开有名管道，否则创建并打开
func OpenOrCreateWriteNamedPipe(pipename string) (pipe *os.File, err error) {
	pipe, err = os.OpenFile(pipename, os.O_WRONLY, os.ModeNamedPipe)
	if err != nil {
		if err := os.MkdirAll(filepath.Dir(pipename), 0755); err != nil {
			return pipe, err
		}

		//syscall.Mkfifo(pipename, 0666)
		pipe, err = os.OpenFile(pipename, os.O_WRONLY|os.O_CREATE, os.ModeNamedPipe)
		if err != nil {
			return nil, err
		}
	}

	return pipe, err
}
