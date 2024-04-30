package resources_runnable

import (
	"fmt"
	"os/exec"
	"rebootx-on-prem/http-server-go/config"
	"syscall"
)

func performOpOnSelf(config *config.Config, op RunnableServiceOperationType) error {
	sysCmdPkg := config.RunnableServiceSelfSysCmdPkg

	switch sysCmdPkg {
	case "exec":
		var cmd *exec.Cmd
		switch op {
		case REBOOT:
			cmd = exec.Command("reboot")
		case STOP:
			cmd = exec.Command("shutdown")
		}
		return cmd.Run()
	case "syscall":
		var cmd int
		switch op {
		case REBOOT:
			cmd = syscall.LINUX_REBOOT_CMD_RESTART
		case STOP:
			cmd = syscall.LINUX_REBOOT_CMD_POWER_OFF
		}
		return syscall.Reboot(cmd)
	default:
		return fmt.Errorf("invalid sysCmdPkg : %s", sysCmdPkg)
	}
}
