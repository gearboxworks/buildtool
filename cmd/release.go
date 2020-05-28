package cmd

import (
	"github.com/gearboxworks/buildtool/cmd/helpers"
	"github.com/gearboxworks/buildtool/ux"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(releaseCmd)
}


// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   helpers.CmdRelease,
	Short: ux.SprintfBlue("Release a gearboxworks repo."),
	Long: ux.SprintfBlue("Release a gearboxworks repo, (public and private repos)."),
	Run: Release,
	DisableFlagParsing: false,
}
func Release(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		var version string
		tmpl := helpers.NewArgTemplate(Cmd.Debug)

		Cmd.State = tmpl.ProcessArgs(cmd, args)
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}


		version, Cmd.State = tmpl.GetBinaryVersion(helpers.DefaultVersionFile...)
		if Cmd.State.IsNotOk() {
			break
		}


		git := helpers.NewGit(nil, Cmd.Debug, ".")
		Cmd.State = git.Open()
		if Cmd.State.IsNotOk() {
			break
		}
		ux.PrintflnBlue("Found git repo. Remote URL: %s", git.Url)


		Cmd.State = git.Push("Release commit v%s)", version)
		if Cmd.State.IsNotOk() {
			break
		}


		Cmd.State = git.DelTag(version)
		if Cmd.State.IsNotOk() {
			break
		}


		Cmd.State = git.AddTag(version, "Release %s", version)
		if Cmd.State.IsNotOk() {
			break
		}


		Cmd.State = tmpl.GoReleaserRelease()
		if Cmd.State.IsNotOk() {
			break
		}
	}
}
