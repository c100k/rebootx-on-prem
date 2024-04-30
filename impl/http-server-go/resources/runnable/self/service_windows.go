package self

import (
	"fmt"
	"os/exec"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/resources/runnable/commons"
)

func performOpOnSelf(config *config.Config, op commons.RunnableServiceOperationType) error {
	sysCmdPkg := config.RunnableServiceSelfSysCmdPkg

	switch sysCmdPkg {
	case "exec":
		var cmd *exec.Cmd
		switch op {
		case commons.REBOOT:
			cmd = exec.Command("shutdown", "/s")
		case commons.STOP:
			cmd = exec.Command("shutdown", "/r")
		}
		return cmd.Run()
	case "syscall":
		return fmt.Errorf("sysCmdPkg not supported on this OS : %s", sysCmdPkg)
	default:
		return fmt.Errorf("invalid sysCmdPkg : %s", sysCmdPkg)
	}
}
