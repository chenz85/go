package event

import (
	"errors"
	"reflect"
)

var (
	_Err_Method_ParamsNumNotMatch = errors.New("method params num not match")
	_Err_Method_WrongParamType    = errors.New("wrong param type")
	_Err_Method_InvalidMethod     = errors.New("invalid method")
)

// rpc method interface
type EventMethod interface {
	Invoke() (result interface{}, err error)
	InvokeA(params []interface{}) (result interface{}, err error)
	InvokeWithParams(params ...interface{}) (result interface{}, err error)
}

// rpc method data
type _EventMethod struct {
	// func object
	rf reflect.Value
	// type of func object
	rft reflect.Type
	// param type list
	rfpt []reflect.Type
}

func (m *_EventMethod) Invoke() (result interface{}, err error) {
	if m.rft.NumIn() != 0 {
		err = _Err_Method_ParamsNumNotMatch
	} else {
		var result_vals = m.rf.Call(nil)
		result = m.return_values(result_vals)
	}
	return
}
func (m *_EventMethod) InvokeA(params []interface{}) (result interface{}, err error) {
	if m.rft.NumIn() != len(params) {
		err = _Err_Method_ParamsNumNotMatch
	} else {
		var param_vals = make([]reflect.Value, len(params))
		for i, p := range params {
			pt := reflect.TypeOf(p)
			if !check_arg_type(pt, m.rfpt[i]) {
				err = _Err_Method_WrongParamType
				return
			}
			param_vals[i] = reflect.ValueOf(p)
		}
		var result_vals = m.rf.Call(param_vals)
		result = m.return_values(result_vals)
	}
	return
}
func (m *_EventMethod) InvokeWithParams(params ...interface{}) (result interface{}, err error) {
	if m.rft.IsVariadic() {
		return m.InvokeVariadic(params...)
	} else {
		return m.InvokeA(params)
	}
}

func (m *_EventMethod) InvokeVariadic(params ...interface{}) (result interface{}, err error) {
	if param_num, arg_num := m.rft.NumIn(), len(params); param_num-1 > arg_num {
		err = _Err_Method_ParamsNumNotMatch
	} else {
		var param_vals = make([]reflect.Value, param_num)
		// normal argument
		for i := 0; i < param_num-1; i++ {
			var arg = params[i]
			var at = reflect.TypeOf(arg)
			if !check_arg_type(at, m.rfpt[i]) {
				err = _Err_Method_WrongParamType
				return
			}
			param_vals[i] = reflect.ValueOf(arg)
		}
		// variadic argument
		var variadic_vals = reflect.MakeSlice(m.rfpt[param_num-1], 0, arg_num-(param_num-1))
		var variadic_param_type = m.rfpt[param_num-1].Elem()
		for i := param_num - 1; i < arg_num; i++ {
			var arg = params[i]
			var at = reflect.TypeOf(arg)
			if !check_arg_type(at, variadic_param_type) {
				err = _Err_Method_WrongParamType
				return
			}
			variadic_vals = reflect.Append(variadic_vals, reflect.ValueOf(arg))
		}
		param_vals[param_num-1] = variadic_vals

		// call
		var result_vals = m.rf.CallSlice(param_vals)
		result = m.return_values(result_vals)
	}
	return
}

func (m *_EventMethod) return_values(vals []reflect.Value) (result interface{}) {
	if result_num := m.rf.Type().NumOut(); result_num == 0 {
		result = nil
	} else if result_num == 1 {
		result = vals[0].Interface()
	} else {
		var results = make([]interface{}, result_num)
		for i, rv := range vals {
			results[i] = rv.Interface()
		}
		result = results
	}
	return
}

// 参数类型判断
func check_arg_type(arg, param reflect.Type) bool {
	return arg.AssignableTo(param)
}

func NewEventMethod(method interface{}) (em EventMethod, err error) {
	rf := reflect.ValueOf(method)
	if !rf.IsValid() || rf.IsNil() || rf.Kind() != reflect.Func {
		err = _Err_Method_InvalidMethod
	} else {
		var _em = &_EventMethod{
			rf:  rf,
			rft: rf.Type(),
		}

		_em.rfpt = make([]reflect.Type, _em.rft.NumIn())
		for i := 0; i < _em.rft.NumIn(); i++ {
			_em.rfpt[i] = _em.rft.In(i)
		}
		em = _em
	}
	return
}
