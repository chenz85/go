package timer

import (
	"time"
)

type TimerFunc = func()

// 在指定时间之后执行函数
// duration 时间，毫秒
func SetTimeout(callback TimerFunc, duration int64) {
	timeout := time.Duration(duration) * time.Millisecond
	time.AfterFunc(timeout, callback)
}

type Interval chan bool

func (interval Interval) Close() {
	interval <- true
	close(interval)
}

// 按间隔时间执行函数
// interval 间隔时间，毫秒
// async 是否异步执行
func SetInterval(callback TimerFunc, interval int64, async bool) (clear Interval) {
	t := time.Duration(interval) * time.Millisecond
	ticker := time.NewTicker(t)
	clear = make(Interval)

	go interval_func(ticker, clear, async, callback)

	return
}

func interval_func(ticker *time.Ticker, clear Interval, async bool, callback TimerFunc) {
	for {
		select {
		case <-ticker.C:
			if async {
				go callback()
			} else {
				callback()
			}
		case <-clear:
			// 停止计时并返回
			ticker.Stop()
			return
		}
	}
}

// 当前时间，秒
func Now() int64 {
	return time.Now().Unix()
}

// 当前时间，毫秒
func NowMS() int64 {
	return time.Now().UnixNano() / 1000000
}
