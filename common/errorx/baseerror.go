package errorx

const defaultCode = 10000

type CodeError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code uint32, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}
func (e *CodeError) StatusCode() uint32 {
	return e.Code
}
func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
