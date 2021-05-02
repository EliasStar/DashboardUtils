package screen

type ScreenRst bool

func (s ScreenRst) IsOK() bool {
	return true
}

func (s ScreenRst) Err() error {
	return nil
}
