package main

import (
	"fmt"
	"os/exec"
)

func performOpOnSelf(config *Config, op RunnableServiceOperationType) error {
	sysCmdPkg := config.runnableServiceSelfSysCmdPkg

	switch sysCmdPkg {
	case "exec":
		var cmd *exec.Cmd
		switch op {
		case REBOOT:
			cmd = exec.Command("shutdown", "/s")
		case STOP:
			cmd = exec.Command("shutdown", "/r")
		}
		return cmd.Run()
	case "syscall":
		return fmt.Errorf("sysCmdPkg not supported on this OS : %s", sysCmdPkg)
	default:
		return fmt.Errorf("invalid sysCmdPkg : %s", sysCmdPkg)
	}
}
