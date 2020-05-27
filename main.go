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

@TODO - Add additional capability to read a *.jtc file which contains both json and template.
@TODO -		- Maybe divide with the following:
@TODO -				- '^-- JSON:BEGIN --$'
@TODO -				- '^-- JSON:END --$'
@TODO -				- '^-- TEMPLATE:BEGIN --$'
@TODO -				- '^-- TEMPLATE:END --$'

*/