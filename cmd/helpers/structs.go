package helpers

import (
	"github.com/newclarity/buildtool/cmd/runtime"
	"github.com/newclarity/buildtool/ux"
	"github.com/spf13/cobra"
	"path/filepath"
)

const OnlyOnce = "1"


type ArgTemplate struct {
	Exec            *runtime.Exec

	Json           *TypeArgFile

	ExecShell      bool // Cmd: "run"
	Chdir          bool // Flag: --chdir
	RemoveTemplate bool // Flag: --rm-tmpl
	ForceOverwrite bool // Flag: --force
	RemoveOutput   bool // Flag: --rm-out
	QuietProgress  bool // Flag: --quiet
	Debug          bool // Flag: --debug

	HelpAll        bool
	HelpFunctions  bool
	HelpVariables  bool
	HelpExamples   bool

	State           *ux.State
	valid           bool
}


func NewArgTemplate(debugMode bool) *ArgTemplate {

	p := ArgTemplate{
		Exec:           runtime.NewExec(debugMode),

		Json:           NewArgFile(debugMode),

		ExecShell:      false,
		Chdir:          false,
		RemoveTemplate: false,
		ForceOverwrite: false,
		RemoveOutput:   false,
		QuietProgress:  false,
		Debug:          false,

		HelpAll:        false,
		HelpFunctions:  false,
		HelpVariables:  false,
		HelpExamples:   false,

		State:          ux.NewState(debugMode),
		valid:          false,
	}

	p.State.SetPackage("")
	p.State.SetFunctionCaller()

	return &p
}


func (at *ArgTemplate) ProcessArgs(cmd *cobra.Command, args []string) *ux.State {
	for range OnlyOnce {
		_ = at.Exec.SetArgs(cmd.Use)
		_ = at.Exec.AddArgs(args...)

		if len(args) >= 1 {
			ext := filepath.Ext(args[0])
			if ext == ".json" {
				at.Json.SetPath(args[0])
			}
		}

		at.ValidateArgs()
		if at.State.IsNotOk() {
			break
		}
	}

	return at.State
}


func (at *ArgTemplate) ValidateArgs() *ux.State {

	for range OnlyOnce {
		at.State.Clear()

		// Validate json and template files/strings.
		if at.Json.Path == SelectDefaultJson {
			at.Json.Path = DefaultJsonString
		}

		at.State = at.Json.SetPath(at.Json.Path)
		if at.State.IsOk() {
			at.State = at.Json.ReadFile()
			if at.State.IsNotOk() {
				break
			}
		}
		at.State.SetOk()	// Ignore if not supplied.


		////////////////////////////////////////////////////
		// Chdir.
		if at.Chdir {
			at.State = at.Json.ChDir()
			if at.State.IsError() {
				break
			}
		}

		at.State.SetOk("")
	}

	return at.State
}
