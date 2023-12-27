package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	pipeOp = "|"
)

type Shell struct{}

func NewShell() *Shell {
	return &Shell{}
}

func (sh *Shell) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	sh.invite()

	for scanner.Scan() {
		cmdStr := scanner.Text()
		err := sh.exec(cmdStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		sh.invite()
	}
}

func (sh *Shell) exec(cmd string) error {
	if strings.Contains(cmd, pipeOp) {
		cmds := strings.Split(cmd, pipeOp)
		for i, c := range cmds {
			cmds[i] = strings.Trim(c, " ")
		}

		out, err := sh.pipe(cmds)
		if err != nil {
			return err
		}

		fmt.Fprintln(os.Stdout, string(out))

	} else {
		cmd, err := sh.findCmd(cmd)
		if err != nil {
			return err
		}

		if err = cmd.Execute(); err != nil {
			return err
		}

		fmt.Fprintln(os.Stdout, string(cmd.Output()))
	}
	return nil
}

func (sh *Shell) invite() {
	path, err := filepath.Abs(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s> ", path)
}

func (sh *Shell) pipe(cmds []string) ([]byte, error) {
	buf := new(bytes.Buffer)

	for _, cmd := range cmds {
		cmd, err := sh.findCmd(cmd)
		if err != nil {
			return nil, err
		}

		nextInput, err := io.ReadAll(buf)
		if err != nil {
			return nil, err
		}

		cmd.SetInput(nextInput)

		if err = cmd.Execute(); err != nil {
			return nil, err
		}

		buf.Reset()
		if _, err := buf.Write(cmd.Output()); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (sh *Shell) findCmd(cmd string) (Command, error) {
	cmdArgs := strings.Split(cmd, " ")
	if len(cmdArgs) == 0 {
		return nil, fmt.Errorf("len(cmdArgs) == 0")
	}

	cmdName := strings.ToLower(cmdArgs[0])
	args := cmdArgs[1:]

	switch cmdName {
	case "cd":
		return &CdCommand{args, NewInOutCommand()}, nil
	case "pwd":
		return &PwdCommand{args, NewInOutCommand()}, nil
	case "echo":
		return &EchoCommand{args, NewInOutCommand()}, nil
	case "kill":
		return &KillCommand{args, NewInOutCommand()}, nil
	case "ps":
		return &PsCommand{args, NewInOutCommand()}, nil
	case "exec":
		return &ExecCommand{args, NewInOutCommand()}, nil
	case "fork":
		return &ForkCommand{args, NewInOutCommand()}, nil
	case "nc":
		return &NetcatCommand{args, NewInOutCommand()}, nil
	case "q":
		return &QCommand{args, NewInOutCommand()}, nil
	default:
		return nil, fmt.Errorf("command '%s' not found", cmdName)
	}
}
