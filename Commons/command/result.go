package command

type Result interface {
	IsOK() bool
	Err() error
}

type OKRst struct{}

func (o OKRst) IsOK() bool {
	return true
}

func (o OKRst) Err() error {
	return nil
}

func NewErrorRst(err error) ErrorRst {
	return ErrorRst{
		Error: err,
	}
}

type ErrorRst struct {
	Error error
}

func (e ErrorRst) IsOK() bool {
	return e.Error == nil
}

func (e ErrorRst) Err() error {
	return e.Error
}
