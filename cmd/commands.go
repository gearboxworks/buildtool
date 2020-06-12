package cmd

import (
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
	CmdGet	 			= "get"
	CmdGhr	 			= "ghr"
	CmdPkgReflect		= "pkgreflect"
	CmdGolang           = "go"
	CmdSync				= "sync"
)


func init() {
	rootCmd.AddCommand(ghrCmd)
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(cloneCmd)
	rootCmd.AddCommand(golangCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(releaseCmd)
	rootCmd.AddCommand(pkgreflectCmd)
	rootCmd.AddCommand(syncCmd)
}


var pushCmd = &cobra.Command{
	Use:   CmdPush,
	Short: ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Push a gearboxworks repo."),
	Long:  ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Push a gearboxworks repo."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			return
		}

		Cmd.State = GitPush(nil)
	},
}

var commitCmd = &cobra.Command{
	Use:   CmdCommit,
	Short: ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Commit changes to a gearboxworks repo."),
	Long:  ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Commit changes to a gearboxworks repo."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			return
		}

		Cmd.State = GitCommit(nil, strings.Join(args, " "))
		//if Cmd.State.IsNotOk() {
		//	return
		//}
		Cmd.State = GitPush(nil)
	},
}

var pullCmd = &cobra.Command{
	Use:   CmdPull,
	Short: ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Pull a gearboxworks repo."),
	Long:  ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Pull a gearboxworks repo."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			return
		}

		Cmd.State = GitPull(nil)
	},
}

var cloneCmd = &cobra.Command{
	Use:   CmdClone,
	Short: ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Clone a gearboxworks repo."),
	Long:  ux.SprintfMagenta("GitHub") + ux.SprintfBlue(" - Clone a gearboxworks repo."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			return
		}

		repo := GitClone(args[0], Cmd.WorkingPath.GetPath())
		Cmd.State = repo.State
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
		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			return
		}
		a := make([]string, 4)
		for i  := range args {
			a[i] = args[i]
		}
		Cmd.State = ReleaseSync(a[0], a[1], a[2], a[3])
	},
	Args:	cobra.RangeArgs(0, 4),
}


var buildCmd = &cobra.Command{
	Use:   CmdBuild,
	Short: ux.SprintfMagenta("Workflow") + ux.SprintfBlue(" - Build a gearboxworks repo."),
	Long: ux.SprintfMagenta("Workflow") + ux.SprintfBlue(" - Build a gearboxworks repo."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			return
		}

		Cmd.State = Build(args...)
	},
}

var releaseCmd = &cobra.Command{
	Use:   CmdRelease,
	Short: ux.SprintfMagenta("Workflow") + ux.SprintfBlue(" - Release a gearboxworks repo."),
	Long: ux.SprintfMagenta("Workflow") + ux.SprintfBlue(" - Release a gearboxworks repo, (public and private repos)."),
	DisableFlagParsing: false,
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			return
		}

		Cmd.State = Release(args...)
	},
}

var ghrCmd = &cobra.Command{
	Use:   CmdGhr,
	Short: ux.SprintfMagenta("Workflow") + ux.SprintfBlue(" - Run github-release package."),
	Long:  ux.SprintfMagenta("Workflow") + ux.SprintfBlue(" - Run github-release package."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = GithubReleaser()
	},
}


var golangCmd = &cobra.Command{
	Use:   CmdGolang,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Run various GoLang commands."),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Run various GoLang commands."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			return
		}

		Cmd.State = Golang(args...)
	},
}

var getCmd = &cobra.Command{
	Use:   CmdGet,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get values from GoLang src code."),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Get values from GoLang src code."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			return
		}

		Cmd.State = Get(args...)
	},
}

var pkgreflectCmd = &cobra.Command{
	Use:   CmdPkgReflect,
	Short: ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Run pkgreflect on a GoLang directory."),
	Long:  ux.SprintfMagenta("GoLang") + ux.SprintfBlue(" - Run pkgreflect on a GoLang directory."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			return
		}

		Cmd.State = PkgReflect(args...)
	},
}
