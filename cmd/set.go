package cmd

import (
	"github.com/newclarity/scribeHelpers/ux"
)


func UpdateMeta(lookfor string, value string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		state = GoMetaFile.UpdateMeta(lookfor, value)
		if state.IsNotOk() {
			break
		}

		ux.PrintflnOk("File updated OK")
	}

	return state
}
