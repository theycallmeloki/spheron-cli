package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var deleteEnvironmentId string

// EnvDeleteCmd is a command to delete an environment variable

var envDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an environment variable",
	Long:  `Delete an environment variable`,
	Run: func(cmd *cobra.Command, args []string) {

		projectId := viper.GetString("project")

		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		// fmt.Println("Deleting environment variable with ID: " + deleteEnvironmentId)

		if(deleteEnvironmentId == "") {
			deleteEnvironmentId = SanitizeInput("Enter environment ID to delete: ")
		}

		deletedEnv, err := spheron.DeleteEnvironmentVariable(projectId, deleteEnvironmentId)

		if err != nil {
			panic(err)
		}

		if(deletedEnv){
			fmt.Println("Deleted environment variable")
		} else {
			fmt.Println("Failed to delete environment variable")
		}
	},
}

func init() {
	envCommand.AddCommand(envDeleteCmd)

	envDeleteCmd.Flags().StringVarP(&deleteEnvironmentId, "environment", "e", "", "Environment ID")
}