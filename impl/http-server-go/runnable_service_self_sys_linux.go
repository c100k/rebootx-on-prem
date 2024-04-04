package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func performOpOnSelf(config *Config, op RunnableServiceOperationType) error {
	sysCmdPkg := config.runnableServiceSelfSysCmdPkg

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
