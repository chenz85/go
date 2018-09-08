package erro

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type _Error struct {
	code int32
	msg  string
	data interface{}
}

var (
	code_checker = make(map[int32]bool)
)

func (ei *_Error) Error() string {
	return fmt.Sprintf("[_Error#%d:%s]", ei.code, ei.msg)
}

func (ei *_Error) Code() int32 {
	return ei.code
}

func (ei *_Error) Msg() string {
	return ei.msg
}
func (ei *_Error) SetData(_data interface{}) {
	ei.data = _data
}
func (ei *_Error) Data() interface{} {
	return ei.data
}

func New(code int32, msg string) Error {
	if _, ex := code_checker[code]; ex {
		log.Fatalln("[error] new a duplicated error:", code, msg)
	} else {
		code_checker[code] = true
	}
	return &_Error{
		code: code,
		msg:  msg,
	}
}

func (ei *_Error) F(format string, args ...interface{}) Error {
	new_ei := &_Error{
		code: ei.code,
		msg:  ei.msg + "::" + fmt.Sprintf(format, args...),
	}
	return new_ei
}

func (ei *_Error) D(data interface{}) Error {
	ei.data = data
	return ei
}

func (ei *_Error) With(err error) Error {
	return ei.F("err: %v", err)
}

func (ei *_Error) Is(err Error) bool {
	if _err, ok := err.(*_Error); !ok {
		return false
	} else {
		return ei.code == _err.code
	}
}
