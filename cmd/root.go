package cmd

import (
	"fmt"
	"github.com/newclarity/buildtool/defaults"
	"github.com/newclarity/buildtool/ux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


//var Cmd TypeCmd
var Cmd *ArgTemplate
var ConfigFile string
const 	flagConfigFile  	= "config"


func init() {
	Cmd = NewArgTemplate(false)

	cobra.OnInitialize(initConfig)

	//rootCmd.PersistentFlags().StringVar(&ConfigFile, flagConfigFile, fmt.Sprintf("%s-config.json", defaults.BinaryName), ux.SprintfBlue("%s: config file.", defaults.BinaryName))
	_ = rootCmd.PersistentFlags().MarkHidden(flagConfigFile)

	rootCmd.PersistentFlags().StringVarP(&Cmd.Json.Path, FlagJsonFile, "j", DefaultJsonFile, ux.SprintfBlue("%s config file.", defaults.BinaryName))

	rootCmd.PersistentFlags().BoolVarP(&Cmd.Chdir, FlagChdir, "c", false, ux.SprintfBlue("Change to directory containing %s", DefaultJsonFile))
	rootCmd.PersistentFlags().BoolVarP(&Cmd.ForceOverwrite, FlagForce, "f", false, ux.SprintfBlue("Force overwrite of output files."))
	//rootCmd.PersistentFlags().BoolVarP(&Cmd.RemoveOutput, FlagRemoveOutput, "", false, ux.SprintfBlue("Remove output file afterwards."))
	rootCmd.PersistentFlags().BoolVarP(&Cmd.QuietProgress, FlagQuiet, "q", false, ux.SprintfBlue("Silence progress in shell scripts."))

	rootCmd.PersistentFlags().BoolVarP(&Cmd.Debug, FlagDebug ,"d", false, ux.SprintfBlue("DEBUG mode."))

	rootCmd.PersistentFlags().BoolVarP(&Cmd.HelpAll, FlagHelpAll, "", false, ux.SprintfBlue("Show all help."))
	//rootCmd.PersistentFlags().BoolVarP(&Cmd.HelpVariables, FlagHelpVariables, "", false, ux.SprintfBlue("Help on template variables."))
	//rootCmd.PersistentFlags().BoolVarP(&Cmd.HelpFunctions, FlagHelpFunctions, "", false, ux.SprintfBlue("Help on template functions."))
	//rootCmd.PersistentFlags().BoolVarP(&Cmd.HelpExamples, FlagHelpExamples, "", false, ux.SprintfBlue("Help on template examples."))

	rootCmd.Flags().BoolP(FlagVersion, "v", false, ux.SprintfBlue("Display version of " + defaults.BinaryName))
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


// rootCmd represents the base command when called without any sub-commands
var rootCmd = &cobra.Command{
	Use:   defaults.BinaryName,
	Short: "gearboxworks build tool.",
	Long: "gearboxworks build tool. Assists in building and releasing all gearboxworks related repos.",
	Run: gbRootFunc,
}

func gbRootFunc(cmd *cobra.Command, args []string) {
	Cmd.State = Cmd.State.EnsureNotNil()

	for range OnlyOnce {
		var ok bool
		fl := cmd.Flags()
		tmpl := NewArgTemplate(false)

		// ////////////////////////////////
		// Show version.
		ok, _ = fl.GetBool(FlagVersion)
		if ok {
			Version(cmd, args)
			Cmd.State.Clear()
			break
		}

		tmpl = ProcessArgs(cmd, args)
		Cmd.State = tmpl.State
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			_ = cmd.Help()
			break
		}

		// Show help if no commands specified.
		if len(args) == 0 {
			Version(cmd, args)
			_ = cmd.Help()
			Cmd.State.Clear()
			break
		}
	}
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() *ux.State {
	for range OnlyOnce {
		var err error

		SetHelp(rootCmd)

		if Cmd == nil {
			Cmd = NewArgTemplate(false)
		}

		err = rootCmd.Execute()
		if err != nil {
			Cmd.State.SetError(err)
			break
		}
	}

	return Cmd.State
}

func _GetUsage(c *cobra.Command) string {
	var str string

	if c.Parent() == nil {
		str += ux.SprintfCyan("%s [flags] ", c.Name())
	} else {
		str += ux.SprintfCyan("%s [flags] ", c.Parent().Name())
		str += ux.SprintfGreen("%s ", c.Use)
	}

	if c.HasAvailableSubCommands() {
		str += ux.SprintfGreen("[command] ")
		str += ux.SprintfCyan("<args> ")
	}

	return str
}

func _GetVersion(c *cobra.Command) string {
	var str string

	if c.Parent() == nil {
		str += ux.SprintfWhite("%s: v%s", defaults.BinaryName, defaults.BinaryVersion)
	}

	return str
}

func SetHelp(c *cobra.Command) {
	var tmplHelp string
	var tmplUsage string

	//fmt.Printf("%s", rootCmd.UsageTemplate())
	//fmt.Printf("%s", rootCmd.HelpTemplate())

	cobra.AddTemplateFunc("GetUsage", _GetUsage)
	cobra.AddTemplateFunc("GetVersion", _GetVersion)

	cobra.AddTemplateFunc("SprintfBlue", ux.SprintfBlue)
	cobra.AddTemplateFunc("SprintfCyan", ux.SprintfCyan)
	cobra.AddTemplateFunc("SprintfGreen", ux.SprintfGreen)
	cobra.AddTemplateFunc("SprintfMagenta", ux.SprintfMagenta)
	cobra.AddTemplateFunc("SprintfRed", ux.SprintfRed)
	cobra.AddTemplateFunc("SprintfWhite", ux.SprintfWhite)
	cobra.AddTemplateFunc("SprintfYellow", ux.SprintfYellow)

	// 	{{ with .Parent }}{{ SprintfCyan .Name }}{{ end }} {{ SprintfGreen .Name }} {{ if .HasAvailableSubCommands }}{{ SprintfGreen "[command]" }}{{ end }}

	tmplUsage += `
{{ SprintfBlue "Usage: " }}
	{{ GetUsage . }}

{{- if gt (len .Aliases) 0 }}
{{ SprintfBlue "\nAliases:" }} {{ .NameAndAliases }}
{{- end }}

{{- if .HasExample }}
{{ SprintfBlue "\nExamples:" }}
	{{ .Example }}
{{- end }}

{{- if .HasAvailableSubCommands }}
{{ SprintfBlue "\nWhere " }}{{ SprintfGreen "[command]" }}{{ SprintfBlue " is one of:" }}
{{- range .Commands }}
{{- if (or .IsAvailableCommand (eq .Name "help")) }}
	{{ rpad (SprintfGreen .Name) .NamePadding}}	- {{ .Short }}{{ end }}
{{- end }}
{{- end }}

{{- if .HasAvailableLocalFlags }}
{{ SprintfBlue "\nFlags:" }}
{{ .LocalFlags.FlagUsages | trimTrailingWhitespaces }}
{{- end }}

{{- if .HasAvailableInheritedFlags }}
{{ SprintfBlue "\nGlobal Flags:" }}
{{ .InheritedFlags.FlagUsages | trimTrailingWhitespaces }}
{{- end }}

{{- if .HasHelpSubCommands }}
{{- SprintfBlue "\nAdditional help topics:" }}
{{- range .Commands }}
{{- if .IsAdditionalHelpTopicCommand }}
	{{ rpad (SprintfGreen .CommandPath) .CommandPathPadding }} {{ .Short }}
{{- end }}
{{- end }}
{{- end }}

{{- if .HasAvailableSubCommands }}
{{ SprintfBlue "\nUse" }} {{ SprintfCyan .CommandPath }} {{ SprintfCyan "help" }} {{ SprintfGreen "[command]" }} {{ SprintfBlue "for more information about a command." }}
{{- end }}
`

	tmplHelp = `{{ GetVersion . }}

{{ SprintfBlue "Commmand:" }} {{ SprintfCyan .Use }}

{{ SprintfBlue "Description:" }} 
	{{ with (or .Long .Short) }}
{{- . | trimTrailingWhitespaces }}
{{- end }}

{{- if or .Runnable .HasSubCommands }}
{{ .UsageString }}
{{- end }}
`

	//c.SetHelpCommand(c)
	//c.SetHelpFunc(PrintHelp)
	c.SetHelpTemplate(tmplHelp)
	c.SetUsageTemplate(tmplUsage)
}
