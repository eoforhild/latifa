package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "carbon",
	PreRun: func(cmd *cobra.Command, args []string) {
		//initConfig()
		//initLogging()
		//initDb()
	},
	Run: rootCmdRun,
}

func Execute() error {
	return rootCmd.Execute()
}

func rootCmdRun(cmd *cobra.Command, _ []string) {

}
