// Version coding related build tools.
package cmd

import (
	"github.com/newclarity/scribeHelpers/toolGo"
	"github.com/newclarity/scribeHelpers/ux"
)


func PrintMetaValue(lookfor string, path ...string) *ux.State {
	for range onlyOnce {
		if lookfor == "" {
			lookfor = toolGo.All
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
	return Cmd.State
}


func FindMetaFile(path ...string) *toolGo.GoFile {
	var metaFile *toolGo.GoFile
	for range onlyOnce {
		goFiles := toolGo.New(Cmd.Runtime)
		if goFiles.State.IsNotOk() {
			Cmd.State = goFiles.State
			break
		}

		metaFile = goFiles.GetMetaFile(false, path...)
		if metaFile != nil {
			if metaFile.State.IsOk() {
				Cmd.State = goFiles.State
				break
			}
		}

		ux.PrintflnWarning("No build metadata found in path '%s'", toolGo.GetDefaultFile())
		ux.PrintflnBlue("Scanning all files for build metadata ...")
		metaFile = goFiles.GetMetaFile(true, ".")
		if metaFile.State.IsOk() {
			Cmd.State = goFiles.State
			break
		}

		Cmd.State = goFiles.State
		ux.PrintflnError("No build metadata found in path '.'")
	}
	return metaFile
}


//func FindMeta(path ...string) *toolGo.GoMeta {
//	var meta *toolGo.GoMeta
//	for range onlyOnce {
//		goFiles := toolGo.New(Cmd.Runtime)
//		if goFiles.State.IsNotOk() {
//			Cmd.State = goFiles.State
//			break
//		}
//
//		meta = goFiles.GetMeta(false, path...)
//		if meta.Valid {
//			break
//		}
//
//		ux.PrintflnWarning("No build metadata found in path '%s'", toolGo.GetDefaultFile())
//		ux.PrintflnBlue("Scanning all files for build metadata ...")
//		meta = goFiles.GetMeta(true, ".")
//		if meta.Valid {
//			break
//		}
//
//		ux.PrintflnError("No build metadata found in path '.'")
//	}
//	return meta
//}
