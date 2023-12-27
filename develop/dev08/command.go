package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

// Интерфейс команды
type Command interface {
	Execute() error
	Output() []byte
	SetInput([]byte)
}

type InOutCommand struct {
	out *bytes.Buffer
	in  *bytes.Buffer
}

func NewInOutCommand() InOutCommand {
	return InOutCommand{
		out: new(bytes.Buffer),
		in:  new(bytes.Buffer),
	}
}

func (c *InOutCommand) Output() []byte {
	return c.out.Bytes()
}

func (c *InOutCommand) SetInput(b []byte) {
	c.in.Write(b)
}

// команда "cd"
type CdCommand struct {
	Args []string
	InOutCommand
}

func (c *CdCommand) Execute() error {
	if len(c.Args) < 1 {
		return fmt.Errorf("len(c.Args) < 1")
	}
	return os.Chdir(c.Args[0])
}

// команда "pwd"
type PwdCommand struct {
	Args []string
	InOutCommand
}

func (c *PwdCommand) Execute() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Fprint(c.out, dir)
	return nil
}

// команда "echo"
type EchoCommand struct {
	Args []string
	InOutCommand
}

func (c *EchoCommand) Execute() error {
	if len(c.Args) > 0 {
		fmt.Fprint(c.out, strings.Join(c.Args, " "))
		return nil
	}

	if c.in.Len() != 0 {
		fmt.Fprint(c.out, c.in.String())
		return nil
	}
	return fmt.Errorf("no data")
}

// команда "kill"
type KillCommand struct {
	Args []string
	InOutCommand
}

func (c KillCommand) Execute() error {
	if len(c.Args) < 1 {
		return fmt.Errorf("len(c.Args) < 1")
	}

	pid, err := strconv.Atoi(c.Args[0])
	if err != nil {
		return err
	}

	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return p.Kill()
}

// команда "ps"
type PsCommand struct {
	Args []string
	InOutCommand
}

func (c PsCommand) Execute() error {
	procs, err := ps.Processes()
	if err != nil {
		return err
	}

	fmt.Fprintf(c.out, "%s\t%s\t%s\n", "N", "PID", "Name")
	for i, proc := range procs {
		fmt.Fprintf(c.out, "%d\t%d\t%s\n", i+1, proc.Pid(), proc.Executable())
	}

	return nil
}

// команда "q"
type QCommand struct {
	Args []string
	InOutCommand
}

func (c QCommand) Execute() error {
	os.Exit(0)
	return nil
}

// команда "exec"
type ExecCommand struct {
	Args []string
	InOutCommand
}

func (c ExecCommand) Execute() error {
	cmd := exec.Command(c.Args[0], c.Args[1:]...)
	cmd.Stdin = c.in
	cmd.Stdout = c.out

	return cmd.Run()
}

// команда "fork"
type ForkCommand struct {
	Args []string
	InOutCommand
}

func (c ForkCommand) Execute() error {
	return fmt.Errorf("not supported :(")
}

// команда "netcat"
type NetcatCommand struct {
	Args []string
	InOutCommand
}

func (c NetcatCommand) Execute() error {
	var (
		protocol, addr string
		port           int
		conn           net.Conn
	)

	fs := flag.NewFlagSet("netcat", flag.ContinueOnError)
	fs.IntVar(&port, "p", 5000, "port value")
	fs.StringVar(&addr, "a", "localhost", "addr value")
	fs.StringVar(&protocol, "P", "tcp", "protocol value")
	if err := fs.Parse(c.Args); err != nil {
		return err
	}

	conn, err := net.Dial(protocol, fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.Write(c.in.Bytes())
	if err != nil {
		return err
	}

	return nil
}
