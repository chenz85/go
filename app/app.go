package app

import (
	"os"
	"path/filepath"

	"github.com/czsilence/go/config"
	"github.com/czsilence/go/log"
)

func GetAppDir() (dir string) {
	var err error
	if dir, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		log.E("[app] can not get app path.", err)
	}
	return
}

func GetDataDir() (dir string) {
	return config.AppDataHome()
}

// 获取数据目录
func GetDataPath(rel ...string) string {
	arr := make([]string, 0, len(rel)+1)
	arr = append(arr, GetDataDir())
	arr = append(arr, rel...)
	return filepath.Join(arr...)
}
