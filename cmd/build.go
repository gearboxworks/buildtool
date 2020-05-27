package cmd

import (
	"github.com/newclarity/buildtool/ux"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(buildCmd)
}


var buildCmd = &cobra.Command{
	Use:   CmdBuild,
	Short: ux.SprintfBlue("Build a gearboxworks repo."),
	Long: ux.SprintfBlue("Build a gearboxworks repo."),
	Run: Build,
}
func Build(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		tmpl := ProcessArgs(cmd, args)
		Cmd.State = tmpl.State
		if Cmd.State.IsNotOk() {
			break
		}
		//Cmd.State.DebugPrint()


		Cmd.State = GoReleaserBuild()
		if Cmd.State.IsNotOk() {
			break
		}
	}
}


const GoReleaserFile = ".goreleaser.yml"
func GoReleaserBuild() *ux.State {
	state := ux.NewState(Cmd.Debug)

	for range OnlyOnce {
		grFile := NewArgFile(Cmd.Debug)
		state = grFile.SetPath(GoReleaserFile)
		if grFile.NotExists() {
			state = grFile.State
			break
		}

		exe := NewExecCommand(Cmd.Debug)
		exe.ShowProgress()

		ux.PrintflnBlue("Found goreleaser file: %s", GoReleaserFile)
		state = exe.Exec("goreleaser", "--snapshot", "--skip-publish", "--rm-dist")
		if state.IsNotOk() {
			ux.PrintflnWarning("No goreleaser file found.")
			break
		}
	}

	return state
}
