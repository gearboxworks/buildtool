package cmd

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/toolGhr"
	"github.com/newclarity/scribeHelpers/toolGoReleaser"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"strings"
)


func Build(path ...string) *ux.State {
	state := Cmd.State

	state = ReleaseGhr(path...)
	os.Exit(1)

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
		state = ReleaseGhr(path...)
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


func ReleaseGhr(path ...string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		if len(path) == 0 {
			path = []string{Cmd.WorkingPath.GetPath()}
		}

		var version string
		version, state = getBinaryVersion()
		if state.IsNotOk() {
			break
		}

		var binrepo string
		binrepo, state = getBinaryRepo()
		if state.IsNotOk() {
			break
		}
		binrepo = "mickmake/test"

		var srcrepo string
		srcrepo, state = getSourceRepo()
		if state.IsNotOk() {
			break
		}

		if binrepo == srcrepo {
			// No need to push to binary repo.
			break
		}

		// Run GHR - copy release to binary repo.
		ghr := toolGhr.New(Cmd.Runtime)
		if ghr.State.IsNotOk() {
			state = ghr.State
			break
		}
		state = ghr.Open("mickmake", "test")

		br := strings.Split(binrepo, "/")
		release := toolGhr.TypeRepo {
			Organization: br[0],
			Name:         br[1],
			TagName:      version,
			Description:  fmt.Sprintf("Release '%s' copied from src repo '%s'", version, srcrepo),
			Draft:        false,
			Prerelease:   false,
			Target:       "",
			Replace:      true,
			Files:        []string{"testing", "pkgreflect.go", "init.go"},
			Auth:         &toolGhr.TypeAuth{ Token: "", AuthUser: "" },
		}

		//state = ghr.OpenUrl("mickmake/test")
		//if ghr.State.IsNotOk() {
		//	break
		//}

		state = ghr.CreateRelease(release)
		if ghr.State.IsNotOk() {
			break
		}

	}

	return state
}
