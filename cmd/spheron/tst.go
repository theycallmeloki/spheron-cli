package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var EnvironmentId string
var DeploymentEnvironmentId string
var DomainId string
var DeploymentId string

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Dev CLI Locally",
	Long:  `Dev CLI Locally - This command is only available in test builds`,
	Run: func(cmd *cobra.Command, args []string) {

		// get project deployments and print them 
		fmt.Println("Getting Project Deployments\n")

		deployments, err := spheron.GetProjectDeployments(ProjectId)
		if err != nil {
			panic(err)
		}

		for _, deployment := range deployments {
			fmt.Println("Deployment ID: ", deployment.ID)
			fmt.Println("Deployment Status: ", deployment.Status)
			fmt.Println("Deployment Commit Message: ", deployment.CommitMessage)
		}

		fmt.Println("-----------------------------")

		/////////////////////////////////////////

		// get project deployment count and print them 
		fmt.Println("Getting Project Deployment Count\n")

		deploymentCount, err := spheron.GetProjectDeploymentsCount(ProjectId)
		if err != nil {
			panic(err)
		}

		fmt.Println("Deployment Total: ", deploymentCount.Total)
		fmt.Println("Deployment Successful: ", deploymentCount.Successful)
		fmt.Println("Deployment Failed: ", deploymentCount.Failed)
		fmt.Println("Deployment Pending: ", deploymentCount.Pending)

		fmt.Println("-----------------------------")

		/////////////////////////////////////////

		// generate an environment variable and push it to the project
		
		
		// get deployment environment variables and print them
		deploymentEnvironmentVariables, err := spheron.GetDeploymentEnvironmentVariables(ProjectId)
		if err != nil {
			panic(err)
		}

		var showingDepEnvs []string

		for _, env := range deploymentEnvironmentVariables {
			showingDepEnvs = append(showingDepEnvs, env.Name)
		}

		fmt.Println(showingDepEnvs)


		// selectedEnvName, err := spheron.SanitizeSelectMultiple(showingDepEnvs, "Select Deployment Environment: ", "Do you want to add another Deployment Environment?")
		selectedEnvName := []string{"Production", "Development"}

		if err != nil {
			panic(err)
		}

		fmt.Println(selectedEnvName)

		// postEnvVarKey := spheron.SanitizeInput("Enter Environment Variable Key: ")
		// postEnvVarValue := spheron.SanitizeInput("Enter Environment Variable Value: ")

		postEnvVarKey := "TEST_KEY"
		postEnvVarValue := "TEST_VALUE"

		constructedEnvVariable := spheron.EnvironmentVariables{
			Name:  postEnvVarKey,	
			Value: postEnvVarValue,
			DeploymentEnvironments: selectedEnvName,
		}
		envsToPush := []spheron.EnvironmentVariables{constructedEnvVariable}

		// push the environment variable to the project
		pushedEnvs, err := spheron.PostEnvironmentVariables(ProjectId, envsToPush)
		if err != nil {
			panic(err)
		}

		selectedEnvIDToDelete := pushedEnvs[0].ID

		for  _, env := range pushedEnvs {
			fmt.Println(env.Name, env.Value, env.ID)
		}

		/////////////////////////////////////////

		// update an existing environment variable and push it to the project

		fmt.Println("Putting shell selected environment variable")

		// putEnvVarKey := spheron.SanitizeInput("Enter Environment Variable Key: ")
		// putEnvVarValue := spheron.SanitizeInput("Enter Environment Variable Value: ")

		putEnvVarKey := "TEST_KEY"
		putEnvVarValue := "TEST_VALUE"

		constructedEnvVariable = spheron.EnvironmentVariables{
			Name:  putEnvVarKey,
			Value: putEnvVarValue,
			DeploymentEnvironments: selectedEnvName,
		}

		updatedEnv, err := spheron.PutEnvironmentVariable(ProjectId, EnvironmentId, constructedEnvVariable)
		if err != nil {
			panic(err)
		}

		fmt.Println(updatedEnv.Name, updatedEnv.Value, updatedEnv.ID)


		/////////////////////////////////////////

		// delete an existing environment variable

		fmt.Println("Deleting above environment variables")

		deletedEnv, err := spheron.DeleteEnvironmentVariable(ProjectId, selectedEnvIDToDelete)

		if err != nil {
			panic(err)
		}

		if(deletedEnv) {
			fmt.Println("Environment Variable Deleted")
		} else {
			fmt.Println("Environment Variable Not Deleted")
		}

		/////////////////////////////////////////

		// create a new deployment environment and push it to the project

		fmt.Println("Creating new deployment environment")

		// deploymentEnvironmentName := spheron.SanitizeInput("Enter Deployment Environment Name: ")
		deploymentEnvironmentName := "Staging"

		// listOfBranches, err := spheron.SanitizeInputMultiple("Enter Branch to Deploy: ", "Do you want to add another branch?")
		listOfBranches := []string{"main", "develop"}

		if err != nil {
			panic(err)
		}

		fmt.Println(listOfBranches)

		constructedDeploymentEnvironment := spheron.DeploymentEnvironment{
			Name: deploymentEnvironmentName,
			Branches: listOfBranches,
			Protocol: "arweave",
		}

		fmt.Println(constructedDeploymentEnvironment)

		postedDepEnv, err := spheron.PostDeploymentEnvironmentVariable(ProjectId, constructedDeploymentEnvironment)
		if err != nil {
			// panic(err)
			// ^ ignore for now, TODO: drop this when using it in cmd
		}

		fmt.Println("Name: " + postedDepEnv.Name + ", ID: " + postedDepEnv.ID)

		/////////////////////////////////////////

		// update an existing deployment environment and push it to the project

		fmt.Println("Updating deployment environment")

		deploymentEnvironmentName = "Staging"

		listOfBranches = []string{"develop"}

		constructedDeploymentEnvironment = spheron.DeploymentEnvironment{
			Name: deploymentEnvironmentName,
			Branches: listOfBranches,
			Protocol: "ipfs-pinata",
		}

		fmt.Println(constructedDeploymentEnvironment)

		updatedDepEnv, err := spheron.PutDeploymentEnvironmentVariable(ProjectId, DeploymentEnvironmentId, constructedDeploymentEnvironment)

		if err != nil {
			// panic(err)
			// ^ ignore for now, TODO: drop this when using it in cmd
		}

		fmt.Println("Name: " + updatedDepEnv.Name + ", ID: " + updatedDepEnv.ID)

		/////////////////////////////////////////

		// delete an existing deployment environment

		fmt.Println("Deleting deployment environment")

		deletedDepEnv, err := spheron.DeleteDeploymentEnvironmentVariable(ProjectId, DeploymentEnvironmentId)

		if err != nil {
			// panic(err)
			// ^ ignore for now, TODO: drop this when using it in cmd
		}

		if(deletedDepEnv) {
			fmt.Println("Deployment Environment Deleted")
		} else {
			fmt.Println("Deployment Environment Not Deleted")
		}

		/////////////////////////////////////////

		// deactivate an existing deployment environment

		fmt.Println("Deactivating deployment environment")

		deactivatedDepEnv, err := spheron.DeactivateDeploymentEnvironmentVariable(ProjectId, DeploymentEnvironmentId)

		if err != nil {
			panic(err)
		}

		fmt.Println(deactivatedDepEnv.Status)

		/////////////////////////////////////////

		// activate an existing deployment environment

		fmt.Println("Activating deployment environment")

		activatedDepEnv, err := spheron.ActivateDeploymentEnvironmentVariable(ProjectId, DeploymentEnvironmentId)

		if err != nil {
			panic(err)
		}

		fmt.Println(activatedDepEnv.Status)

		/////////////////////////////////////////

		// get all domains for a project

		fmt.Println("Getting all domains for a project")

		domains, err := spheron.GetDomains(ProjectId)

		if err != nil {
			panic(err)
		}

		for _, domain := range domains {
			fmt.Println("Domain ID: ", domain.ID)
			fmt.Println("Domain Name: ", domain.Name)
			fmt.Println("Domain Link: ", domain.Link)
			fmt.Println("Domain Is Latest: ", domain.IsLatest)
			fmt.Println("Domain Is Verified: ", domain.Verified)
			fmt.Println("Domain Type: ", domain.Type)
			fmt.Println("Domain Deployment Environment IDs: ")
			for _, env := range domain.DeploymentEnvironmentIds {
				fmt.Println("Domain Environment: ", env)
			}
			fmt.Println("Domain Version: ", domain.Version)
		}

		/////////////////////////////////////////

		// create a domain for a project

		fmt.Println("Creating a domain for a project")

		// link := spheron.SanitizeInput("Enter Domain Link: ")
		// name := spheron.SanitizeInput("Enter Domain Name: ")
		
		// DomainItems := []string{"subdomain", "domain"}
		// domainType, err := spheron.SanitizeFixedSelect(DomainItems, "Select Domain Type")
		// if err != nil {
		// 	panic(err)
		// }

		domainType := "subdomain"

		domain := spheron.CreateDomainPayload{
			// Link: link,
			Link: "somethingsomething.spheron.app",
			Type: domainType,
			DeploymentEnvironments: []string{"Production", "Development"},
			IsLatest: true,
			// Name: name,
			Name: "somethingsomething",
		}

		createdDomain, err := spheron.PostDomain(ProjectId, domain)

		if err != nil {
			//panic(err)
			// ^ ignore for now, TODO: drop this when using it in cmd
		}

		fmt.Println("Domain ID: ", createdDomain.ID)
		fmt.Println("Domain Name: ", createdDomain.Name)
		fmt.Println("Domain Link: ", createdDomain.Link)
		fmt.Println("Domain Is Latest: ", createdDomain.IsLatest)
		fmt.Println("Domain Is Verified: ", createdDomain.Verified)
		fmt.Println("Domain Type: ", createdDomain.Type)
		fmt.Println("Domain Deployment Environment IDs: ")
		for _, env := range createdDomain.DeploymentEnvironmentIds {
			fmt.Println("Domain Environment: ", env)
		}
		fmt.Println("Domain Version: ", createdDomain.Version)

		/////////////////////////////////////////

		// get a domain with domainId for a project

		fmt.Println("Getting a domain with domainId for a project")

		getDomain, err := spheron.GetDomain(ProjectId, DomainId)

		if err != nil {
			panic(err)
		}

		fmt.Println("Domain ID: ", getDomain.ID)
		fmt.Println("Domain Name: ", getDomain.Name)
		fmt.Println("Domain Link: ", getDomain.Link)
		fmt.Println("Domain Is Latest: ", getDomain.IsLatest)
		fmt.Println("Domain Is Verified: ", getDomain.Verified)
		fmt.Println("Domain Type: ", getDomain.Type)
		fmt.Println("Domain Deployment Environment IDs: ")
		for _, env := range getDomain.DeploymentEnvironmentIds {
			fmt.Println("Domain Environment: ", env)
		}
		fmt.Println("Domain Version: ", getDomain.Version)

		/////////////////////////////////////////

		// update a domain with domainId for a project

		fmt.Println("Updating a domain with domainId for a project")

		// link := spheron.SanitizeInput("Enter Domain Link: ")
		// name := spheron.SanitizeInput("Enter Domain Name: ")

		link := "somethingsomething.spheron.app"
		name := "somethingsomething"

		deploymentEnvironments := []string{"Production", "Development"}
		isLatest := true

		domainUpdate := spheron.UpdateDomainPayload{
			Link: link,
			Name: name,
			DeploymentEnvironments: deploymentEnvironments,
			IsLatest: isLatest,
		}

		updatedDomain, err := spheron.PatchDomain(ProjectId, DomainId, domainUpdate)

		if err != nil {
			// panic(err)
			// ^ ignore for now, TODO: drop this when using it in cmd
		}

		fmt.Println("Domain ID: ", updatedDomain.ID)
		fmt.Println("Domain Name: ", updatedDomain.Name)
		fmt.Println("Domain Link: ", updatedDomain.Link)
		fmt.Println("Domain Is Latest: ", updatedDomain.IsLatest)
		fmt.Println("Domain Is Verified: ", updatedDomain.Verified)
		fmt.Println("Domain Type: ", updatedDomain.Type)
		fmt.Println("Domain Deployment Environment IDs: ")
		for _, env := range updatedDomain.DeploymentEnvironmentIds {
			fmt.Println("Domain Environment: ", env)
		}
		fmt.Println("Domain Version: ", updatedDomain.Version)

		/////////////////////////////////////////

		// delete a domain with domainId for a project

		// fmt.Println("Deleting a domain with domainId for a project")

		// deletedDomain, err := spheron.DeleteDomain(ProjectId, DomainId)
		// ^ Dont actually delete the domain

		// if err != nil {
		// 	panic(err)
		// }

		// if(deletedDomain) {
		// 	fmt.Println("Domain Deleted")
		// }


		/////////////////////////////////////////

		// verify a domain with domainId for a project

		fmt.Println("Verifying a domain with domainId for a project")

		verifyDomain, err := spheron.VerifyDomain(ProjectId, DomainId)

		if err != nil {
			// panic(err)
			// ^ ignore for now, TODO: drop this when using it in cmd
		}

		if(verifyDomain) {
			fmt.Println("Domain Verified")
		}

		/////////////////////////////////////////

		// create a deployment (Not working, possibly need more internal system context wrt CI/CD of repo to get this to work)

		// fmt.Println("Creating a deployment")

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

		/////////////////////////////////////////

		// get a deployment with deploymentId for a project

		fmt.Println("Getting a deployment with deploymentId")

		deployment, err := spheron.GetDeployment(DeploymentId)

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
		fmt.Println("Project: ", deployment.Project)
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

		/////////////////////////////////////////

		// authorize deployment with deploymentId

		fmt.Println("Authorizing a deployment with deploymentId")

		authorizedDeployment, err := spheron.PostAuthorizeDeployment(DeploymentId)

		if err != nil {
			panic(err)
		}

		fmt.Println("Deployment Authorized: ", authorizedDeployment)

		/////////////////////////////////////////

		// cancel a deployment with deploymentId

		fmt.Println("Canceling a deployment with deploymentId")

		cancelDeployment, err := spheron.PostCancelDeployment(DeploymentId)

		if err != nil {
			panic(err)
		}

		fmt.Println("Deployment Canceled: ", cancelDeployment)

		/////////////////////////////////////////

		// redploy a deployment with deploymentId

		fmt.Println("Redeploying a deployment with deploymentId")

		redeployDeployment, err := spheron.PostRedeployDeployment(DeploymentId)

		if err != nil {
			panic(err)
		}

		fmt.Println("Deployment Redeployed: ", redeployDeployment)

		/////////////////////////////////////////

	},
}

func init(){
	if(version == "TEST_BUILD") {
		rootCmd.AddCommand(testCmd)
	}

	testCmd.PersistentFlags().StringVarP(&ProjectId, "project", "p", "", "Project ID")
	// testCmd.PersistentFlags().StringVarP(&EnvironmentId, "environment", "e", "", "Environment ID")
	testCmd.PersistentFlags().StringVarP(&DeploymentEnvironmentId, "deploymentEnvironment", "d", "", "Deployment Environment ID")
	testCmd.PersistentFlags().StringVarP(&DomainId, "domain", "o", "", "Domain ID")
	testCmd.PersistentFlags().StringVarP(&DeploymentId, "deployment", "l", "", "Deployment ID")
}
