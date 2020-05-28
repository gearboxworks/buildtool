package helpers

import (
	"github.com/newclarity/buildtool/ux"
)

const GoReleaserFile = ".goreleaser.yml"


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
