package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var viewDeploymentId string

var viewDeploymentCmd = &cobra.Command{
	Use:   "view",
	Short: "View a deployment",
	Long:  `View a deployment`,
	Run: func(cmd *cobra.Command, args []string) {

		if(viewDeploymentId == "") {
			viewDeploymentId = SanitizeInput("Enter deployment ID to view: ")
		}

		deployment, err := spheron.GetDeployment(viewDeploymentId)

		if err != nil {
			panic(err)
		}

		fmt.Println("Deployment ID: ", deployment.ID)
		fmt.Println("Site Preview: ", deployment.SitePreview)
		fmt.Println("Commit ID: ", deployment.CommitID)
		fmt.Println("Commit Message: ", deployment.CommitMessage)
		fmt.Println("ContentHash: ", deployment.ContentHash)
		for _, bd := range deployment.BuildDirectory {
			fmt.Println("Build Directory: ", bd)
		}
		fmt.Println("Topic: ", deployment.Topic)
		fmt.Println("Status: ", deployment.Status)
		fmt.Println("Payment ID: ", deployment.PaymentID)
		fmt.Println("Build Time: ", deployment.BuildTime)
		fmt.Println("Memory Used: ", deployment.MemoryUsed)
		fmt.Println("-----------Project-----------")
		fmt.Println("ID: ", deployment.Project.ID)
		fmt.Println("Name: ", deployment.Project.Name)
		fmt.Println("URL: ", deployment.Project.URL)
		// TODO: use a marshaller to retrieve it as a string
		// fmt.Println("Latest Deployment: ", deployment.Project.LatestDeployment)
		fmt.Println("----------/Project-----------")
		fmt.Println("Screenshot Fee: ", deployment.Screenshot.Fee)
		fmt.Println("Screenshot URL: ", deployment.Screenshot.URL)
		fmt.Println("Deployment Initiator: ", deployment.DeploymentInitiator)
		fmt.Println("Branch: ", deployment.Branch)
		fmt.Println("ExternalRespositoryName: ", deployment.ExternalRepositoryName)
		fmt.Println("Protocol: ", deployment.Protocol)
		fmt.Println("DeploymentEnvironmentName: ", deployment.DeploymentEnvironmentName)
		fmt.Println("Failed Message: ", deployment.FailedMessage)
		fmt.Println("Is From Request: ", deployment.IsFromRequest)
		fmt.Println("Configuration - Build Command: ", deployment.Configuration.BuildCommand)
		fmt.Println("Configuration - Install Command: ", deployment.Configuration.InstallCommand)
		fmt.Println("Configuration - Workspace: ", deployment.Configuration.Workspace)
		fmt.Println("Configuration - Publish Directory: ", deployment.Configuration.PublishDir)
		fmt.Println("Configuration - Framework: ", deployment.Configuration.Framework)
		fmt.Println("Configuration - Node Version: ", deployment.Configuration.NodeVersion)

	},
}

func init() {
	deploymentCmd.AddCommand(viewDeploymentCmd)
	
	viewDeploymentCmd.Flags().StringVarP(&viewDeploymentId, "deployment", "d", "", "Deployment ID to view")
}