package command

type Result interface {
	IsOK() bool
}

type ErrorRst struct {
	Error error
}

func (e ErrorRst) IsOK() bool {
	return e.Error == nil
}
