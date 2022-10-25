package spheron

import "github.com/spf13/cobra"

// Set is a command to set the project and deployment

var setCmd = &cobra.Command{
	Use:  "set",
	Short: "Sets the project",
	Long: `Allows you to fix a project ID for use throughout the session`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
