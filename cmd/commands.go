package cmd

import (
	"github.com/gearboxworks/buildtool/defaults"
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(ghrCmd)
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(selfUpdateCmd)
	rootCmd.AddCommand(golangCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(releaseCmd)
	rootCmd.AddCommand(pkgreflectCmd)
}


var ghrCmd = &cobra.Command{
	Use:   CmdGhr,
	Short: ux.SprintfBlue("Run github-release package."),
	Long:  ux.SprintfBlue("Run github-release package."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = GithubReleaser()
	},
}


var pushCmd = &cobra.Command{
	Use:   CmdPush,
	Short: ux.SprintfBlue("Push a gearboxworks repo to GitHub."),
	Long:  ux.SprintfBlue("Push a gearboxworks repo to GitHub."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = GitPush(args[0], args[1:]...)
	},
}


var versionCmd = &cobra.Command{
	Use:   loadTools.CmdVersion,
	Short: ux.SprintfBlue("Show version of %s.", defaults.BinaryName),
	Long:  ux.SprintfBlue("Show version of %s.", defaults.BinaryName),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = Version()
	},
}


var selfUpdateCmd = &cobra.Command{
	Use:   "selfupdate",
	Short: ux.SprintfBlue("Update version of %s.", defaults.BinaryName),
	Long:  ux.SprintfBlue("Check and update the latest version of %s.", defaults.BinaryName),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = VersionUpdate()
	},
}


var golangCmd = &cobra.Command{
	Use:   CmdGolang,
	Short: ux.SprintfBlue("Run various go commands."),
	Long:  ux.SprintfBlue("Run various go commands."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = Golang(args...)
	},
}


var getCmd = &cobra.Command{
	Use:   CmdGet,
	Short: ux.SprintfBlue("Get value from GoLang src."),
	Long:  ux.SprintfBlue("Get value from GoLang src code."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = Get(args...)
	},
}


var buildCmd = &cobra.Command{
	Use:   CmdBuild,
	Short: ux.SprintfBlue("Build a gearboxworks repo."),
	Long: ux.SprintfBlue("Build a gearboxworks repo."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = Build(args...)
	},
}


var releaseCmd = &cobra.Command{
	Use:   CmdRelease,
	Short: ux.SprintfBlue("Release a gearboxworks repo."),
	Long: ux.SprintfBlue("Release a gearboxworks repo, (public and private repos)."),
	DisableFlagParsing: false,
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = Release(args...)
	},
}


var pkgreflectCmd = &cobra.Command{
	Use:   CmdPkgReflect,
	Short: ux.SprintfBlue("Run pkgreflect on a directory."),
	Long:  ux.SprintfBlue("Run pkgreflect on a directory."),
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.State = PkgReflect(args...)
	},
}

