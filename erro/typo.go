package erro

type Error interface {
	error
	Code() int32
	Msg() string

	D(data interface{}) Error
	F(format string, args ...interface{}) Error
	With(err error) Error

	Data() interface{}
	Is(err Error) bool
}
