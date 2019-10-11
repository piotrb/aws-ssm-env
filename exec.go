package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/kballard/go-shellquote"
)

var childCmd *exec.Cmd

func execWithParams(params []*ssm.Parameter, shell string, upcase bool) {
	bashCmd := shellquote.Join(flag.Args()...)
	words := []string{shell, "-c", fmt.Sprintf("exec %s", bashCmd)}

	command := words[0]
	remainingParts := words[1:len(words)]
	childCmd = exec.Command(command, remainingParts...)

	fmt.Printf("[aws-ssm-env] Running command: %v ...\n", words)

	childCmd.Env = append(os.Environ(), paramsToEnv(params, upcase)...)

	childCmd.Stdin = os.Stdin
	childCmd.Stderr = os.Stderr
	childCmd.Stdout = os.Stdout

	handleSingnals("aws-ssm-env", []os.Signal{syscall.SIGINT, syscall.SIGTERM}, func(signal os.Signal) {
		// nothing
	})

	err := childCmd.Run()
	handleCmdExit(childCmd, err, "[aws-ssm-env] ")
}

func paramsToEnv(params []*ssm.Parameter, upcase bool) []string {
	var result []string
	for _, param := range params {
		split := strings.Split(*param.Name, "/")
		name := split[len(split)-1]
		if upcase {
			name = strings.ToUpper(name)
		}
		value := fmt.Sprintf("%s=%s", name, *param.Value)
		result = append(result, value)
	}
	return result
}
