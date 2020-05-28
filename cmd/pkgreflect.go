package cmd

import (
	"github.com/newclarity/buildtool/cmd/helpers"
	"github.com/newclarity/buildtool/ux"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(pkgreflectCmd)
}


var pkgreflectCmd = &cobra.Command{
	Use:   helpers.CmdPkgReflect,
	Short: ux.SprintfBlue("Run pkgreflect on a directory."),
	Long:  ux.SprintfBlue("Run pkgreflect on a directory."),
	Run:   RunPkgReflect,
}
func RunPkgReflect(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		//tmpl := helpers.NewArgTemplate(Cmd.Debug)
		//
		//Cmd.State = tmpl.ProcessArgs(cmd, args)
		//if Cmd.State.IsNotOk() {
		//	Cmd.State.PrintResponse()
		//	break
		//}
		//
		//var repo string
		//repo, Cmd.State = tmpl.GetBinaryRepo(helpers.DefaultVersionFile...)
		//if Cmd.State.IsNotOk() {
		//	break
		//}

		//args = []string{"ux"}	// DEBUG
		if len(args) == 0 {
			Cmd.State.SetError("No directory specified.")
			break
		}

		pkgreflect := helpers.NewPkgReflect(Cmd.Debug)
		if pkgreflect.State.IsNotOk() {
			break
		}
		pkgreflect.SetArgs()	//"--stdout")

		Cmd.State = pkgreflect.Run(args)
		if Cmd.State.IsNotOk() {
			//
		}
	}
}
