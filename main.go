package main

import (
	"github.com/gearboxworks/buildtool/cmd"
	"github.com/gearboxworks/buildtool/defaults"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"strings"
)


func init() {
	_ = ux.Open(strings.ToUpper(defaults.BinaryName) + ": ")
}

func main() {
	state := cmd.Execute()
	state.PrintResponse()
	ux.Close()
	os.Exit(state.ExitCode)
}


/*
@TODO - When 'release' do some sanity checking on existing releases and offer a "do you want to remove?" option.

@TODO - Add a 'sync all releases' option to src -> binary repo release syncing.

@TODO - Add "upgrade to" command to enable semver upgrades.
@TODO - 	In particular, handle the defaults.version.go file better.

@TODO - Add 'setup' command that
@TODO -		Creates a Makefile.
@TODO -		Adds the Git 'assume-unchanged' thingy.
@TODO -		Sets up a GitHub action.
@TODO -		Creates buildtools directory.

@TODO - Add git ignore tag to the 'buildtool' binary.
@TODO -		git update-index --assume-unchanged buildtool

*/


/*

@DONE - BUG - pkgreflect args not being properly set.

*/


/*
************************************************************************
Notes:

Makefile:
        @echo "Pushing to: $(shell git branch)"
        @git config core.hooksPath .git-hooks

doc:
ifeq ($(GODOCMD),)
        @echo "godocdown - Installing"
        @go install github.com/robertkrimen/godocdown/godocdown
else
        @echo "godocdown - already installed here $(GODOCMD)"
endif
        @$(GODOCMD)


BINARYREPO="$(tools/getBinaryRepo.sh)"
USER="$(echo "${BINARYREPO}" | awk -F/ '{print$1}')"
REPO="$(echo "${BINARYREPO}" | awk -F/ '{print$1}')"

VERSION="$(tools/getBinaryVersion.sh)"

echo github-release release \
        --user "${USER}" \
        --repo "${GB_GITREPO}" \
        --tag "${VERSION}" \
        --name "Release ${VERSION}" \
        --description "${DESCRIPTION}"
*/