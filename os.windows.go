// +build windows

package main

import (
	"os"
	"os/exec"
	"syscall"
)

var handledSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGHUP,
	syscall.SIGTERM,
	// syscall.SIGTTIN,
	// syscall.SIGTTOU,
	// syscall.SIGUSR1,
	// syscall.SIGUSR2,
}

func detach(cmd *exec.Cmd) {
	// noop on windows
}
