package typo

import (
	"encoding/json"
	"errors"
)

// map字段检查
func HasFields(m Map, fields Fields) bool {
	for _, f := range fields {
		if _, ex := m[f]; !ex {
			return false
		}
	}
	return true
}

func GetField(m Map, f Field) (v Any, ex bool) {
	v, ex = m[f]
	return
}

func GetIntField(m Map, f Field) (vv int64, err error) {
	var v Any
	var ex bool
	if v, ex = m[f]; !ex {
		err = errors.New("not exist:" + f)
	} else if jn, ok := v.(json.Number); !ok {
		err = errors.New("not number:" + f)
	} else {
		vv, err = jn.Int64()
	}
	return
}

func GetStringField(m Map, f Field) (vv string, err error) {
	var v Any
	var ex, ok bool
	if v, ex = m[f]; !ex {
		err = errors.New("not exist:" + f)
	} else if vv, ok = v.(string); !ok {
		err = errors.New("not string:" + f)
	}
	return
}

func GetBoolField(m Map, f Field) (vv bool, err error) {
	var v Any
	var ex, ok bool
	if v, ex = m[f]; !ex {
		err = errors.New("not exist:" + f)
	} else if vv, ok = v.(bool); !ok {
		err = errors.New("not bool:" + f)
	}
	return
}

func GetArrayField(m Map, f Field) (vv Array, err error) {
	var v Any
	var ex, ok bool
	if v, ex = m[f]; !ex {
		err = errors.New("not exist:" + f)
	} else if vv, ok = v.(Array); !ok {
		err = errors.New("not array:" + f)
	}
	return
}

func GetMapField(m Map, f Field) (vv Map, err error) {
	var v Any
	var ex, ok bool
	if v, ex = m[f]; !ex {
		err = errors.New("not exist:" + f)
	} else if vv, ok = v.(Map); !ok {
		err = errors.New("not map:" + f)
	}
	return
}

// 复制map
func CopyMap(src map[string]Any, exclude []string) (dst map[string]Any) {
	dst = make(map[string]Any)
	for k, v := range src {
		if -1 != IndexOf(&exclude, k) {
			continue
		}
		dst[k] = v
	}
	return
}
