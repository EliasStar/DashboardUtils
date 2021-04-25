package screen

type ScreenRst struct {
	Value bool
	Error error
}

func (s ScreenRst) OK() bool {
	return s.Error == nil
}

func (s ScreenRst) Err() error {
	return s.Error
}
