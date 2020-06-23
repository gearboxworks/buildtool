// High level build tools.
package cmd

import (
	"github.com/newclarity/scribeHelpers/toolGhr"
	"github.com/newclarity/scribeHelpers/toolGo"
	"github.com/newclarity/scribeHelpers/toolGoReleaser"
	"github.com/newclarity/scribeHelpers/ux"
)


func Build(path ...string) *ux.State {
	state := CmdScribe.State

	for range onlyOnce {
		if len(path) == 0 {
			path = []string{CmdScribe.WorkingPath.GetPath()}
		}

		gr := toolGoReleaser.New(CmdScribe.Runtime)
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
	state := CmdScribe.State

	for range onlyOnce {
		if len(path) == 0 {
			path = []string{CmdScribe.WorkingPath.GetPath()}
		}

		// Ensure repo builds properly.
		state = Build(path...)
		if state.IsNotOk() {
			state.SetError("Failed to build. Aborting...")
			break
		}

		GoMeta := FindMetaFile().GetMeta()
		if GoMeta.IsNotValid() {
			state.SetError("Current source files do not have any build meta data.")
			break
		}

		// Commit changes.
		state = ReleaseCommit(GoMeta)
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


func ReleaseCommit(meta *toolGo.GoMeta) *ux.State {
	state := CmdScribe.State

	for range onlyOnce {
		ux.PrintflnBlue("Committing changes prior to release.")

		// Fetch version from GoLang files.
		version := meta.GetBinaryVersion()
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
	state := CmdScribe.State

	for range onlyOnce {
		ux.PrintflnBlue("Running GoReleaser.")

		if len(path) == 0 {
			path = []string{CmdScribe.WorkingPath.GetPath()}
		}

		// Run GoReleaser.
		gr := toolGoReleaser.New(CmdScribe.Runtime)
		if gr.State.IsNotOk() {
			state = gr.State
			break
		}

		gr.ShowProgress()
		state = gr.Release(false, path...)
		if state.IsNotOk() {
			break
		}

	}

	return state
}


func ReleaseSync(version string, path string, srcrepo string, binrepo string) *ux.State {
	state := CmdScribe.State

	for range onlyOnce {
		GoMeta := FindMetaFile().GetMeta()
		if GoMeta.IsNotValid() {
			state.SetError("Current source files do not have any build meta data.")
			break
		}

		if version == "" {
			// Fetch version from GoLang files.
			version = GoMeta.GetBinaryVersion().Name()
			if version == "" {
				state.SetError("BinaryVersion is invalid")
				break
			}
		}

		if path == "" {
			//path = CmdScribe.WorkingPath.GetPath() + "/dist"
			path = CmdScribe.WorkingPath.GetPath() + "/dist"
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
