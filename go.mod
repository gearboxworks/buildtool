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

require (
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/github-release/github-release v0.8.1
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/go-github v17.0.0+incompatible
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c // indirect
	github.com/inconshreveable/log15 v0.0.0-20200109203555-b30bc20e4fd1 // indirect
	github.com/kevinburke/rest v0.0.0-20200429221318-0d2892b400f8
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db
	github.com/mitchellh/mapstructure v1.3.1 // indirect
	github.com/newclarity/scribeHelpers/loadTools v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolExec v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/toolGit v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolGoReleaser v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/toolSelfUpdate v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/ux v0.0.0-20200604000029-dbb313f0fedc
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/rhysd/go-github-selfupdate v1.2.2 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.0
	github.com/tcnksm/go-gitconfig v0.1.2
	github.com/tcnksm/go-latest v0.0.0-20170313132115-e3007ae9052e
	github.com/tj/go-prompt v1.3.0
	github.com/tomnomnom/linkheader v0.0.0-20180905144013-02ca5825eb80
	github.com/tsuyoshiwada/go-gitcmd v0.0.0-20180205145712-5f1f5f9475df
	github.com/ulikunitz/xz v0.5.7 // indirect
	github.com/voxelbrain/goptions v0.0.0-20180630082107-58cddc247ea2
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/ini.v1 v1.57.0 // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1
)
