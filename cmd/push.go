package cmd

import (
	"github.com/gearboxworks/buildtool/cmd/helpers"
	"github.com/gearboxworks/buildtool/ux"
	"github.com/spf13/cobra"
	"strings"
)


func init() {
	rootCmd.AddCommand(pushCmd)
}


var pushCmd = &cobra.Command{
	Use:   helpers.CmdPush,
	Short: ux.SprintfBlue("Push a gearboxworks repo."),
	Long: ux.SprintfBlue("Push a gearboxworks repo."),
	Run: Push,
}
func Push(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		tmpl := helpers.NewArgTemplate(Cmd.Debug)

		Cmd.State = tmpl.ProcessArgs(cmd, args)
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}


		git := helpers.NewGit(nil, Cmd.Debug, ".")
		Cmd.State = git.Open()
		if Cmd.State.IsNotOk() {
			break
		}
		ux.PrintflnBlue("Found git repo. Remote URL: %s", git.Url)


		Cmd.State = git.Push("%s", strings.Join(args, " "))
		if Cmd.State.IsNotOk() {
			break
		}
	}
}
