package cmd

const (
	SelectDefaultJson = "select:json"

	DefaultJsonString 		= "{}"

	DefaultJsonFile 		= "buildtool.json"
	DefaultJsonFileSuffix 	= ".json"

	CmdBuild 			= "build"
	CmdPush 			= "push"
	CmdVersion 			= "version"
	CmdRelease 			= "release"
	CmdGet	 			= "get"
	CmdGhr	 			= "ghr"
	CmdPkgReflect		= "pkgreflect"
	CmdGolang           = "go"

	FlagJsonFile     	= "json"

	FlagChdir       	= "chdir"
	FlagForce 			= "force"
	FlagRemoveTemplate	= "rm-tmpl"
	FlagRemoveOutput	= "rm-out"
	FlagDebug 			= "debug"
	FlagQuiet			= "quiet"

	FlagVersion 		= "version"
	FlagHelpFunctions	= "help-functions"
	FlagHelpVariables	= "help-variables"
	FlagHelpExamples	= "help-examples"
	FlagHelpAll			= "help-all"
)
