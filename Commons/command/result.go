package command

type Result interface {
	OK() bool
	Err() error
}
