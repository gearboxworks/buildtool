// Version coding related build tools.
package cmd

import (
	"github.com/newclarity/scribeHelpers/toolGo"
	"github.com/newclarity/scribeHelpers/ux"
)


func PrintMetaValue(lookfor string, path ...string) *ux.State {
	state := CmdScribe.State

	for range onlyOnce {
		if lookfor == "" {
			lookfor = toolGo.All
		}

		GoMeta := FindMetaFile().GetMeta()
		if GoMeta.IsNotValid() {
			state.SetError("Current source files do not have any build meta data.")
			break
		}

		GoMeta.Print(lookfor)

		//name := meta.GetBinaryName()
		//ux.Printf("%v\n", name)
		//version := meta.GetBinaryVersion()
		//ux.Printf("%v\n", version)
		//binrepo := meta.GetBinaryRepo()
		//ux.PrintflnBlue("BinaryRepo URL: %s", binrepo.GetUrl())
		//ux.PrintflnBlue("BinaryRepo owner: %v", binrepo.GetOwner())
		//ux.PrintflnBlue("BinaryRepo name: %v", binrepo.GetName())
		//srcrepo := meta.GetSourceRepo()
		//ux.PrintflnBlue("SourceRepo URL: %s", srcrepo.GetUrl())
		//ux.PrintflnBlue("SourceRepo owner: %v", srcrepo.GetOwner())
		//ux.PrintflnBlue("SourceRepo name: %v", srcrepo.GetName())
		//ux.PrintflnBlue("\n\n\n")
	}
	return state
}


func FindMetaFile(path ...string) *toolGo.GoFile {
	var metaFile *toolGo.GoFile
	for range onlyOnce {
		goFiles := toolGo.New(CmdScribe.Runtime)
		if goFiles.State.IsNotOk() {
			CmdScribe.State = goFiles.State
			break
		}

		metaFile = goFiles.GetMetaFile(false, path...)
		if metaFile != nil {
			if metaFile.State.IsOk() {
				CmdScribe.State = goFiles.State
				break
			}
		}

		ux.PrintflnWarning("No build metadata found in path '%s'", toolGo.GetDefaultFile())
		ux.PrintflnBlue("Scanning all files for build metadata ...")
		metaFile = goFiles.GetMetaFile(true, ".")
		if metaFile == nil {
			ux.PrintflnError("No build metadata found in path '.'")
			break
		}
		if metaFile.State.IsOk() {
			CmdScribe.State = goFiles.State
			break
		}

		CmdScribe.State = goFiles.State
		//CmdScribe.State.SetError("Current source files do not have any build meta data.")
		ux.PrintflnError("No build metadata found in path '.'")
	}
	return metaFile
}
