package main

import (
	"fmt"
	"os/exec"
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
		return fmt.Errorf("sysCmdPkg not supported on this OS : %s", config.sysCmdPkg)
	default:
		return fmt.Errorf("invalid sysCmdPkg : %s", config.sysCmdPkg)
	}
}
