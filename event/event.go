package event

import (
	"sync"

	"github.com/czsilence/go/typo"
)

// 简单的事件处理

type EventHandlerList = []EventMethod

type EventPackage struct {
	normal EventHandlerList
	once   EventHandlerList
}

var (
	eventMap    map[string]*EventPackage
	mu_eventMap sync.RWMutex
)

func init() {
	checkInit()
}

func checkInit() {
	mu_eventMap.Lock()
	defer mu_eventMap.Unlock()

	if eventMap == nil {
		eventMap = make(map[string]*EventPackage)
	}
}

// 绑定事件
func On(eventName string, method interface{}) {
	if em, err := NewEventMethod(method); err == nil {
		addEventHandler(eventName, em, false)
	}
}

// 绑定一次性执行的事件
func Once(eventName string, method interface{}) {
	if em, err := NewEventMethod(method); err == nil {
		addEventHandler(eventName, em, true)
	}
}

func addEventHandler(name string, method EventMethod, once bool) {
	mu_eventMap.Lock()
	defer mu_eventMap.Unlock()

	pkg, ex := eventMap[name]
	if !ex {
		pkg = &EventPackage{
			normal: make(EventHandlerList, 0, 10),
			once:   make(EventHandlerList, 0, 10),
		}
	}
	eventMap[name] = pkg

	if once {
		pkg.once = append(pkg.once, method)
	} else {
		pkg.normal = append(pkg.normal, method)
	}
}

// 触发事件
// 如果事件响应过程中出错，返回第一个错误。
func Emit(eventName string, args ...typo.Any) (err error) {
	mu_eventMap.Lock()

	var tmp EventHandlerList
	if pkg, ok := eventMap[eventName]; ok {
		tmp = make(EventHandlerList, len(pkg.normal)+len(pkg.once))
		copy(tmp, pkg.normal)
		copy(tmp[len(pkg.normal):], pkg.once)
		// 清除once数组
		pkg.once = make(EventHandlerList, 0, 10)
	}
	mu_eventMap.Unlock()

	for _, v := range tmp {
		if _, ie := v.InvokeWithParams(args...); ie != nil && err == nil {
			err = ie
		}
	}
	return
}
