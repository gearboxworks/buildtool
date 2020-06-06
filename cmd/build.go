package cmd

import (
	"github.com/newclarity/scribeHelpers/toolGoReleaser"
	"github.com/newclarity/scribeHelpers/ux"
)


func Build(path ...string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		if len(path) == 0 {
			path = []string{Cmd.WorkingPath.GetPath()}
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

	for range onlyOnce {
		if len(path) == 0 {
			path = []string{Cmd.WorkingPath.GetPath()}
		}


		// Ensure repo builds properly.
		state = Build(path...)
		if state.IsNotOk() {
			state.SetError("Failed to build. Aborting...")
			break
		}


		// Fetch version from GoLang files.
		var version string
		version, state = getBinaryVersion(DefaultVersionFile...)
		if state.IsNotOk() {
			break
		}


		// Sync GitHub repo.
		repo := GitOpen()
		if repo.State.IsNotOk() {
			break
		}
		state = GitCommit(repo, "Commit for Release v%s", version)
		if state.IsNotOk() {
			break
		}
		state = GitPush(repo)
		if state.IsNotOk() {
			break
		}
		state = GitDelTag(repo, version)
		if state.IsNotOk() {
			break
		}
		state = GitAddTag(repo, version, "Release %s", version)
		if state.IsNotOk() {
			break
		}


		// Run GoReleaser.
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
