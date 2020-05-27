package cmd

import (
	"bytes"
	"github.com/newclarity/buildtool/ux"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)


type TypeExecCommand struct {
	exe    string
	args   []string

	show   bool
	stdout []byte
	stderr []byte
	exit   int

	debug  bool

	State  *ux.State
}


func NewExecCommand(debugMode bool) *TypeExecCommand {
	ret := &TypeExecCommand {
		exe:    "",
		args:   nil,

		show:   false,
		stdout: []byte{},
		stderr: []byte{},
		exit:   0,

		debug:  debugMode,

		State:  ux.NewState(debugMode),
	}
	ret.State.SetPackage("")
	ret.State.SetFunctionCaller()

	return ret
}


func ExecCommand(cmd string, args ...string) *TypeExecCommand {
	e := NewExecCommand(Cmd.Debug)

	for range OnlyOnce {
		if e.State.IsNotOk() {
			e.State.PrintResponse()
			break
		}

		Cmd.State = e.SetPath(cmd)
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		if e.IsNotRunnable() {
			e.State.PrintResponse()
			break
		}

		Cmd.State = e.SetArgs(args...)
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}
		e.ShowProgress()

		if Cmd.Debug {
			ux.PrintflnBlue("# Executing: %s %s", e.exe, strings.Join(e.args, " "))
		}
		Cmd.State = e.Run()
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}
	}

	return e
}


func (e *TypeExecCommand) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


func (e *TypeExecCommand) Exec(cmd string, args ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		e.State = e.SetPath(cmd)
		if e.State.IsNotOk() {
			e.State.PrintResponse()
			break
		}

		if e.IsNotRunnable() {
			e.State.PrintResponse()
			break
		}

		e.State = e.SetArgs(args...)
		if e.State.IsNotOk() {
			e.State.PrintResponse()
			break
		}

		if e.debug {
			ux.PrintflnBlue("# Executing: %s %s", e.exe, strings.Join(e.args, " "))
		}
		e.State = e.Run()
		if e.State.IsNotOk() {
			e.State.PrintResponse()
			break
		}
	}

	return e.State
}


func (e *TypeExecCommand) Run() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}

	for range OnlyOnce {
		if e.State == nil {
			e.State = ux.NewState(e.debug)
		}

		c := exec.Command(e.exe, e.args...)

		var err error
		if e.show {
			var stdoutBuf, stderrBuf bytes.Buffer
			c.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
			c.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

			err := c.Run()
			if err != nil {
				e.State.SetError(err)
			}

			e.State.SetError(err)
			e.stdout = stdoutBuf.Bytes()
			e.stderr = stderrBuf.Bytes()

		} else {
			e.stdout, err = c.CombinedOutput()
			e.State.SetError(err)
		}

		if e.State.IsError() {
			if exitError, ok := err.(*exec.ExitError); ok {
				waitStatus := exitError.Sys().(syscall.WaitStatus)
				e.exit = waitStatus.ExitStatus()
				e.State.SetExitCode(e.exit)
			}
			break
		}

		waitStatus := c.ProcessState.Sys().(syscall.WaitStatus)
		e.exit = waitStatus.ExitStatus()
		e.State.SetExitCode(e.exit)
	}

	return e.State
}


func (e *TypeExecCommand) IsRunnable() bool {
	if state := e.IsNil(); state.IsError() {
		return false
	}
	var ok bool

	for range OnlyOnce {
		var err error

		e.exe, err = exec.LookPath(e.exe)
		if e.debug {
			ux.PrintflnBlue("Found executable at: %s", e.exe)
		}
		//fmt.Printf("exe: %s\n", e.exe)
		if err != nil {
			e.State.SetError("Executable not found.")
			break
		}

		ok = true
	}

	return ok
}
func (e *TypeExecCommand) IsNotRunnable() bool {
	return !e.IsRunnable()
}

func (e *TypeExecCommand) ShowProgress() {
	if state := e.IsNil(); state.IsError() {
		return
	}
	e.show = true
}
func (e *TypeExecCommand) SilenceProgress() {
	if state := e.IsNil(); state.IsError() {
		return
	}
	e.show = false
}


func (e *TypeExecCommand) GetExe() string {
	return e.exe
}
func (e *TypeExecCommand) GetPath() string {
	return e.exe
}
func (e *TypeExecCommand) SetPath(path ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}

	for range OnlyOnce {
		e.exe = filepath.Join(path...)
		if e.exe == "" {
			e.State.SetError("No path specified.")
			break
		}

		if e.IsNotRunnable() {
			break
		}

		e.State.SetOk()
	}

	return e.State
}


func (e *TypeExecCommand) GetArgs() []string {
	return e.args
}
func (e *TypeExecCommand) SetArgs(args ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}
	e.args = []string{}
	e.State = e.AddArgs(args...)
	return e.State
}
func (e *TypeExecCommand) AddArgs(args ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}

	e.args = append(e.args, args...)
	return e.State
}
func (e *TypeExecCommand) AppendArgs(args ...string) *ux.State {
	e.State = e.AddArgs(args...)
	return e.State
}


func (e *TypeExecCommand) GetStdout() []byte {
	return e.stdout
}
func (e *TypeExecCommand) GetStdoutString() string {
	return string(e.stdout)
}


func (e *TypeExecCommand) GetStderr() []byte {
	return e.stderr
}
func (e *TypeExecCommand) GetStderrString() string {
	return string(e.stderr)
}


func (e *TypeExecCommand) GetExitCode() int {
	return e.exit
}
