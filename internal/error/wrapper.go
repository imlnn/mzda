package error

type ErrWrapper struct {
	ErrorCode    string
	ErrorMessage string
}

func (w *ErrWrapper) Error() string {
	return w.ErrorMessage
}
