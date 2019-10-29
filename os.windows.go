// +build windows

package main

import (
	"os"
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