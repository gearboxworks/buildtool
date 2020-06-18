// High level build tools.
package cmd

import (
	"github.com/newclarity/scribeHelpers/toolGhr"
	"github.com/newclarity/scribeHelpers/toolGo"
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

		state = gr.Build(false, path...)
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


func ReleaseCommit() *ux.State {
	state := Cmd.State

	for range onlyOnce {
		ux.PrintflnBlue("Committing changes prior to release.")

		// Fetch version from GoLang files.
		version := GoMeta.GetBinaryVersion()
		if version.IsNotValid() {
			state.SetError("BinaryVersion is invalid")
			break
		}

		// Sync GitHub repo.
		repo := GitOpen()
		if repo.State.IsNotOk() {
			break
		}
		state = GitCommit(repo, "Commit for Release %s", version.Name())
		if state.IsNotOk() {
			break
		}
		state = GitPush(repo)
		if state.IsNotOk() {
			break
		}
		state = GitDelTag(repo, version.Name())
		if state.IsNotOk() {
			break
		}
		state = GitAddTag(repo, version.Name(), "Release %s", version.Name())
		if state.IsNotOk() {
			break
		}
	}

	return state
}


func ReleaseGoReleaser(path ...string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		ux.PrintflnBlue("Running GoReleaser.")

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
		if version == "" {
			// Fetch version from GoLang files.
			version = GoMeta.GetBinaryVersion().Name()
			if version == "" {
				state.SetError("BinaryVersion is invalid")
				break
			}
		}

		if path == "" {
			//path = Cmd.WorkingPath.GetPath() + "/dist"
			path = Cmd.WorkingPath.GetPath() + "/dist"
		}

		if srcrepo == "" {
			// Fetch version from GoLang files.
			srcrepo = GoMeta.GetSourceRepo().GetUrl()
			if srcrepo == "" {
				state.SetError("BinaryVersion is invalid")
				break
			}
		}

		if binrepo == "" {
			// Fetch version from GoLang files.
			binrepo = GoMeta.GetBinaryRepo().GetUrl()
			if binrepo == "" {
				state.SetError("BinaryVersion is invalid")
				break
			}
		}

		ux.PrintflnBlue("Syncing Git repos...")
		if binrepo == srcrepo {
			ux.PrintflnBlue("Source and Binary repos identical, no action taken.")
			// No need to push to binary repo.
			// GoReleaser will handle this.
			break
		}
		ux.PrintflnBlue("Syncing Git repos...")
		ux.PrintflnBlue("Release version:  %s", version)
		ux.PrintflnBlue("Source repo:      %s", srcrepo)
		ux.PrintflnBlue("Binary repo:      %s", binrepo)
		ux.PrintflnBlue("Binary directory: %s", path)

		// Now sync the release in the destination repo.
		state = toolGhr.CopyReleases(srcrepo, version, binrepo, path)
	}

	return state
}


func argSourceRepo(srcrepo string) string {
	state := Cmd.State
	var ret string

	for range onlyOnce {
		foundRepo := GoMeta.GetSourceRepo()
		if srcrepo == "" {
			ret = foundRepo.GetUrl()
			break
		}

		var argRepo toolGo.Repo
		if err := argRepo.Set(srcrepo); err != nil {
			state.SetError("%s: %v", toolGo.SourceRepo, err)
			break
		}

		if foundRepo.IsNotSame(&argRepo) {
			ux.PrintflnWarning("Requested %s (%s) differs to found %s (%s)",
				toolGo.SourceRepo, argRepo.GetUrl(),
				toolGo.SourceRepo, foundRepo.GetUrl(),
			)
			break
		}
		ret = foundRepo.String()
	}

	Cmd.State = state
	return ret
}
