package spheron

import (
	"github.com/spf13/cobra"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var authorizeDeploymentId string

var authorizeDeploymentCmd = &cobra.Command{
	Use:   "authorize",
	Short: "Authorize a deployment",
	Long:  `Authorize a deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		deployment, err := spheron.PostAuthorizeDeployment(authorizeDeploymentId)

		if err != nil {
			panic(err)
		}

		if(deployment) {
			println("Deployment authorized successfully")
		} else {
			println("Deployment not authorized")
		}
	},
}

func init() {
	deploymentCmd.AddCommand(authorizeDeploymentCmd)
	
	authorizeDeploymentCmd.Flags().StringVarP(&authorizeDeploymentId, "deployment", "d", "", "Deployment ID to authorize")
}

