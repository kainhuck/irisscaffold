package errno

type ErrorType int

func (e *ErrorType) NewError(msg string) Error {
	*e++
	return NewError(int(*e)-1, msg)
}
