package errr

type Err struct {
	code    string
	message string
}

func NewErr(code string, message string) *Err {
	return &Err{
		code:    code,
		message: message,
	}
}

func (e *Err) GetCode() string {
	return e.code
}

func (e *Err) Error() string {
	return e.message
}
