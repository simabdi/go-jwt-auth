package exception

type HTTPError struct {
	Code int
	Err  error
}

func NewHTTPError(code int, err error) error {
	return &HTTPError{
		Code: code,
		Err:  err,
	}
}

func (e *HTTPError) Error() string {
	return e.Err.Error()
}
