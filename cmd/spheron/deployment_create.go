package spheron

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var createDeploymentCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a deployment",
	Long:  `Create a deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		// stub, replace with real code

		gitUrl := SanitizeInput("Enter Git URL: ")
		repoName := SanitizeInput("Enter Repo Name: ")
		uniqueTopicId := uuid.New().String()
		fmt.Println("Using a Unique Topic ID of ", uniqueTopicId)
		buildCommand := SanitizeInput("Enter Build Command: ")
		installCommand := SanitizeInput("Enter Install Command: ")
		workspace := SanitizeInput("Enter Workspace (Also called Root Directory): ")
		publishDir := SanitizeInput("Enter Publish Dir: ")
		frameworkSelection, err := SanitizeFixedSelect(frameworkList, "Select Framework: ")
		if err != nil {
			panic(err)
		}
		framework := GetFrameworkEnum(frameworkSelection)
		nodeSelection, err := SanitizeFixedSelect(nodeList, "Select Node Version: ")
		if err != nil {
			panic(err)
		}
		nodeVersion := GetNodeEnum(nodeSelection)
		protocolSelection, err := SanitizeFixedSelect(protocolList, "Select Protocol: ")
		if err != nil {
			panic(err)
		}
		protocol := GetProtocolEnum(protocolSelection)

		askForCreationOfWebhook, err := SanitizeFixedSelect([]string{"Yes", "No"}, "Create Webhook? ")

		if err != nil {
			panic(err)
		}
		var createDefaultWebhook bool
		if(askForCreationOfWebhook == "Yes") {
			createDefaultWebhook = true
		} else {
			createDefaultWebhook = false
		}

		// provider := SanitizeInput("Enter Provider: ")
		providerSelection, err := SanitizeFixedSelect(providerList, "Select Provider: ")
		if err != nil {
			panic(err)
		}
		provider := GetProviderEnum(providerSelection)
		branch := SanitizeInput("Enter Branch: ")
		// TODO: Do we need to do anything about this? 
		prComments := true
		commitComments := true
		buildStatus := true
		githubDeployment := true

		organizationID := viper.GetString("organization")

		deployment := spheron.CreateDeploymentPayload{
			OrganizationID: organizationID,
			GitURL: gitUrl,
			RepoName: repoName,
			UniqueTopicID: uniqueTopicId,
			Configuration: spheron.Configuration{
				BuildCommand: buildCommand,
				InstallCommand: installCommand,
				Workspace: workspace,
				PublishDir: publishDir,
				Framework: framework,
				NodeVersion: nodeVersion,
			},
			Protocol: protocol,
			CreateDefaultWebhook: createDefaultWebhook,
			Provider: provider,
			Branch: branch,
			GitProviderPreferences: spheron.GitProviderPreferences{
				PrComments: prComments,
				CommitComments: commitComments,
				BuildStatus: buildStatus,
				GithubDeployment: githubDeployment,
			},
		}

		createdDeployment, err := spheron.PostDeployment(deployment)

		if err != nil {
			panic(err)
		}

		fmt.Println("Deployment Success: ", createdDeployment.Success)
		fmt.Println("Deployment Message: ", createdDeployment.Message)
		fmt.Println("Deployment Topic: ", createdDeployment.Topic)
		fmt.Println("Deployment ID: ", createdDeployment.DeploymentID)
		fmt.Println("Project ID: ", createdDeployment.ProjectID)
		fmt.Println("Deployment ID: ", createdDeployment.Body.DeploymentID)
		fmt.Println("GithubURL: ", createdDeployment.Body.GithubURL)
		fmt.Println("FolderName: ", createdDeployment.Body.FolderName)
		fmt.Println("Topic: ", createdDeployment.Body.Topic)
		fmt.Println("Framework: ", createdDeployment.Body.Framework)
		fmt.Println("Branch: ", createdDeployment.Body.Branch)
		fmt.Println("Install Command: ", createdDeployment.Body.InstallCommand)
		fmt.Println("Build Command: ", createdDeployment.Body.BuildCommand)
		fmt.Println("Workspace: ", createdDeployment.Body.Workspace)
		fmt.Println("Publish Dir: ", createdDeployment.Body.PublishDirectory)
		fmt.Println("Protocol: ", createdDeployment.Body.Protocol)
		fmt.Println("IsWorkspace: ", createdDeployment.Body.IsWorkspace)
		for _, log := range createdDeployment.Body.LogsToCapture {
			fmt.Println("Logs: ", log.Key, log.Value)
		}
		fmt.Println("PaidViaSubscription: ", createdDeployment.Body.PaidViaSubscription)
		fmt.Println("CommitId: ", createdDeployment.Body.CommitID)
	},
}

func init() {
	deploymentCmd.AddCommand(createDeploymentCmd)
}