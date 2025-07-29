package error_wrapper

type ErrorCode string

func (e ErrorCode) String() string {
	return string(e)
}
