package display

import (
	"context"
	"os/exec"
	"syscall"

	"github.com/EliasStar/Dashboard/DashD/command"
)

const Browser = "browser"

type ContextKey struct{}

type Command struct {
	Action Action
	URL    string
}

func (c Command) IsValid(ctx context.Context) bool {
	return c.Action.IsValid()
}

func (c Command) Execute(ctx context.Context) command.Result {
	cmd, ok := ctx.Value(ContextKey{}).(*exec.Cmd)
	if !ok {
		return command.ErrorRst("display not initialized")
	}

	if c.Action == ActionGet {
		var url string
		if len(cmd.Args) > 1 {
			url = cmd.Args[1]
		}

		return Result(url)
	}

	if cmd.Process != nil {
		cmd.Process.Signal(syscall.SIGTERM)
		cmd.Wait()

		*cmd = exec.Cmd{}
	}

	if c.Action == ActionSet {
		*cmd = *exec.Command(Browser, c.URL)
		cmd.Start()
	}

	return command.OKRst{}
}
