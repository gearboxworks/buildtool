//
package cmd

import (
	"fmt"
	"github.com/gearboxworks/buildtool/defaults"
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/toolCobraHelp"
	"github.com/newclarity/scribeHelpers/toolSelfUpdate"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const onlyOnce = "1"
var onlyTwice = []string{"", ""}


var CobraHelp *toolCobraHelp.TypeCommands
var CmdSelfUpdate *toolSelfUpdate.TypeSelfUpdate
var CmdScribe *loadTools.TypeScribeArgs
var ConfigFile string
const 	flagConfigFile  	= "config"


func init() {
	SetCmd()

	CobraHelp.AddCommands("GitHub", rootCmd, pushCmd, commitCmd, pullCmd, cloneCmd, syncCmd)
	CobraHelp.AddCommands("Workflow", rootCmd, buildCmd, releaseCmd, ghrCmd)
	CobraHelp.AddCommands("GoLang", rootCmd, golangCmd, pkgreflectCmd, vfsGenCmd, getCmd, setCmd)

	getCmd.AddCommand(getAllCmd)
	getCmd.AddCommand(getNameCmd)
	getCmd.AddCommand(getVersionCmd)
	getCmd.AddCommand(getSourceRepoCmd)
	getCmd.AddCommand(getBinaryRepoCmd)

	setCmd.AddCommand(setNameCmd)
	setCmd.AddCommand(setVersionCmd)
	setCmd.AddCommand(setSourceRepoCmd)
	setCmd.AddCommand(setBinaryRepoCmd)

	cobra.OnInitialize(initConfig)
	cobra.EnableCommandSorting = false

	rootCmd.PersistentFlags().StringVar(&ConfigFile, flagConfigFile, fmt.Sprintf("%s-config.json", defaults.BinaryName), ux.SprintfBlue("%s: config file.", defaults.BinaryName))
	_ = rootCmd.PersistentFlags().MarkHidden(flagConfigFile)

	//rootCmd.PersistentFlags().StringVarP(&CmdScribe.Json.File, loadTools.FlagJsonFile, "j", loadTools.DefaultJsonFile, ux.SprintfBlue("%s config file.", defaults.BinaryName))
	//rootCmd.PersistentFlags().StringVarP(&CmdScribe.WorkingPath.File, loadTools.FlagWorkingPath, "p", loadTools.DefaultWorkingPath, ux.SprintfBlue("Set working path."))
	//
	//rootCmd.PersistentFlags().BoolVarP(&CmdScribe.Chdir, loadTools.FlagChdir, "c", false, ux.SprintfBlue("Change to directory containing %s.", DefaultJsonFile))
	//rootCmd.PersistentFlags().BoolVarP(&CmdScribe.ForceOverwrite, loadTools.FlagForce, "f", false, ux.SprintfBlue("Force overwrite of output files."))
	////rootCmd.PersistentFlags().BoolVarP(&CmdScribe.RemoveOutput, loadTools.FlagRemoveOutput, "", false, ux.SprintfBlue("Remove output file afterwards."))
	//rootCmd.PersistentFlags().BoolVarP(&CmdScribe.QuietProgress, loadTools.FlagQuiet, "q", false, ux.SprintfBlue("Silence progress in shell scripts."))
	//
	//rootCmd.PersistentFlags().BoolVarP(&CmdScribe.Debug, loadTools.FlagDebug ,"d", false, ux.SprintfBlue("DEBUG mode."))
	//
	//rootCmd.PersistentFlags().BoolVarP(&CmdScribe.HelpAll, loadTools.FlagHelpAll, "", false, ux.SprintfBlue("Show all help."))
	////rootCmd.PersistentFlags().BoolVarP(&CmdScribe.HelpVariables, loadTools.FlagHelpVariables, "", false, ux.SprintfBlue("Help on template variables."))
	////rootCmd.PersistentFlags().BoolVarP(&CmdScribe.HelpFunctions, loadTools.FlagHelpFunctions, "", false, ux.SprintfBlue("Help on template functions."))
	////rootCmd.PersistentFlags().BoolVarP(&CmdScribe.HelpExamples, loadTools.FlagHelpExamples, "", false, ux.SprintfBlue("Help on template examples."))
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(ConfigFile)
	} else {
		// Find home directory.
		//home, err := homedir.Dir()
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}

		viper.AddConfigPath(".")
		viper.SetConfigName(defaults.BinaryName + "-config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}


func SetCmd() {
	for range onlyOnce {
		if CmdScribe == nil {
			CmdScribe = loadTools.New(defaults.BinaryName, defaults.BinaryVersion, false)
			CmdScribe.Runtime.SetRepos(defaults.SourceRepo, defaults.BinaryRepo)
			if CmdScribe.State.IsNotOk() {
				break
			}

			// Import additional tools.
			//CmdScribe.ImportTools(&buildtools.GetHelpers)
			//if CmdScribe.State.IsNotOk() {
			//	break
			//}

			CmdScribe.LoadCommands(rootCmd, true)
			if CmdScribe.State.IsNotOk() {
				break
			}

			CmdScribe.Template.Ignore()
			if CmdScribe.State.IsNotOk() {
				break
			}
		}

		if CobraHelp == nil {
			CobraHelp = toolCobraHelp.New(CmdScribe.Runtime)
			CobraHelp.SetHelp(rootCmd)
		}

		if CmdSelfUpdate == nil {
			CmdSelfUpdate = toolSelfUpdate.New(CmdScribe.Runtime)
			CmdSelfUpdate.LoadCommands(rootCmd, false)
			if CmdSelfUpdate.State.IsNotOk() {
				break
			}
		}
	}
}


var rootCmd = &cobra.Command{
	Use:   defaults.BinaryName,
	Short: "gearboxworks build tool.",
	Long: "gearboxworks build tool. Assists in building and releasing all gearboxworks related repos.",
	Run: gbRootFunc,
}

func gbRootFunc(cmd *cobra.Command, args []string) {
	for range onlyOnce {
		if CmdSelfUpdate.FlagCheckVersion(nil) {
			CmdScribe.State.SetOk()
			break
		}

		// Show help if no commands specified.
		if len(args) == 0 {
			_ = cmd.Help()
			CmdScribe.State.SetOk()
			break
		}
	}
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() *ux.State {
	for range onlyOnce {
		//SetHelp(rootCmd)
		//SetCmd()

		err := rootCmd.Execute()
		if err != nil {
			CmdScribe.State.SetError(err)
			break
		}

		CmdScribe.State = CheckReturns()
	}

	return CmdScribe.State
}


func CheckReturns() *ux.State {
	state := CmdScribe.State
	for range onlyOnce {
		if CmdScribe.State.IsNotOk() {
			state = CmdScribe.State
			break
		}

		if CmdSelfUpdate.State.IsNotOk() {
			state = CmdSelfUpdate.State
			break
		}

		if CobraHelp.State.IsNotOk() {
			state = CobraHelp.State
			break
		}
	}
	return state
}


//func ProcessArgs(toolArgs *loadTools.TypeScribeArgs, cmd *cobra.Command, args []string) *ux.State {
//	state := CmdScribe.State
//
//	for range onlyOnce {
//		err := toolArgs.Runtime.SetArgs(cmd.Use)
//		if err != nil {
//			state.SetError(err)
//			break
//		}
//
//		err = toolArgs.Runtime.AddArgs(args...)
//		if err != nil {
//			state.SetError(err)
//			break
//		}
//
//		for range onlyTwice {
//			if len(args) >= 1 {
//				ext := filepath.Ext(args[0])
//				if ext == ".json" {
//					toolArgs.Json.File = args[0]
//					args = args[1:]
//				} else if ext == ".tmpl" {
//					toolArgs.Template.File = args[0]
//					args = args[1:]
//				} else {
//					break
//				}
//			}
//		}
//		_ = CmdScribe.Runtime.SetArgs(args...)
//
//		toolArgs.Template.File = loadTools.SelectIgnore
//
//		state = toolArgs.ValidateArgs()
//		if state.IsNotOk() {
//			break
//		}
//	}
//
//	return state
//}
