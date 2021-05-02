package launch

type LaunchRst string

func (l LaunchRst) IsOK() bool {
	return true
}

func (l LaunchRst) Err() error {
	return nil
}
