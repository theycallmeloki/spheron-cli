package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var projectDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "List deployments of your project",
	Long:  `List deployments of your project`,
	Run: func(cmd *cobra.Command, args []string) {
		if(ProjectId == "") {
			ProjectId = viper.GetString("project")
			if(ProjectId == "") {
				ProjectId = SanitizeInput("Enter project ID: ")
			}
		}

		deployments, err := spheron.GetProjectDeployments(ProjectId)

		if err != nil {
			panic(err)
		}

		deploymentCount, err := spheron.GetProjectDeploymentsCount(ProjectId)
		if err != nil {
			panic(err)
		}
		
		fmt.Println("Your project has been deployed " + string(deploymentCount.Total) + " times\n")
		fmt.Println("Successful deployments: ", deploymentCount.Successful)
		fmt.Println("Failed deployments: ", deploymentCount.Failed)
		fmt.Println("Pending deployments: ", deploymentCount.Pending)

		for _, deployment := range deployments {
			fmt.Println("--------------Deployment--------------")
			fmt.Println("Deployment ID: ", deployment.ID)
			fmt.Println("Deployment Status: ", deployment.Status)
			fmt.Println("Deployment Commit Message: ", deployment.CommitMessage)
			fmt.Println("-------------/Deployment--------------")
		}
	},
}

func init() {
	projectCmd.AddCommand(projectDeploymentCmd)

	projectDeploymentCmd.Flags().StringVarP(&ProjectId, "project", "p", "", "Project ID")
}
