package fileutil

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// FileExist 判断文件是否存在
func FileExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

// WritePidFile 将PID写入PID文件
func WritePidFile(filename string) error {
	dir, _ := filepath.Split(filename)
	if !FileExist(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(filename, []byte(strconv.Itoa(os.Getpid())), 0666); err != nil {
		return err
	}

	return nil
}

// ListDir 按照文件后缀匹配目录文件
func ListDir(dirPath, suffix string) ([]string, error) {
	files := []string{}
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	suffix = strings.ToLower(suffix)
	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		if strings.HasSuffix(strings.ToLower(file.Name()), suffix) {
			files = append(files, path.Join(dirPath, file.Name()))
		}
	}

	return files, nil
}

//获取文件修改时间 返回unix时间戳
func GetFileModTime(path string) (int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return fi.ModTime().Unix(), err
}

// WriteFileWithName write data into file.
func WriteFileWithName(filename string, data []byte) error {
	dir, _ := filepath.Split(filename)
	if !FileExist(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// CreateDateDir 根据当前日期来创建文件夹 $path/$today/
func CreateDateDir(Path string) (string, error) {
	today := time.Now().Format("2006-01-02")
	dir := path.Join(Path, today)

	if fi, err := os.Stat(dir); err == nil || os.IsExist(err) {
		if fi.IsDir() {
			return dir, nil
		}
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		//logger.Errorf("CreateDateDir fail:%s", err.Error())
		return "", err
	}

	return dir, nil
}

func GetBinPath() (string, error) {
	bindir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return bindir, nil
}

func GetBinParentPath() (string, error) {
	bindir, err := GetBinPath()
	if err != nil {
		return "", err
	}

	pos := strings.LastIndex(bindir, string(filepath.Separator))
	if pos < 0 {
		return bindir, nil
	}
	return bindir[:pos], nil
}
