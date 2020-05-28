package main

import (
	"github.com/newclarity/buildtool/cmd"
	"github.com/newclarity/buildtool/defaults"
	"github.com/newclarity/buildtool/ux"
	"os"
	"strings"
)

func init() {
	_ = ux.Open(strings.ToUpper(defaults.BinaryName) + ": ")
}

func main() {
	state := cmd.Execute()
	if state.IsNotOk() {
		state.PrintResponse()
	}
	ux.Close()
	os.Exit(state.ExitCode)
}

/*

@TODO - Add 'setup' command that
@TODO -		Creates a Makefile.
@TODO -		Adds the Git 'assume-unchanged' thingy.
@TODO -		Sets up a GitHub action.
@TODO -		Creates buildtools directory.

@TODO - Add git ignore tag to the 'buildtool' binary.
@TODO -		git update-index --assume-unchanged buildtool

*/