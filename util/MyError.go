package util

type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return e.Msg
}

func NewError(code int, msg string) error {
	return &MyError{
		Code: code,
		Msg:  msg,
	}

}
