package erro

var (
	E_UnknownError = New(-1, "unknown error")

	E_InvalidIntValue = New(10001, "invalid int value")
)
