// +build windows

package command

import (
	"fmt"
	"github.com/ghaoo/rboot"
	"os/exec"
	"strings"
)

func setup(bot *rboot.Robot, in *rboot.Message) []*rboot.Message {
	rule := in.Header.Get("rule")

	cmd := command[rule]

	for _, c := range cmd.Cmd {
		args := strings.SplitN(c, " ", 2)

		runCmd := exec.Command(args[0], args[1])

		output, err := runCmd.CombinedOutput()
		if err != nil {
			bot.Outgoing(rboot.NewMessage(fmt.Sprintf("error running command: %v: %q", err, string(output))))
		}

		bot.Outgoing(rboot.NewMessage(string(output)))
	}

	return nil
}