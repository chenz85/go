package event

import (
	"sync"

	"github.com/czsilence/go/typo"
)

// 简单的事件处理

type EventHandlerFunc = func(args ...typo.Any)

type EventHandler struct {
	handleFunc EventHandlerFunc
}

type EventHandlerList = []*EventHandler

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
func On(eventName string, handlerFunc EventHandlerFunc) {
	handler := &EventHandler{handleFunc: handlerFunc}
	addEventHandler(eventName, handler, false)
}

// 绑定一次性执行的事件
func Once(eventName string, handlerFunc EventHandlerFunc) {
	handler := &EventHandler{handleFunc: handlerFunc}
	addEventHandler(eventName, handler, true)
}

func addEventHandler(name string, h *EventHandler, once bool) {
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
		pkg.once = append(pkg.once, h)
	} else {
		pkg.normal = append(pkg.normal, h)
	}
}

// 触发事件
func Emit(eventName string, args ...typo.Any) {
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
		v.handleFunc(args...)
	}
}
