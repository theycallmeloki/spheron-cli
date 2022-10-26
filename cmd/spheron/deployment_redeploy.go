package spheron

import (
	"github.com/spf13/cobra"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var redeployDeploymentId string

var redeployDeploymentCmd = &cobra.Command{
	Use:   "redeploy",
	Short: "Redeploy a deployment",
	Long:  `Redeploy a deployment`,
	Run: func(cmd *cobra.Command, args []string) {

		if(redeployDeploymentId == "") {
			redeployDeploymentId = SanitizeInput("Enter deployment ID to redeploy: ")
		}

		deployment, err := spheron.PostRedeployDeployment(redeployDeploymentId)

		if err != nil {
			panic(err)
		}

		if(deployment) {
			println("Deployment redeployed successfully")
		} else {
			println("Deployment not redeployed")
		}
	},
}

func init() {
	deploymentCmd.AddCommand(redeployDeploymentCmd)
	
	redeployDeploymentCmd.Flags().StringVarP(&redeployDeploymentId, "deployment", "d", "", "Deployment ID to redeploy")
}
