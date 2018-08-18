package erro

import (
	"fmt"

	"github.com/czsilence/go/typo"

	"github.com/czsilence/go/log"
)

type Error interface {
	error
	Code() int32
	Msg() string

	D(data typo.Any) Error
	F(format string, args ...typo.Any) Error
}

type _Error struct {
	code int32
	msg  string
	data interface{}
}

var (
	code_checker = make(map[int32]bool)
)

func (ei *_Error) Error() string {
	return fmt.Sprintf("[Error#%d:%s]", ei.code, ei.msg)
}

func (ei *_Error) Code() int32 {
	return ei.code
}

func (ei *_Error) Msg() string {
	return ei.msg
}
func (ei *_Error) SetData(_data typo.Any) {
	ei.data = _data
}
func (ei *_Error) Data() interface{} {
	return ei.data
}

func New(code int32, msg string) Error {
	if _, ex := code_checker[code]; ex {
		log.E("[errdef] new a duplicated error:", code, msg)
	} else {
		code_checker[code] = true
	}
	return &_Error{
		code: code,
		msg:  msg,
	}
}

func (ei *_Error) F(format string, args ...typo.Any) Error {
	new_ei := &_Error{
		code: ei.code,
		msg:  ei.msg + " " + fmt.Sprintf(format, args...),
	}
	return new_ei
}

func (ei *_Error) D(data interface{}) Error {
	ei.data = data
	return ei
}
