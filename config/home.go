package config

import (
	"flag"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/czsilence/go/log"
)

var (
	_app_data_root string
)

func init() {
	flag.StringVar(&_app_data_root, "datadir", "", "specify the data path to store key & database.")
}

func AppDataHome() string {
	if _app_data_root == "" {
		if app_name, ex := GetString("app_name"); ex {
			_app_data_root = filepath.Join(homeDir(), "."+app_name)
		} else if app_name = filepath.Base(os.Args[0]); len(app_name) > 2 { // should be a valid execute file name.
			_app_data_root = filepath.Join(homeDir(), "."+app_name)
		} else {
			log.E("[app] can not get data dir, because of invalid app name.")
		}
	}

	// if err := os.MkdirAll(_app_data_root, 0755); err != nil {
	// 	log.E2("[app] create data root failed, err: %v", err)
	// }
	return _app_data_root
}

//////////////////////////////////////////////////////////
// home目录
func homeDir() (dir string) {
	user, err := user.Current()
	if nil == err && len(user.HomeDir) > 0 {
		return user.HomeDir
	}

	switch runtime.GOOS {
	case "windows":
		dir = _homeWindows()
	case
		"darwin",
		"linux":
		dir = _homeUnix()
	default:
		log.E("[app] unsupported os:", runtime.GOOS)
	}
	return
}

func _homeUnix() (dir string) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		dir = home
	} else {
		log.E("[app] can not get home path to store app data.")
	}
	return
}

func _homeWindows() string {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		log.E("[app] can not get home path to store app data.")
	}

	return home
}
