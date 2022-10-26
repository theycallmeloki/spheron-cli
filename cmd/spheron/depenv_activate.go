package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var activateDepEnvId string

var activateDepEnvCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activate a deployment environment",
	Long:  `Activate a deployment environment`,
	Run: func(cmd *cobra.Command, args []string) {
		
		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		if(activateDepEnvId == "") {
			activateDepEnvId = SanitizeInput("Enter deployment environment ID to activate: ")
		}

		activatedEnv, err := spheron.ActivateDeploymentEnvironmentVariable(projectId, activateDepEnvId)

		if err != nil {
			panic(err)
		}

		fmt.Println("Deployment Environment Status: ", activatedEnv.Status)
		
	},
}

func init() {
	depEnvCmd.AddCommand(activateDepEnvCmd)

	activateDepEnvCmd.Flags().StringVarP(&activateDepEnvId, "deployment", "d", "", "Deployment Environment ID")
}