package common

const (
	EndOfIteratorErrorMsg string = `End of Iterator`
)

type (
	EndOfIteratorError struct {
		Err error
	}
)

func (e *EndOfIteratorError) Error() string {
	return EndOfIteratorErrorMsg
}
func (e *EndOfIteratorError) UnWrap() error {
	return e.Err
}
func (e *EndOfIteratorError) Wrap(err error) *EndOfIteratorError {
	return &EndOfIteratorError{Err: err}
}
