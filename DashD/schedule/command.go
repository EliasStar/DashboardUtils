package schedule

import (
	"context"
	"os"
	"strings"

	"github.com/EliasStar/Dashboard/DashD/command"
	"github.com/EliasStar/Dashboard/DashD/launch"
)

const CrontabLocation = "/etc/cron.d/dashd"

// * * * * * root /bin/sh /home/root/script.sh

type Command struct {
	Action         Action
	CronExpression string
	Command        launch.Command
}

func (c Command) IsValid(ctx context.Context) bool {
	x := c.Action.IsValid()
	y := c.Command.IsValid(ctx)
	z := len(strings.Fields(c.CronExpression)) == 5

	if c.Action != ActionWrite {
		z = z || c.CronExpression == ""
	}

	return x && y && z
}

func (c Command) Execute(ctx context.Context) command.Result {
	cronLines := []string{}

	if _, err := os.Stat(CrontabLocation); err == nil {
		content, err := os.ReadFile(CrontabLocation)
		if err != nil {
			return command.ErrorRst(err.Error())
		}

		cronLines = strings.Split(string(content), "\n")
	}

	cmd := c.Command.Executable + " " + strings.Join(c.Command.Arguments, " ")

	switch c.Action {
	case ActionRead:
		rst := Result{}

		for _, v := range cronLines {
			if strings.Contains(v, c.CronExpression) && strings.Contains(v, cmd) {
				rst = append(rst, v)
			}
		}

		return rst

	case ActionWrite:
		for _, v := range cronLines {
			if strings.Contains(v, c.CronExpression) && strings.Contains(v, cmd) {
				return command.OKRst{}
			}
		}

		cronLines = append(cronLines, c.CronExpression+" root "+cmd)

	case ActionRemove:
		for i, v := range cronLines {
			if strings.Contains(v, c.CronExpression) && strings.Contains(v, cmd) {
				cronLines = append(cronLines[:i], cronLines[i+1:]...)
			}
		}
	}

	err := os.WriteFile(CrontabLocation, []byte(strings.Join(cronLines, "\n")), os.ModePerm)
	return command.ResultFromError(err)
}
