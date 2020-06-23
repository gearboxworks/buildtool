//
package cmd

import (
	"github.com/newclarity/scribeHelpers/toolGo"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"strings"
)


//noinspection ALL
const (
	DefaultJsonFile 		= "buildtool.json"

	CmdBuild 			= "build"
	CmdPush 			= "push"
	CmdCommit 			= "commit"
	CmdPull 			= "pull"
	CmdClone 			= "clone"
	CmdVersion 			= "version"
	CmdRelease 			= "release"
	CmdGhr	 			= "ghr"
	CmdPkgReflect		= "pkgreflect"
	CmdGolang           = "go"
	CmdSync				= "sync"
	CmdVfsGen			= "vfsgen"

	CmdSet	 			= "set"

	CmdGet	 			= "get"
	CmdGetAll 			= "all"
	CmdGetName 			= "name"
	CmdGetVersion		= "version"
	CmdGetSourceRepo	= "src"
	CmdGetBinaryRepo	= "bin"

)


func init() {
	//// GitHub
	//rootCmd.AddCommand(pushCmd)
	//rootCmd.AddCommand(commitCmd)
	//rootCmd.AddCommand(pullCmd)
	//rootCmd.AddCommand(cloneCmd)
	//rootCmd.AddCommand(syncCmd)
	//
	////Workflow
	//rootCmd.AddCommand(buildCmd)
	//rootCmd.AddCommand(releaseCmd)
	//rootCmd.AddCommand(ghrCmd)
	//
	////GoLang
	//rootCmd.AddCommand(golangCmd)
	//rootCmd.AddCommand(pkgreflectCmd)
	//rootCmd.AddCommand(vfsGenCmd)
	//
	//rootCmd.AddCommand(setCmd)
	//setCmd.AddCommand(setNameCmd)
	//setCmd.AddCommand(setVersionCmd)
	//setCmd.AddCommand(setSourceRepoCmd)
	//setCmd.AddCommand(setBinaryRepoCmd)
	//
	//rootCmd.AddCommand(getCmd)
	//getCmd.AddCommand(getAllCmd)
	//getCmd.AddCommand(getNameCmd)
	//getCmd.AddCommand(getVersionCmd)
	//getCmd.AddCommand(getSourceRepoCmd)
	//getCmd.AddCommand(getBinaryRepoCmd)
}


var pushCmd = &cobra.Command{
	Use:   CmdPush,
	Short: ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Push a gearboxworks repo."),
	Long:  ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Push a gearboxworks repo."),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		CmdScribe.State = GitPush(nil)
	},
}

var commitCmd = &cobra.Command{
	Use:   CmdCommit,
	Short: ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Commit changes to a gearboxworks repo."),
	Long:  ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Commit changes to a gearboxworks repo."),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		CmdScribe.State = GitCommit(nil, strings.Join(args, " "))
		//if CmdScribe.State.IsNotOk() {
		//	return
		//}
		CmdScribe.State = GitPush(nil)
	},
}

var pullCmd = &cobra.Command{
	Use:   CmdPull,
	Short: ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Pull a gearboxworks repo."),
	Long:  ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Pull a gearboxworks repo."),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		CmdScribe.State = GitPull(nil)
	},
}

var cloneCmd = &cobra.Command{
	Use:   CmdClone,
	Short: ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Clone a gearboxworks repo."),
	Long:  ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Clone a gearboxworks repo."),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		repo := GitClone(args[0], CmdScribe.WorkingPath.GetPath())
		CmdScribe.State = repo.State
	},
}

var syncCmd = &cobra.Command{
	Use:   CmdSync,
	Short: ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Sync source and binary releases."),
	Long:  ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Sync source and binary releases.") + `
Arguments:
	version	- Release to sync.
	path	- Path to assets cache.
	srcrepo	- Source repo.
	binrepo	- Binary repo.`,
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		a := make([]string, 4)
		for i  := range args {
			a[i] = args[i]
		}
		CmdScribe.State = ReleaseSync(a[0], a[1], a[2], a[3])
	},
	Args:	cobra.RangeArgs(0, 4),
}


var buildCmd = &cobra.Command{
	Use:   CmdBuild,
	Short: ux.SprintfMagenta("Workflow") + ux.SprintfBlue(" - Build a gearboxworks repo."),
	Long: ux.SprintfMagenta("Workflow") + ux.SprintfBlue(" - Build a gearboxworks repo."),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		CmdScribe.State = Build(args...)
	},
}

var releaseCmd = &cobra.Command{
	Use:   CmdRelease,
	Short: ux.SprintfMagenta("Workflow") + ux.SprintfBlue(" - Release a gearboxworks repo."),
	Long: ux.SprintfMagenta("Workflow") + ux.SprintfBlue(" - Release a gearboxworks repo, (public and private repos)."),
	DisableFlagParsing: false,
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		CmdScribe.State = Release(args...)
	},
}

var ghrCmd = &cobra.Command{
	Use:   CmdGhr,
	Short: ux.SprintfMagenta("Workflow") + ux.SprintfBlue(" - Run github-release package."),
	Long:  ux.SprintfMagenta("Workflow") + ux.SprintfBlue(" - Run github-release package."),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		CmdScribe.State = GithubReleaser()
	},
}


var golangCmd = &cobra.Command{
	Use:   CmdGolang,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Run various GoLang commands."),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Run various GoLang commands."),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		CmdScribe.State = Golang(args...)
	},
}

var getCmd = &cobra.Command{
	Use:   CmdGet,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get values from src code."),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get values from src code."),
	Args: cobra.RangeArgs(0,1),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		err := cmd.Help()
		if err != nil {
			CmdScribe.State.SetError(err)
		}
	},
}
var getAllCmd = &cobra.Command{
	Use:   CmdGetAll,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get all values from src code."),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get all values from src code."),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		PrintMetaValue(toolGo.All, args...)
	},
}
var getNameCmd = &cobra.Command{
	Use:   CmdGetName,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get %s within src code.", toolGo.BinaryName),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get %s within src code.", toolGo.BinaryName),
	Example: ux.SprintfMagenta("%s %s", CmdGet, CmdGetName) + ux.SprintfBlue(" - Update the binary name within src.\n"),
	//Args:	cobra.ExactArgs( 1),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		PrintMetaValue(toolGo.BinaryName, args...)
	},
}
var getVersionCmd = &cobra.Command{
	Use:   CmdGetVersion,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get %s within src code.", toolGo.BinaryVersion),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get %s within src code.", toolGo.BinaryVersion),
	Example: ux.SprintfMagenta("%s %s", CmdGet, CmdGetVersion) + ux.SprintfBlue(" - Update the binary version within src.\n"),
	//Args:	cobra.ExactArgs( 1),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		PrintMetaValue(toolGo.BinaryVersion, args...)
	},
}
var getBinaryRepoCmd = &cobra.Command{
	Use:   CmdGetBinaryRepo,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get %s within src code.", toolGo.BinaryRepo),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get %s within src code.", toolGo.BinaryRepo),
	Example: ux.SprintfMagenta("%s %s", CmdGet, CmdGetBinaryRepo) + ux.SprintfBlue(" - Update the binary repo within src.\n"),
	//Args:	cobra.ExactArgs( 1),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		PrintMetaValue(toolGo.BinaryRepo, args...)
	},
}
var getSourceRepoCmd = &cobra.Command{
	Use:   CmdGetSourceRepo,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get %s within src code.", toolGo.SourceRepo),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get %s within src code.", toolGo.SourceRepo),
	Example: ux.SprintfMagenta("%s %s", CmdGet, CmdGetSourceRepo) + ux.SprintfBlue(" - Update the source repo within src.\n"),
	//Args:	cobra.ExactArgs( 1),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		PrintMetaValue(toolGo.SourceRepo, args...)
	},
}

var setCmd = &cobra.Command{
	Use:   CmdSet,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Sync source and binary releases."),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Sync source and binary releases."),
	Args: cobra.RangeArgs(0,1),
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			CmdScribe.State.SetError(err)
		}
	},
}
var setNameCmd = &cobra.Command{
	Use:   CmdGetName,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Set %s within src code.", toolGo.BinaryName),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Set %s within src code.", toolGo.BinaryName),
	Example: ux.SprintfMagenta("%s %s", CmdSet, CmdGetName) + ux.SprintfBlue(" - Update the binary name within src.\n"),
	Args:	cobra.ExactArgs( 1),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		CmdScribe.State = UpdateMeta(toolGo.BinaryName, args[0])
	},
}
var setVersionCmd = &cobra.Command{
	Use:   CmdGetVersion,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Set %s within src code.", toolGo.BinaryVersion),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Set %s within src code.", toolGo.BinaryVersion),
	Example: ux.SprintfMagenta("%s %s", CmdSet, CmdGetVersion) + ux.SprintfBlue(" - Update the binary version within src.\n"),
	Args:	cobra.ExactArgs( 1),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		CmdScribe.State = UpdateMeta(toolGo.BinaryVersion, args[0])
	},
}
var setBinaryRepoCmd = &cobra.Command{
	Use:   CmdGetBinaryRepo,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Set %s within src code.", toolGo.BinaryRepo),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Set %s within src code.", toolGo.BinaryRepo),
	Example: ux.SprintfMagenta("%s %s", CmdSet, CmdGetBinaryRepo) + ux.SprintfBlue(" - Update the binary repo within src.\n"),
	Args:	cobra.ExactArgs( 1),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		CmdScribe.State = UpdateMeta(toolGo.BinaryRepo, args[0])
	},
}
var setSourceRepoCmd = &cobra.Command{
	Use:   CmdGetSourceRepo,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Set %s within src code.", toolGo.SourceRepo),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Set %s within src code.", toolGo.SourceRepo),
	Example: ux.SprintfMagenta("%s %s", CmdSet, CmdGetSourceRepo) + ux.SprintfBlue(" - Update the source repo within src.\n"),
	Args:	cobra.ExactArgs( 1),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		CmdScribe.State = UpdateMeta(toolGo.SourceRepo, args[0])
	},
}

var pkgreflectCmd = &cobra.Command{
	Use:   CmdPkgReflect,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Run pkgreflect on a GoLang directory."),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Run pkgreflect on a GoLang directory."),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		CmdScribe.State = PkgReflect(args...)
	},
}

var vfsGenCmd = &cobra.Command{
	Use:   CmdVfsGen,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Run VfsGen on a GoLang directory."),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Run VfsGen on a GoLang directory."),
	Args: cobra.RangeArgs(1,2),
	Run: func(cmd *cobra.Command, args []string) {
		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			return
		}

		switch len(args) {
			case 0:
			case 1:
				CmdScribe.State = VfsGen(args[0], "")
			case 2:
				CmdScribe.State = VfsGen(args[0], args[1])
		}
	},
}
