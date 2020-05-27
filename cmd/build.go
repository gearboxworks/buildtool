package cmd

import (
	"github.com/newclarity/buildtool/cmd/helpers"
	"github.com/newclarity/buildtool/ux"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(buildCmd)
}


var buildCmd = &cobra.Command{
	Use:   helpers.CmdBuild,
	Short: ux.SprintfBlue("Build a gearboxworks repo."),
	Long: ux.SprintfBlue("Build a gearboxworks repo."),
	Run: Build,
}
func Build(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		tmpl := helpers.NewArgTemplate(Cmd.Debug)

		Cmd.State = tmpl.ProcessArgs(cmd, args)
		if Cmd.State.IsNotOk() {
			break
		}
		//Cmd.State.DebugPrint()


		Cmd.State = tmpl.GoReleaserBuild()
		if Cmd.State.IsNotOk() {
			break
		}
	}
}
