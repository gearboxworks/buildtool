package cmd

import (
	"github.com/newclarity/buildtool/ux"
	"github.com/spf13/cobra"
	"strings"
)


func init() {
	rootCmd.AddCommand(releaseCmd)
}


// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   CmdRelease,
	Short: ux.SprintfBlue("Release a gearboxworks repo."),
	Long: ux.SprintfBlue("Release a gearboxworks repo, (public and private repos)."),
	Run: Release,
	DisableFlagParsing: false,
}
func Release(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		var version string

		tmpl := ProcessArgs(cmd, args)
		Cmd.State = tmpl.State
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}


		version, Cmd.State = DiscoverVersion("BinaryVersion", "defaults", "version.go")
		if Cmd.State.IsNotOk() {
			break
		}


		Cmd.State = GitPush("Commit before release v%s", version)
		if Cmd.State.IsNotOk() {
			break
		}


		Cmd.State = GitTag(version, "Release ", version)
		if Cmd.State.IsNotOk() {
			break
		}


		Cmd.State = GoReleaserRelease()
		if Cmd.State.IsNotOk() {
			break
		}
	}
}


func GoReleaserRelease() *ux.State {
	state := ux.NewState(Cmd.Debug)

	for range OnlyOnce {
		exe := NewExecCommand(Cmd.Debug)
		exe.ShowProgress()

		grFile := NewArgFile(Cmd.Debug)
		state = grFile.SetPath(GoReleaserFile)
		if grFile.NotExists() {
			state = grFile.State
			break
		}

		ux.PrintflnBlue("Found goreleaser file: %s", GoReleaserFile)
		state = exe.Exec("goreleaser", "--rm-dist")
		if state.IsNotOk() {
			ux.PrintflnWarning("Error with goreleaser.")
			break
		}
	}

	return state
}


func DiscoverVersion(lookfor string, path ...string) (string, *ux.State) {
	state := ux.NewState(Cmd.Debug)
	var version string

	for range OnlyOnce {
		grFile := NewArgFile(Cmd.Debug)
		state = grFile.SetPath(path...)
		if grFile.NotExists() {
			state = grFile.State
			break
		}

		state = grFile.ReadFile()
		if state.IsNotOk() {
			break
		}

		for _, v := range strings.Split(grFile.String, "\n") {
			if !strings.Contains(v, lookfor) {
				continue
			}

			sa := strings.Split(v, "=")
			if len(sa) != 2 {
				continue
			}

			version = strings.TrimSpace(sa[1])
			version = strings.TrimPrefix(version, "\"")
			version = strings.TrimSuffix(version, "\"")
			break
		}

		//
	}

	return version, state
}
