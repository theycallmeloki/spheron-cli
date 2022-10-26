package spheron

import "github.com/spf13/cobra"

var createDeploymentCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a deployment",
	Long:  `Create a deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		// stub, replace with real code

		// gitUrl := spheron.SanitizeInput("Enter Git URL: ")
		// repoName := spheron.SanitizeInput("Enter Repo Name: ")
		// uniqueTopicId := spheron.SanitizeInput("Enter Unique Topic ID: ")
		// buildCommand := spheron.SanitizeInput("Enter Build Command: ")
		// installCommand := spheron.SanitizeInput("Enter Install Command: ")
		// workspace := spheron.SanitizeInput("Enter Workspace: ")
		// publishDir := spheron.SanitizeInput("Enter Publish Dir: ")
		// framework := spheron.SanitizeInput("Enter Framework: ")
		// nodeVersion := spheron.SanitizeInput("Enter Node Version: ")
		// protocol := spheron.SanitizeInput("Enter Protocol: ")
		// createDefaultWebhook := true
		// provider := spheron.SanitizeInput("Enter Provider: ")
		// branch := spheron.SanitizeInput("Enter Branch: ")
		// prComments := true
		// commitComments := true
		// buildStatus := true
		// githubDeployment := true

		// organizationID := viper.GetString("organization")

		// deployment := spheron.CreateDeploymentPayload{
		// 	OrganizationID: organizationID,
		// 	GitURL: gitUrl,
		// 	RepoName: repoName,
		// 	UniqueTopicID: uniqueTopicId,
		// 	Configuration: spheron.Configuration{
		// 		BuildCommand: buildCommand,
		// 		InstallCommand: installCommand,
		// 		Workspace: workspace,
		// 		PublishDir: publishDir,
		// 		Framework: framework,
		// 		NodeVersion: nodeVersion,
		// 	},
		// 	Protocol: protocol,
		// 	CreateDefaultWebhook: createDefaultWebhook,
		// 	Provider: provider,
		// 	Branch: branch,
		// 	GitProviderPreferences: spheron.GitProviderPreferences{
		// 		PrComments: prComments,
		// 		CommitComments: commitComments,
		// 		BuildStatus: buildStatus,
		// 		GithubDeployment: githubDeployment,
		// 	},
		// }

		// createdDeployment, err := spheron.PostDeployment(ProjectId, deployment)

		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Println("Deployment Success: ", createdDeployment.Success)
		// fmt.Println("Deployment Message: ", createdDeployment.Message)
		// fmt.Println("Deployment Topic: ", createdDeployment.Topic)
		// fmt.Println("Deployment ID: ", createdDeployment.DeploymentID)
		// fmt.Println("Project ID: ", createdDeployment.ProjectID)
		// fmt.Println("Deployment ID: ", createdDeployment.Body.DeploymentID)
		// fmt.Println("GithubURL: ", createdDeployment.Body.GithubURL)
		// fmt.Println("FolderName: ", createdDeployment.Body.FolderName)
		// fmt.Println("Topic: ", createdDeployment.Body.Topic)
		// fmt.Println("Framework: ", createdDeployment.Body.Framework)
		// fmt.Println("Branch: ", createdDeployment.Body.Branch)
		// fmt.Println("Install Command: ", createdDeployment.Body.InstallCommand)
		// fmt.Println("Build Command: ", createdDeployment.Body.BuildCommand)
		// fmt.Println("Workspace: ", createdDeployment.Body.Workspace)
		// fmt.Println("Publish Dir: ", createdDeployment.Body.PublishDirectory)
		// fmt.Println("Protocol: ", createdDeployment.Body.Protocol)
		// fmt.Println("IsWorkspace: ", createdDeployment.Body.IsWorkspace)
		// for _, log := range createdDeployment.Body.LogsToCapture {
		// 	fmt.Println("Logs: ", log.Key, log.Value)
		// }
		// fmt.Println("PaidViaSubscription: ", createdDeployment.Body.PaidViaSubscription)
		// fmt.Println("CommitId: ", createdDeployment.Body.CommitID)
	},
}

func init() {
	deploymentCmd.AddCommand(createDeploymentCmd)
}