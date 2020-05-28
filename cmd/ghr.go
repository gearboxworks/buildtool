package cmd

import (
	"github.com/gearboxworks/buildtool/cmd/helpers"
	"github.com/gearboxworks/buildtool/ux"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(ghrCmd)
}


var ghrCmd = &cobra.Command{
	Use:   helpers.CmdGhr,
	Short: ux.SprintfBlue("Get value from GoLang src."),
	Long:  ux.SprintfBlue("Get value from GoLang src code."),
	Run:   RunGhr,
}
func RunGhr(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		tmpl := helpers.NewArgTemplate(Cmd.Debug)

		Cmd.State = tmpl.ProcessArgs(cmd, args)
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		var repo string
		repo, Cmd.State = tmpl.GetBinaryRepo(helpers.DefaultVersionFile...)
		if Cmd.State.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Repo Url: %s", repo)
		owner, name := helpers.GetRepoComponents(repo)
		ux.PrintflnBlue("Repo Owner: %s", owner)
		ux.PrintflnBlue("Repo Name: %s", name)

		ghr := helpers.NewGitHubRelease(Cmd.Debug)
		if ghr.State.IsNotOk() {
			break
		}
		ghr.SetOwner(owner)
		ghr.SetRepo(name)
		Cmd.State = ghr.Run("info")
		if Cmd.State.IsNotOk() {
			//
		}


		if len(args) == 0 {
			break
		}

		if args[0] == "all" {
			args = []string{"name", "version", "src", "bin"}
		}

	}
}
