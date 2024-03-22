package main

import (
	"errors"
	"fmt"
	"os/exec"
	"syscall"
)

func performOpOnSelf(config *Config, op ServiceOperationType) error {
	switch config.sysCmdPkg {
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
		return errors.New(fmt.Sprintf("Invalid sysCmdPkg : %s", config.sysCmdPkg))
	}
}
