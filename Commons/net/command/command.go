package command

import "github.com/EliasStar/DashboardUtils/Commons/net/result"

type Command interface {
	Execute()
	IsDone() bool
	GetResult() result.Result
	GetError() error
}

type GeneralCmd struct {
	isDone bool
	result result.Result
	err    error
}

func (c GeneralCmd) Execute() {
	c.isDone = true
}

func (c GeneralCmd) IsDone() bool {
	return c.isDone
}

func (c GeneralCmd) GetResult() result.Result {
	return c.result
}

func (c GeneralCmd) GetError() error {
	return c.err
}
