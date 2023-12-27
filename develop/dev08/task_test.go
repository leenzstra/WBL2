package main

import (
	"bytes"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShellCommands(t *testing.T) {
	tests := []struct {
		name string
		cmd  string
		err  bool
	}{
		{"cd ok", "cd .", false},
		{"cd err", "cd ./here1", true},
		{"random command", "asd asda sdagr", true},
		{"ps ok", "ps", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := NewShell()
			err := sh.exec(tt.cmd)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestShellPipeCommands(t *testing.T) {
	tests := []struct {
		name string
		cmd  string
		err  bool
	}{
		{"pipe ok", "exec go version | echo", false},
		{"pipe err", "asdasd | echo", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := NewShell()
			err := sh.exec(tt.cmd)

			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}


		})
	}
}

func TestCdCommand(t *testing.T) {
	tests := []struct {
		name string
		args []string
		err  bool
	}{
		{"cd ok", []string{"."}, false},
		{"cd err", []string{"./here1"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := CdCommand{tt.args, NewInOutCommand()}
			err := cmd.Execute()

			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPwdCommand(t *testing.T) {
	tests := []struct {
		name string
		args []string
		err  bool
	}{
		{"pwd ok", []string{}, false},
	}

	expected, _ := os.Getwd()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := PwdCommand{tt.args, NewInOutCommand()}
			err := cmd.Execute()

			dir := cmd.out.String()

			assert.Equal(t, expected, dir)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestEchoCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
		err      bool
	}{
		{"echo ok", []string{"input"}, "input", false},
		{"echo err", []string{}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := EchoCommand{tt.args, NewInOutCommand()}
			err := cmd.Execute()

			out := cmd.out.String()

			assert.Equal(t, tt.expected, out)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestKillCommand(t *testing.T) {
	// steam pid
	// pid := "10956"
	tests := []struct {
		name string
		args []string
		err  bool
	}{
		// {"kill ok", []string{pid}, false},
		{"kill err incorrect param", []string{"123321123"}, true},
		{"kill err", []string{"pid"}, true},
		{"kill err no args", []string{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := KillCommand{tt.args, NewInOutCommand()}
			err := cmd.Execute()

			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPsCommand(t *testing.T) {
	tests := []struct {
		name string
		args []string
		err  bool
	}{
		{"ps ok", []string{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := PsCommand{tt.args, NewInOutCommand()}
			err := cmd.Execute()

			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestExecCommand(t *testing.T) {
	tests := []struct {
		name string
		args []string
		err  bool
	}{
		{"exec ok", []string{"go", "version"}, false},
		{"exec err", []string{"notexists"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := ExecCommand{tt.args, NewInOutCommand()}
			err := cmd.Execute()

			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNetcatCommand(t *testing.T) {
	s, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	defer s.Close()

	fullAddr := s.Addr()
	addrAndPort := strings.Split(fullAddr.String(), ":")
	addr := addrAndPort[0]
	port := addrAndPort[1]

	input := bytes.NewBufferString("Input message")
	tests := []struct {
		name string
		args []string
		err  bool 
	}{
		{"nc ok", []string{"-p", port, "-P", fullAddr.Network(), "-a", addr}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]byte, input.Len())

			read := func() {
				conn, err := s.Accept()
				if tt.err {
					assert.Error(t, err)
					return
				}
				defer conn.Close()
				conn.Read(buf)
			}

			go read()

			cmd := NetcatCommand{tt.args, NewInOutCommand()}
			cmd.in = input
			err := cmd.Execute()

			if tt.err {
				require.Error(t, err)
				assert.Equal(t, make([]byte, input.Len()), buf)
			} else {
				require.NoError(t, err)
				assert.Equal(t, input.Bytes(), buf)
			}
		})
	}
}
