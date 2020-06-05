package cmd

import (
	"github.com/newclarity/scribeHelpers/toolSelfUpdate"
	"github.com/newclarity/scribeHelpers/ux"
)


func Version(args ...string) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		switch {
			case len(args) == 0:
				state = VersionShow()

			case args[0] == "info":
				state = VersionInfo()

			case args[0] == "check":
				state = VersionCheck()

			case args[0] == "update":
				state = VersionUpdate()

			default:
				state = VersionShow()
		}
	}

	return state
}


func VersionShow() *ux.State {
	state := Cmd.State

	//fmt.Printf("%s %s\n",
	//	ux.SprintfBlue(defaults.BinaryName),
	//	ux.SprintfCyan("v%s", defaults.BinaryVersion),
	//)

	return state
}


func VersionInfo() *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		update := toolSelfUpdate.New(Cmd.Runtime)
		if update.State.IsError() {
			state = update.State
			break
		}

		state = update.PrintVersion(nil)
		if state.IsNotOk() {
			state = update.State
			break
		}
	}

	return state
}


func VersionCheck() *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		update := toolSelfUpdate.New(Cmd.Runtime)
		if update.State.IsError() {
			state = update.State
			break
		}

		state = update.IsUpdated(false)
		if update.State.IsError() {
			break
		}
	}

	return state
}


func VersionUpdate() *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		update := toolSelfUpdate.New(Cmd.Runtime)
		if update.State.IsError() {
			state = update.State
			break
		}

		state = update.IsUpdated(true)
		if update.State.IsError() {
			break
		}

		state = update.Update()
		if state.IsNotOk() {
			break
		}
	}

	return state
}
