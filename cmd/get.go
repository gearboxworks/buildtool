package cmd

import (
	"github.com/newclarity/buildtool/cmd/helpers"
	"github.com/newclarity/buildtool/ux"
	"github.com/spf13/cobra"
	"strings"
)


func init() {
	rootCmd.AddCommand(getCmd)
}


var getCmd = &cobra.Command{
	Use:   helpers.CmdGet,
	Short: ux.SprintfBlue("Get value from GoLang src."),
	Long:  ux.SprintfBlue("Get value from GoLang src code."),
	Run:   Get,
}
func Get(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		if len(args) == 0 {
			ux.PrintflnYellow("Need to supply one of: 'all', 'name', 'version', 'srcrepo', 'src', 'binrepo', 'bin'")
			break
		}

		if args[0] == "all" {
			args = []string{"name", "version", "src", "bin"}
		}

		tmpl := helpers.NewArgTemplate(Cmd.Debug)

		Cmd.State = tmpl.ProcessArgs(cmd, args)
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		for _, v := range args {
			k := ""
			v = strings.ToLower(v)
			switch {
				case v == "name":
					k = "Name"
					v = helpers.BinaryName

				case v == "ver":
					fallthrough
				case v == "version":
					k = "Version"
					v = helpers.BinaryVersion

				case v == "src":
					fallthrough
				case v == "source":
					fallthrough
				case v == "srcrepo":
					k = "Source Repo"
					v = helpers.SourceRepo

				case v == "bin":
					fallthrough
				case v == "binary":
					fallthrough
				case v == "binrepo":
					k = "Binary Repo"
					v = helpers.BinaryRepo
			}

			v, Cmd.State = tmpl.GetValue(v, helpers.DefaultVersionFile...)
			if Cmd.State.IsNotOk() {
				continue
			}

			if len(args) > 1 {
				ux.PrintfBlue("%s: ", k)
			}
			ux.PrintflnCyan("%s", v)
		}
	}
}
