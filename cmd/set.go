package cmd

import (
	"github.com/newclarity/scribeHelpers/ux"
)


func UpdateMeta(lookfor string, value string) *ux.State {
	state := CmdScribe.State

	for range onlyOnce {
		GoMetaFile := FindMetaFile()
		if GoMetaFile.State.IsNotOk() {
			state.SetError("Current source files do not have any build meta data.")
			break
		}

		state = GoMetaFile.UpdateMeta(lookfor, value)
		if state.IsNotOk() {
			break
		}

		ux.PrintflnOk("File updated OK")
	}

	return state
}
