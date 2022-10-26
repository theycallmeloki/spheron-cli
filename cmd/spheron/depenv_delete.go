package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var deleteDepEnvId string

var deleteDepEnvCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a deployment environment",
	Long:  `Delete a deployment environment`,
	Run: func(cmd *cobra.Command, args []string) {
		
		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		if(deleteDepEnvId == "") {
			deleteDepEnvId = SanitizeInput("Enter deployment environment ID to delete: ")
		}

		deletedEnv, err := spheron.DeleteDeploymentEnvironmentVariable(projectId, deleteDepEnvId)

		if err != nil {
			panic(err)
		}

		if(deletedEnv){
			fmt.Println("Deleted Deployment Environment")
		} else {
			fmt.Println("Failed to delete deployment environment")
		}

		
	},
}

func init() {
	depEnvCmd.AddCommand(deleteDepEnvCmd)

	deleteDepEnvCmd.Flags().StringVarP(&deleteDepEnvId, "deployment", "d", "", "Deployment Environment ID")
}