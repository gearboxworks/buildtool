package cmd

import (
	"github.com/newclarity/buildtool/ux"
	"github.com/spf13/cobra"
	"strings"
)


func init() {
	rootCmd.AddCommand(pushCmd)
}


var pushCmd = &cobra.Command{
	Use:   CmdPush,
	Short: ux.SprintfBlue("Push a gearboxworks repo."),
	Long: ux.SprintfBlue("Push a gearboxworks repo."),
	Run: Push,
}
func Push(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		tmpl := ProcessArgs(cmd, args)
		Cmd.State = tmpl.State
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		// //////////
		Cmd.State = GitPush(args...)
		if Cmd.State.IsNotOk() {
			break
		}
	}
}


func GitPush(comment ...string) *ux.State {
	state := ux.NewState(Cmd.Debug)

	for range OnlyOnce {
		c := strings.Join(comment, " ")
		if c == "" {
			state.SetError("Missing comment to git commit.")
			break
		}

		git := NewGit(nil, Cmd.Debug, ".")
		state = git.Open()
		if state.IsNotOk() {
			break
		}
		ux.PrintflnBlue("Found git repo. Pushing to remote: %s", git.Url)

		ux.PrintflnBlue("Adding files to repo...")
		state = git.Exec("add", ".")
		if state.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Committing files to repo...")
		state = git.Exec("commit", "-m", c, ".")
		if state.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Pushing repo...")
		state = git.Exec("push")
		if state.IsNotOk() {
			break
		}
	}

	return state
}
