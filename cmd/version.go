package cmd

import (
	"github.com/gearboxworks/buildtool/defaults"
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/toolSelfUpdate"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(selfUpdateCmd)
}
var versionCmd = &cobra.Command{
	Use:   loadTools.CmdVersion,
	Short: ux.SprintfMagenta(defaults.BinaryName) + ux.SprintfBlue(" - Show version of executable."),
	Long:  ux.SprintfMagenta(defaults.BinaryName) + ux.SprintfBlue(" - Show version of executable."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			return
		}

		Cmd.State = Version()
	},
}
var selfUpdateCmd = &cobra.Command{
	Use:   loadTools.CmdSelfUpdate,
	Short: ux.SprintfMagenta(defaults.BinaryName) + ux.SprintfBlue(" - Update version of executable."),
	Long: ux.SprintfMagenta(defaults.BinaryName) + ux.SprintfBlue(" - Check and update the latest version."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			return
		}

		Cmd.State = VersionUpdate()
	},
}


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
