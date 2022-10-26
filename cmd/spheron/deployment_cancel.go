package spheron

import (
	"github.com/spf13/cobra"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var cancelDeploymentId string

var cancelDeploymentCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel a deployment",
	Long:  `Cancel a deployment`,
	Run: func(cmd *cobra.Command, args []string) {

		if(cancelDeploymentId == "") {
			cancelDeploymentId = SanitizeInput("Enter deployment ID to cancel: ")
		}

		deployment, err := spheron.PostCancelDeployment(cancelDeploymentId)

		if err != nil {
			panic(err)
		}

		if(deployment) {
			println("Deployment canceled successfully")
		} else {
			println("Deployment not canceled")
		}
	},
}

func init() {
	deploymentCmd.AddCommand(cancelDeploymentCmd)
	
	cancelDeploymentCmd.Flags().StringVarP(&cancelDeploymentId, "deployment", "d", "", "Deployment ID to cancel")
}