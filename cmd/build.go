package cmd

import (
	"github.com/newclarity/scribeHelpers/toolGit"
	"github.com/newclarity/scribeHelpers/toolGoReleaser"
	"github.com/newclarity/scribeHelpers/ux"
)


func Build(path ...string) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		if len(path) == 0 {
			path = []string{"."}
		}

		gr := toolGoReleaser.New(Cmd.Runtime)
		if gr.State.IsNotOk() {
			state = gr.State
			break
		}

		state = gr.Build(path...)
		if state.IsNotOk() {
			break
		}
	}

	return state
}


func Release(path ...string) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		if len(path) == 0 {
			path = []string{"."}
		}


		var version string
		version, state = getBinaryVersion(DefaultVersionFile...)
		if state.IsNotOk() {
			break
		}


		git := toolGit.New(Cmd.Runtime)
		if git.State.IsNotOk() {
			state = git.State
			break
		}

		state = git.SetPath(path...)
		if state.IsNotOk() {
			break
		}

		state = git.Open()
		if state.IsNotOk() {
			break
		}
		ux.PrintflnBlue("Found git repo. Remote URL: %s", git.Url)


		state = git.Push("Pre-commit Release v%s", version)
		if state.IsNotOk() {
			break
		}


		state = git.DelTag(version)
		if state.IsNotOk() {
			break
		}


		state = git.AddTag(version, "Release %s", version)
		if state.IsNotOk() {
			break
		}


		gr := toolGoReleaser.New(Cmd.Runtime)
		if gr.State.IsNotOk() {
			state = gr.State
			break
		}

		gr.ShowProgress()
		state = gr.Release(path...)
		if state.IsNotOk() {
			break
		}
	}

	return state
}
