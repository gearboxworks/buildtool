package cmd

import (
	"github.com/newclarity/scribeHelpers/toolGhr"
	"github.com/newclarity/scribeHelpers/toolGoReleaser"
	"github.com/newclarity/scribeHelpers/ux"
)


func Build(path ...string) *ux.State {
	state := Cmd.State

	//state = ReleaseSync("", "mickmake/test", "v1.1.0", "")
	//os.Exit(1)

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

		// Commit changes.
		state = ReleaseCommit()
		if state.IsNotOk() {
			break
		}

		// Run GoReleaser.
		state = ReleaseGoReleaser(path...)
		if state.IsNotOk() {
			break
		}

		// Run GHR - copy release to binary repo.
		state = ReleaseSync("", "", "", "")
		if state.IsNotOk() {
			break
		}
	}

	return state
}


func ReleaseCommit(path ...string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		// Fetch version from GoLang files.
		var version string
		version, state = getBinaryVersion()
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
	}

	return state
}


func ReleaseGoReleaser(path ...string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		if len(path) == 0 {
			path = []string{Cmd.WorkingPath.GetPath()}
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


func ReleaseSync(version string, path string, srcrepo string, binrepo string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		if srcrepo == "" {
			srcrepo, state = getSourceRepo()
			if state.IsNotOk() {
				break
			}
		}

		if binrepo == "" {
			binrepo, state = getBinaryRepo()
			if state.IsNotOk() {
				break
			}
		}

		if version == "" {
			version, state = getBinaryVersion()
			if state.IsNotOk() {
				break
			}
		}

		if path == "" {
			//path = Cmd.WorkingPath.GetPath() + "/dist"
			path = Cmd.WorkingPath.GetPath() + "/dist"
		}

		ux.PrintflnBlue("Syncing Git repos...")
		if binrepo == srcrepo {
			ux.PrintflnBlue("Source and Binary repos identical, no action taken.")
			// No need to push to binary repo.
			// GoReleaser will handle this.
			break
		}
		ux.PrintflnBlue("Syncing Git repos...")
		ux.PrintflnBlue("Source repo:	 %s", srcrepo)
		ux.PrintflnBlue("Binary repo:	 %s", binrepo)
		ux.PrintflnBlue("Release version: %s", version)
		ux.PrintflnBlue("Asset directory: %s", path)

		// Now sync the release in the destination repo.
		state = toolGhr.CopyReleases(srcrepo, version, binrepo, path)
	}

	return state
}
