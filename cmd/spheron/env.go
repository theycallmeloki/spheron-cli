package spheron

import "github.com/spf13/cobra"

// env command
var envCommand = &cobra.Command{
	Use:   "env",
	Short: "Manage environment variables",
	Long:  `Manage environment variables`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(envCommand)
}