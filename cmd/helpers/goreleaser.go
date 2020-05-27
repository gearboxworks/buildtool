package helpers

import (
	"github.com/newclarity/buildtool/ux"
	"strings"
)


func (at *ArgTemplate) GoReleaserRelease() *ux.State {
	for range OnlyOnce {
		exe := NewExecCommand(at.Debug)
		exe.ShowProgress()

		grFile := NewArgFile(at.Debug)
		at.State = grFile.SetPath(GoReleaserFile)
		if grFile.NotExists() {
			at.State = grFile.State
			break
		}

		ux.PrintflnBlue("Found goreleaser file: %s", GoReleaserFile)
		at.State = exe.Exec("goreleaser", "--rm-dist")
		if at.State.IsNotOk() {
			ux.PrintflnWarning("Error with goreleaser.")
			break
		}
	}

	return at.State
}


func (at *ArgTemplate) DiscoverVersion(lookfor string, path ...string) (string, *ux.State) {
	var version string

	for range OnlyOnce {
		grFile := NewArgFile(at.Debug)
		at.State = grFile.SetPath(path...)
		if grFile.NotExists() {
			at.State = grFile.State
			break
		}

		at.State = grFile.ReadFile()
		if at.State.IsNotOk() {
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

	return version, at.State
}


const GoReleaserFile = ".goreleaser.yml"
func (at *ArgTemplate) GoReleaserBuild() *ux.State {
	for range OnlyOnce {
		grFile := NewArgFile(at.Debug)
		at.State = grFile.SetPath(GoReleaserFile)
		if grFile.NotExists() {
			at.State = grFile.State
			break
		}

		exe := NewExecCommand(at.Debug)
		exe.ShowProgress()

		ux.PrintflnBlue("Found goreleaser file: %s", GoReleaserFile)
		at.State = exe.Exec("goreleaser", "--snapshot", "--skip-publish", "--rm-dist")
		if at.State.IsNotOk() {
			ux.PrintflnWarning("No goreleaser file found.")
			break
		}
	}

	return at.State
}
