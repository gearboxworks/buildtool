package cmd

import (
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/ux"
)


func PkgReflect(path ...string) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		pr := loadTools.PkgReflect {
			Notypes:    false,
			Nofuncs:    false,
			Novars:     false,
			Noconsts:   false,
			Unexported: false,
			Norecurs:   false,
			Stdout:     false,
			Gofile:     "",
			Notests:    false,
			Debug:      false,
			State:      nil,
		}
		state = loadTools.PackageReflect(pr, path...)
		if state.IsNotOk() {
			break
		}
	}

	return state
}
