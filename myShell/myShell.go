package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("shell> ")
		cmdLine, _ := reader.ReadString('\n')
		cmdLine = strings.TrimSuffix(cmdLine, "\n")
		if cmdLine == "" {
			continue
		}

		executeCommand(cmdLine)
	}
}

func executeCommand(cmdLine string) {
	if strings.Contains(cmdLine, "|") {
		executePipe(cmdLine)
		return
	}

	args := strings.Split(cmdLine, " ")
	switch args[0] {
	case "cd":
		changeDirectory(args)
	case "pwd":
		printWorkingDirectory()
	case "echo":
		echo(args)
	case "kill":
		killProcess(args)
	case "ps":
		printProcesses()
	case "quit":
		fmt.Println("Exiting shell...")
		os.Exit(0)
	default:
		executeCommandDirectly(args)
	}
}

func executeCommandDirectly(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("exec:", err)
	}
}

func changeDirectory(args []string) {
	if len(args) < 2 {
		fmt.Println("cd: expected argument")
	} else {
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func printWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dir)
	}
}

func echo(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func killProcess(args []string) {
	if len(args) < 2 {
		fmt.Println("kill: expected argument")
	} else {
		pid, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("kill: invalid pid")
		} else {
			process, err := os.FindProcess(pid)
			if err != nil {
				fmt.Println(err)
			} else {
				err := process.Kill()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func printProcesses() {
	cmd := exec.Command("tasklist")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func executePipe(cmdLine string) {
	commands := strings.Split(cmdLine, "|")
	var cmds []*exec.Cmd

	for _, cmdString := range commands {
		cmdArgs := strings.Fields(strings.TrimSpace(cmdString))
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		cmds = append(cmds, cmd)
	}

	var prevStdout io.ReadCloser
	for i, cmd := range cmds {
		if i != 0 {
			cmd.Stdin = prevStdout
		}
		if i != len(cmds)-1 {
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				fmt.Println(err)
				return
			}
			prevStdout = stdout
		} else {
			cmd.Stdout = os.Stdout
		}

		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			return
		}

		if i != len(cmds)-1 {
			if err := cmd.Wait(); err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	if err := cmds[len(cmds)-1].Wait(); err != nil {
		fmt.Println(err)
	}
}
