package cmd

import (
	"fmt"
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


		Cmd.State = GitPush("%s", strings.Join(args, " "))
		if Cmd.State.IsNotOk() {
			break
		}
	}
}


func GitChanges() []string {
	state := ux.NewState(Cmd.Debug)
	var changes []string

	for range OnlyOnce {
		git := NewGit(nil, Cmd.Debug, ".")
		state = git.Open()
		if state.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Checking files in repo...")
		state = git.Exec("status", "--porcelain")
		if state.IsNotOk() {
			break
		}

		for _, fp := range state.OutputArray {
			s := strings.Fields(fp)
			if len(s) > 1 {
				changes = append(changes, fp)
			}
		}
	}

	return changes
}


func GitPush(comment string, args ...interface{}) *ux.State {
	state := ux.NewState(Cmd.Debug)

	for range OnlyOnce {
		c := fmt.Sprintf(comment, args...)
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

		changes := GitChanges()
		if len(changes) > 0 {
			ux.PrintflnBlue("Committing files to repo...")
			state = git.Exec("commit", "-m", c, ".")
			if state.IsNotOk() {
				break
			}
		}

		ux.PrintflnBlue("Pushing repo...")
		state = git.Exec("push")
		if state.IsNotOk() {
			break
		}
	}

	return state
}


func GitAddTag(version string, comment string, args ...interface{}) *ux.State {
	state := ux.NewState(Cmd.Debug)

	for range OnlyOnce {
		if version == "" {
			state.SetError("Missing tag version.")
			break
		}
		if !strings.HasPrefix(version, "v") {
			version = "v" + version
		}

		c := fmt.Sprintf(comment, args...)
		if c == "" {
			c = fmt.Sprintf("Commit before release %s", version)
		}

		git := NewGit(nil, Cmd.Debug, ".")
		state = git.Open()
		if state.IsNotOk() {
			break
		}
		ux.PrintflnBlue("Found git repo: %s", git.Url)

		ux.PrintflnBlue("Tagging version in repo...")
		state = git.Exec("tag", "-a", version, "-m", c)
		if state.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Pushing to origin...")
		state = git.Exec("push", "origin", version)
		if state.IsNotOk() {
			break
		}
	}

	return state
}


func GitDelTag(version string) *ux.State {
	state := ux.NewState(Cmd.Debug)

	for range OnlyOnce {
		if version == "" {
			state.SetError("Missing tag version.")
			break
		}
		if !strings.HasPrefix(version, "v") {
			version = "v" + version
		}

		git := NewGit(nil, Cmd.Debug, ".")
		state = git.Open()
		if state.IsNotOk() {
			break
		}

		if !IsTagExisting(version) {
			state.SetOk()
			break
		}

		ux.PrintflnBlue("Removing version tag in repo...")
		state = git.Exec("tag", "-d", version)
		if state.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Pushing to origin...")
		state = git.Exec("push", "origin", version)
		if state.IsNotOk() {
			break
		}

		state.SetOk()
	}

	return state
}


func IsTagExisting(version string) bool {
	state := ux.NewState(Cmd.Debug)
	var ok bool

	for range OnlyOnce {
		git := NewGit(nil, Cmd.Debug, ".")
		state = git.Open()
		if state.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Checking tag %s in repo...", version)
		state = git.Exec("tag", "-l", version)
		if state.IsNotOk() {
			break
		}

		if len(state.OutputArray) == 0 {
			ok = false
			break
		}

		ok = true
	}

	return ok
}
