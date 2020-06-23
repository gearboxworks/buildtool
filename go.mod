module github.com/gearboxworks/buildtool

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../scribeHelpers/ux

replace github.com/newclarity/scribeHelpers/loadTools => ../scribeHelpers/loadTools

replace github.com/newclarity/scribeHelpers/toolCopy => ../scribeHelpers/toolCopy

replace github.com/newclarity/scribeHelpers/toolExec => ../scribeHelpers/toolExec

replace github.com/newclarity/scribeHelpers/toolGit => ../scribeHelpers/toolGit

replace github.com/newclarity/scribeHelpers/toolGitHub => ../scribeHelpers/toolGitHub

replace github.com/newclarity/scribeHelpers/toolPath => ../scribeHelpers/toolPath

replace github.com/newclarity/scribeHelpers/toolPrompt => ../scribeHelpers/toolPrompt

replace github.com/newclarity/scribeHelpers/toolService => ../scribeHelpers/toolService

replace github.com/newclarity/scribeHelpers/toolSystem => ../scribeHelpers/toolSystem

replace github.com/newclarity/scribeHelpers/toolTypes => ../scribeHelpers/toolTypes

replace github.com/newclarity/scribeHelpers/toolUx => ../scribeHelpers/toolUx

replace github.com/newclarity/scribeHelpers/toolRuntime => ../scribeHelpers/toolRuntime

replace github.com/newclarity/scribeHelpers/toolGoReleaser => ../scribeHelpers/toolGoReleaser

replace github.com/newclarity/scribeHelpers/toolSelfUpdate => ../scribeHelpers/toolSelfUpdate

replace github.com/newclarity/scribeHelpers/toolGo => ../scribeHelpers/toolGo

replace github.com/newclarity/scribeHelpers/toolGhr => ../scribeHelpers/toolGhr

replace github.com/newclarity/scribeHelpers/toolCobraHelp => ../scribeHelpers/toolCobraHelp

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/mitchellh/mapstructure v1.3.1 // indirect
	github.com/newclarity/scribeHelpers/loadTools v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolCobraHelp v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolExec v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/toolGhr v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolGit v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolGo v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolGoReleaser v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200621234507-ba6f08c6b68d
	github.com/newclarity/scribeHelpers/toolSelfUpdate v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/ux v0.0.0-20200621234507-ba6f08c6b68d
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.0
	gopkg.in/ini.v1 v1.57.0 // indirect
)
