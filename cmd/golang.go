package cmd

import (
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/toolExec"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)


func Golang(args ...string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		if len(args) == 0 {
			goLangHelp()
			break
		}

		switch strings.ToLower(args[0]) {
			case "update":
				state = goLangUpdate(args[1:]...)

			default:
				goLangHelp()
		}
	}

	return state
}


func goLangHelp() *ux.State {
	state := Cmd.State

	ux.PrintflnYellow("Need to supply one of:")
	ux.PrintflnYellow("\tupdate - Detect go.mod files recursively and update modules.")

	return state
}


func goLangUpdate(path ...string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		if len(path) == 0 {
			path = []string{"."}
		}

		ux.PrintflnBlue("Updating go modules...")

		e := toolExec.NewMultiExec(Cmd.Runtime)
		if e.State.IsNotOk() {
			state = e.State
			break
		}

		state = e.Set("go", "get", "-u")
		if state.IsNotOk() {
			break
		}

		state = e.SetDontAppendFile()
		if state.IsNotOk() {
			break
		}

		state = e.SetChdir()
		if state.IsNotOk() {
			break
		}

		state = e.ShowProgress()
		if state.IsNotOk() {
			break
		}

		state = e.FindRegex("go.mod", path...)
		if state.IsNotOk() {
			break
		}

		p := e.GetPaths()
		ux.PrintflnBlue("Updating go modules in %d paths...", len(p))

		state = e.Run()
		if state.IsNotOk() {
			break
		}

		state.SetOk("go module update OK")
	}

	return state
}


func PkgReflect(paths ...string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		if len(paths) == 0 {
			paths = []string{"."}
		}

		pr := loadTools.PkgReflect {
			Notypes:    false,
			Nofuncs:    false,
			Novars:     false,
			Noconsts:   false,
			Unexported: false,
			Norecurs:   false,
			Stdout:     false,
			Gofile:     "",
			Notests:    false,
			Debug:      false,
			State:      nil,
		}
		state = loadTools.PackageReflect(pr, paths...)
		if state.IsNotOk() {
			break
		}
	}

	return state
}
