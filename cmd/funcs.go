package cmd

import (
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"path/filepath"
)


func ProcessArgs(toolArgs *loadTools.TypeScribeArgs, cmd *cobra.Command, args []string) *ux.State {
	state := Cmd.State

	for range OnlyOnce {
		err := toolArgs.Runtime.SetArgs(cmd.Use)
		if err != nil {
			state.SetError(err)
			break
		}

		err = toolArgs.Runtime.AddArgs(args...)
		if err != nil {
			state.SetError(err)
			break
		}

		if len(args) >= 1 {
			ext := filepath.Ext(args[0])
			if ext == ".json" {
				toolArgs.Json.Filename = args[0]
			}
		}

		toolArgs.Template.Filename = loadTools.SelectIgnore

		state = toolArgs.ValidateArgs()
		if state.IsNotOk() {
			break
		}
	}

	return state
}
