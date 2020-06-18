// GoLang related build tools.
package cmd

import (
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/toolExec"
	"github.com/newclarity/scribeHelpers/toolGo"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/shurcooL/vfsgen"
	"net/http"
	"path/filepath"
	"strings"
)


func Golang(args ...string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		if len(args) == 0 {
			goLangHelp()
			break
		}

		switch strings.ToLower(args[0]) {
			case "update":
				state = goLangUpdate(args[1:]...)

			default:
				goLangHelp()
		}
	}

	return state
}


func goLangHelp() *ux.State {
	state := Cmd.State

	ux.PrintflnYellow("Need to supply one of:")
	ux.PrintflnYellow("\tupdate - Detect go.mod files recursively and update modules.")

	return state
}


func goLangUpdate(path ...string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		if len(path) == 0 {
			path = []string{"."}
		}

		ux.PrintflnBlue("Updating go modules...")

		e := toolExec.NewMultiExec(Cmd.Runtime)
		if e.State.IsNotOk() {
			state = e.State
			break
		}

		state = e.Set("go", "get", "-u")
		if state.IsNotOk() {
			break
		}

		state = e.SetDontAppendFile()
		if state.IsNotOk() {
			break
		}

		state = e.SetChdir()
		if state.IsNotOk() {
			break
		}

		state = e.ShowProgress()
		if state.IsNotOk() {
			break
		}

		state = e.FindRegex("go.mod", path...)
		if state.IsNotOk() {
			break
		}

		p := e.GetPaths()
		ux.PrintflnBlue("Updating go modules in %d paths...", len(p))

		state = e.Run()
		if state.IsNotOk() {
			break
		}

		state.SetOk("go module update OK")
	}

	return state
}


func PkgReflect(paths ...string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		if len(paths) == 0 {
			paths = []string{"."}
		}

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
		state = loadTools.PackageReflect(pr, paths...)
		if state.IsNotOk() {
			break
		}
	}

	return state
}


func VfsGen(assetDir string, goDir string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		if len(assetDir) == 0 {
			break
		}
		ux.PrintflnBlue("Running VFSGEN...")


		// Fetch the assets directory.
		dir := toolPath.New(nil)
		if dir.State.IsNotOk() {
			state = dir.State
			break
		}

		dir.SetPath(assetDir)
		state = dir.StatPath()
		if state.IsNotOk() {
			break
		}
		ux.PrintflnBlue("VFSGEN: Asset directory is '%s'.", dir.GetPath())


		// Fetch the GoLang destination directory.
		if goDir == "" {
			//goDir = dir.GetParentDirAbs()
			goDir = "."
		}

		goFiles := toolGo.New(Cmd.Runtime)
		if goFiles.State.IsError() {
			state = goFiles.State
			break
		}

		packageName := goFiles.GetPackageName(goDir)
		if packageName == "" {
			packageName = filepath.Base(goDir)
			//state.SetError("No Go files in destination directory.")
			//break
		}
		ux.PrintflnBlue("VFSGEN: GoLang directory is '%s'.", goFiles.GetPath())


		// Run vfsgen
		ux.PrintflnBlue("VFSGEN: Generating code into file '%s'.", filepath.Join(goFiles.GetPath(), "vfsdata.go"))
		fs := http.Dir(dir.GetDirname())
		err := vfsgen.Generate(fs, vfsgen.Options{
			Filename:        filepath.Join(goFiles.GetPath(), "vfsdata.go"),
			//Filename:        "vfsdata.go",
			PackageName:     packageName,
			BuildTags:       "",
			VariableName:    "",
			VariableComment: "",
		})

		if err != nil {
			state.SetError(err)
			ux.PrintflnError("%s", err)
			break
		}

		state.SetOk()
	}

	return state
}
