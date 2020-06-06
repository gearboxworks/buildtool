package cmd

import (
	"github.com/newclarity/scribeHelpers/toolGit"
	"github.com/newclarity/scribeHelpers/ux"
)


func GitOpen() *toolGit.TypeGit {
	return toolGit.GitOpen(Cmd.WorkingPath.GetPath())
}


func GitClone(url string, path ...string) *toolGit.TypeGit {
	return toolGit.GitClone(url, path...)
}


func GitCommit(repo *toolGit.TypeGit, comment string, args ...interface{}) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		if repo == nil {
			repo = GitOpen()
			if repo.State.IsNotOk() {
				break
			}
		}

		ux.PrintflnBlue("Pushing to Git repo '%s'", repo.Url)
		state = repo.Commit([]string{"."}, comment, args...)
		if state.IsNotOk() {
			break
		}

		ux.PrintflnGreen("OK")
	}

	return state
}


func GitPush(repo *toolGit.TypeGit) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		if repo == nil {
			repo = GitOpen()
			if repo.State.IsNotOk() {
				break
			}
		}

		state = repo.Push()
		if state.IsNotOk() {
			break
		}

		ux.PrintflnGreen("OK")
	}

	return state
}


func GitPull(repo *toolGit.TypeGit) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		if repo == nil {
			repo = GitOpen()
			if repo.State.IsNotOk() {
				break
			}
		}

		ux.PrintflnBlue("Pulling Git repo '%s'", repo.Url)
		state = repo.Pull()
		if state.IsNotOk() {
			break
		}

		ux.PrintflnGreen("OK")
	}

	return state
}


func GitAddTag(repo *toolGit.TypeGit, tag string, comment string) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		if repo == nil {
			repo = GitOpen()
			if repo.State.IsNotOk() {
				break
			}
		}

		ux.PrintflnBlue("Add tag '%s' on Git repo '%s'", tag, repo.Url)
		state = repo.AddTag(tag, comment)
		if state.IsNotOk() {
			break
		}

		ux.PrintflnGreen("OK")
	}

	return state
}


func GitDelTag(repo *toolGit.TypeGit, tag string) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		if repo == nil {
			repo = GitOpen()
			if repo.State.IsNotOk() {
				break
			}
		}

		ux.PrintflnBlue("Delete tag '%s' on Git repo '%s'", tag, repo.Url)
		state = repo.DelTag(tag)
		if state.IsNotOk() {
			break
		}

		ux.PrintflnGreen("OK")
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
