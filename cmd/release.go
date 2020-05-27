package cmd

import (
	"github.com/newclarity/buildtool/ux"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(releaseCmd)
}


// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   CmdRelease,
	Short: ux.SprintfBlue("Release a gearboxworks repo."),
	Long: ux.SprintfBlue("Release a gearboxworks repo, (public and private repos)."),
	Run: ReleaseTemplate,
	DisableFlagParsing: false,
}
func ReleaseTemplate(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		tmpl := ProcessArgs(cmd, args)
		Cmd.State = tmpl.State
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

	}
}
