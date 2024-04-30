package self

import (
	"fmt"
	"os/exec"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/resources/runnable/commons"
	"syscall"
)

func performOpOnSelf(config *config.Config, op commons.RunnableServiceOperationType) error {
	sysCmdPkg := config.RunnableServiceSelfSysCmdPkg

	switch sysCmdPkg {
	case "exec":
		var cmd *exec.Cmd
		switch op {
		case commons.REBOOT:
			cmd = exec.Command("reboot")
		case commons.STOP:
			cmd = exec.Command("shutdown")
		}
		return cmd.Run()
	case "syscall":
		var cmd int
		switch op {
		case commons.REBOOT:
			cmd = syscall.LINUX_REBOOT_CMD_RESTART
		case commons.STOP:
			cmd = syscall.LINUX_REBOOT_CMD_POWER_OFF
		}
		return syscall.Reboot(cmd)
	default:
		return fmt.Errorf("invalid sysCmdPkg : %s", sysCmdPkg)
	}
}
