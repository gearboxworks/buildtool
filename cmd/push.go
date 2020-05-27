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


		git := NewGit(nil, Cmd.Debug, ".")
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


func (g *TypeGit) FileChanges() []string {
	state := ux.NewState(Cmd.Debug)
	var changes []string

	for range OnlyOnce {
		ux.PrintflnBlue("Checking files in repo...")
		state = g.Exec("status", "--porcelain")
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


func (g *TypeGit) Push(comment string, args ...interface{}) *ux.State {
	state := ux.NewState(Cmd.Debug)

	for range OnlyOnce {
		c := fmt.Sprintf(comment, args...)
		if c == "" {
			state.SetError("Missing comment to git commit.")
			break
		}

		ux.PrintflnBlue("Adding files to repo...")
		state = g.Exec("add", ".")
		if state.IsNotOk() {
			break
		}

		changes := g.FileChanges()
		if len(changes) > 0 {
			ux.PrintflnBlue("Committing files to repo...")
			state = g.Exec("commit", "-m", c, ".")
			if state.IsNotOk() {
				break
			}
		}

		ux.PrintflnBlue("Pushing repo...")
		state = g.Exec("push")
		if state.IsNotOk() {
			break
		}
	}

	return state
}


func (g *TypeGit) AddTag(version string, comment string, args ...interface{}) *ux.State {
	state := ux.NewState(Cmd.Debug)

	for range OnlyOnce {
		if version == "" {
			state.SetError("Missing tag version.")
			break
		}
		if !strings.HasPrefix(version, "v") {
			version = "v" + version
		}

		if g.IsTagExisting(version) {
			state.SetOk()
			break
		}


		c := fmt.Sprintf(comment, args...)
		if c == "" {
			c = fmt.Sprintf("Commit before release %s", version)
		}

		ux.PrintflnBlue("Tagging version in repo...")
		state = g.Exec("tag", "-a", version, "-m", c)
		if state.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Pushing to origin...")
		state = g.Exec("push", "origin", version)
		if state.IsNotOk() {
			break
		}
	}

	return state
}


func (g *TypeGit) DelTag(version string) *ux.State {
	state := ux.NewState(Cmd.Debug)

	for range OnlyOnce {
		if version == "" {
			state.SetError("Missing tag version.")
			break
		}
		if !strings.HasPrefix(version, "v") {
			version = "v" + version
		}

		if !g.IsTagExisting(version) {
			state.SetOk()
			break
		}

		ux.PrintflnBlue("Removing version tag in repo...")
		state = g.Exec("tag", "-d", version)
		if state.IsNotOk() {
			break
		}

		state.SetOk()
	}

	return state
}


func (g *TypeGit) IsTagExisting(version string) bool {
	state := ux.NewState(Cmd.Debug)
	var ok bool

	for range OnlyOnce {
		ux.PrintflnBlue("Checking tag %s in repo...", version)
		state = g.Exec("tag", "-l", version)
		if state.IsNotOk() {
			break
		}

		if !strings.HasPrefix(version, "v") {
			version = "v" + version
		}

		ok = false
		for _, t := range state.OutputArray {
			if t == version {
				ok = true
				break
			}
		}
	}

	return ok
}
