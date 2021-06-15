package model

// 重写Error，实现error接口
type ServerError struct {
	Code int
	Message string
}

func NewServerError(code int,message string)*ServerError{
	return &ServerError{
		Code:    code,
		Message: message,
	}
}

func (e *ServerError)Error()string{
	return e.Message
}
