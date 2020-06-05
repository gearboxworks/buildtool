package cmd

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/toolGit"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)


func GitPush(path string, comment ...string) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		git := toolGit.New(Cmd.Runtime)

		state = git.SetPath(path)
		if state.IsNotOk() {
			break
		}

		state = git.Open()
		if state.IsNotOk() {
			break
		}
		ux.PrintflnBlue("Found git repo. Remote URL: %s", git.Url)

		ctxt := fmt.Sprintf("%s", strings.Join(comment, " "))
		state = git.Push(ctxt)
		if state.IsNotOk() {
			break
		}
	}

	return state
}


func GitTag(path string, comment ...string) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		//git := toolGit.New(Cmd.Runtime)
		//
		//state = git.SetPath(path)
		//if state.IsNotOk() {
		//	break
		//}
		//
		//state = git.Open()
		//if state.IsNotOk() {
		//	break
		//}
		//ux.PrintflnBlue("Found git repo. Remote URL: %s", git.Url)
	}

	return state
}


func GithubReleaser(args ...string) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		//tmpl := loadTools.NewArgTemplate(Cmd.Debug)
		//
		//state = tmpl.ProcessArgs(cmd, args)
		//if state.IsNotOk() {
		//	state.PrintResponse()
		//	break
		//}

		//var repo string
		//repo, state = tmpl.GetBinaryRepo(helpers.DefaultVersionFile...)
		//if state.IsNotOk() {
		//	break
		//}
		//
		//ux.PrintflnBlue("Repo Url: %s", repo)
		//owner, name := helpers.GetRepoComponents(repo)
		//ux.PrintflnBlue("Repo Owner: %s", owner)
		//ux.PrintflnBlue("Repo Name: %s", name)
		//
		//ghr := helpers.NewGitHubRelease(Cmd.Debug)
		//if ghr.State.IsNotOk() {
		//	break
		//}
		//ghr.SetOwner(owner)
		//ghr.SetRepo(name)
		//state = ghr.Run("info")
		//if state.IsNotOk() {
		//	//
		//}
		//
		//
		//if len(args) == 0 {
		//	break
		//}
		//
		//if args[0] == "all" {
		//	args = []string{"name", "version", "src", "bin"}
		//}

	}

	return state
}
