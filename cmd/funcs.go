package cmd

import (
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"path/filepath"
)


func ProcessArgs(toolArgs *loadTools.TypeScribeArgs, cmd *cobra.Command, args []string) *ux.State {
	for range OnlyOnce {
		_ = toolArgs.Runtime.SetArgs(cmd.Use)
		_ = toolArgs.Runtime.AddArgs(args...)

		if len(args) >= 1 {
			ext := filepath.Ext(args[0])
			if ext == ".json" {
				toolArgs.Json.Filename = args[0]
			}
		}

		toolArgs.Template.Filename = loadTools.SelectIgnore

		toolArgs.ValidateArgs()
		if toolArgs.State.IsNotOk() {
			break
		}
	}

	return toolArgs.State
}
