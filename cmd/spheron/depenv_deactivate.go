package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var deactivateDepEnvId string 

var deactivateDepEnvCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivate a deployment environment",
	Long:  `Deactivate a deployment environment`,
	Run: func(cmd *cobra.Command, args []string) {
		
		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		if(deactivateDepEnvId == "") {
			deactivateDepEnvId = SanitizeInput("Enter deployment environment ID to deactivate: ")
		}

		deactivatedEnv, err := spheron.DeactivateDeploymentEnvironmentVariable(projectId, deactivateDepEnvId)

		if err != nil {
			panic(err)
		}

		fmt.Println("Deployment Environment Status: ", deactivatedEnv.Status)
		
	},
}

func init() {
	depEnvCmd.AddCommand(deactivateDepEnvCmd)

	deactivateDepEnvCmd.Flags().StringVarP(&deactivateDepEnvId, "deployment", "d", "", "Deployment Environment ID")
}
