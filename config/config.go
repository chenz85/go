package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/czsilence/go/log"
	"github.com/czsilence/go/typo"
)

const (
	// 默认的配置文件名
	// 可以通过参数指定配置文件
	default_config_file_name string = "config.json"
)

type ConfigEntry = typo.Any

var (
	// 配置信息表
	config map[string]ConfigEntry
	// 配置文件路径
	config_file string
)

// 默认配置
func defaultConfig() {
	if config != nil {
		return
	}

	config = map[string]ConfigEntry{
		// 默认的配置信息
	}
}

// 直接设置/修改配置
func Set(name string, val typo.Any) {
	defaultConfig()
	config[name] = val
}

// 解析配置文件
func Parse() {
	if !flag.Parsed() {
		flag.Parse()
	}

	defaultConfig()

	loadConfig()
}

func loadConfig() {
	//加载配置文件并应用配置
	if data, err := ioutil.ReadFile(config_file); err == nil {
		var config_obj typo.Map
		if err := json.Unmarshal(data, &config_obj); err == nil {
			for k, v := range config_obj {
				switch v := v.(type) {
				case json.Number:
					if v, err := v.Int64(); err != nil {
						log.W("[config] invalid number in config:", config_file, k)
					} else {
						config[k] = v
					}
				default:
					config[k] = v
				}
				log.D("[config] custom config:", fmt.Sprintf("%s => %v (type: %T)", k, v, v))
			}
			log.I("[config] load conifg from:", config_file)

		} else {
			log.W("[config] parse config file failed!", err)
		}
	} else {
		log.I("[config] no config file")
	}
}

func init() {
	flag.StringVar(&config_file, "c", default_config_file_name, "specify the config file path.")
	flag.StringVar(&config_file, "config", default_config_file_name, "specify the config file path.")
}

// 获取配置信息
func GetConfig(name string) (val interface{}, exist bool) {
	if config == nil {
		log.E("[config] can not get config before load!")
	}
	val, exist = config[name]
	return val, exist
}

// 获取配置信息，如果不存在则返回指定的默认值def
func GetConfigWithDefault(name string, def interface{}) interface{} {
	val, exist := GetConfig(name)
	if exist {
		return val
	} else {
		return def
	}
}

func GetString(name string) (val string, ex bool) {
	var _val interface{}
	var ok bool
	_val, ex = GetConfig(name)
	if !ex {
		return
	} else if val, ok = _val.(string); ok {
		return
	}
	return
}

func GetStringArray(name string) (val []string, ex bool) {
	var _val interface{}
	_val, ex = GetConfig(name)
	if !ex {
		return
	} else {
		switch _val := _val.(type) {
		case []string:
			val = _val
		default:
			log.W2("[config] invalid int value: %v, type: %T", _val, _val)
			ex = false
		}
	}
	return
}

func GetInt(name string) (val int64, ex bool) {
	var _val interface{}
	_val, ex = GetConfig(name)
	if !ex {
		return
	} else {
		switch _val := _val.(type) {
		case int64:
			val = _val
		case int:
			val = int64(_val)
		case float64:
			val = int64(_val)
		case string:
			var err error
			if val, err = strconv.ParseInt(_val, 10, 64); err != nil {
				ex = false
			}
		default:
			log.W2("[config] invalid int value: %v, type: %T", _val, _val)
			ex = false
		}
	}
	return
}
