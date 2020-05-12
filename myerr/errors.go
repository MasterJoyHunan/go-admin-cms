package myerr

//正常校验错误类
type NormalValidateError struct {
	s string
}

func NewNormalValidateError(str string) error {
	return &NormalValidateError{s: str}
}

func (e *NormalValidateError) Error() string {
	return e.s
}


// 数据库校验错误类
type DbValidateError struct {
	s string
}

func NewDbValidateError(str string) error {
	return &DbValidateError{s: str}
}

func (d *DbValidateError) Error() string {
	return d.s
}


