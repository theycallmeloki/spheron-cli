package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var ProjectId string

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long:  `Manage projects`,
	Run: func(cmd *cobra.Command, args []string) {
		
		if(ProjectId == "") {
			ProjectId = viper.GetString("project")
			if(ProjectId == "") {
				ProjectId = SanitizeInput("Enter project ID: ")
			}
		}
		
		// fetch the project with the provided project ID and print them
		project, err := spheron.GetProject(ProjectId)
		if err != nil {
			panic(err)
		}
		fmt.Println("Project ID: ", project.ID)
		fmt.Println("Project Name: ", project.Name)
		fmt.Println("Project URL: ", project.URL)
		fmt.Println("Project Created At: ", project.CreatedAt)
		fmt.Println("State: ", project.State)
		fmt.Println("Organization", project.Organization)
		fmt.Println("Provider: ", project.Provider)
		for _, deploymentEnvironment := range project.DeploymentEnvironments {
			fmt.Println("-----------Deployment Environment--- " + deploymentEnvironment.Name + " --------")
			fmt.Println("Deployment Environment ID: ", deploymentEnvironment.ID)
			fmt.Println("------Branches------")
			for _, branch := range deploymentEnvironment.Branches {
				fmt.Println(branch)
			}
			fmt.Println("-----/Branches------")
			fmt.Println("Deployment Environment Name: ", deploymentEnvironment.Name)
			fmt.Println("Deployment Environment Status: ", deploymentEnvironment.Status)
			fmt.Println("Deployment Environment Protocol: ", deploymentEnvironment.Protocol)
			fmt.Println("----------/Deployment Environment--- " + deploymentEnvironment.Name + " --------")
		}
		fmt.Println("\nLatest Deployment\n")
		fmt.Println("Deployment ID: ", project.LatestDeployment.ID)
		fmt.Println("Deployment Commit Message: ", project.LatestDeployment.CommitMessage)
		fmt.Println("Deployment Commit Status: ", project.LatestDeployment.Status)
		fmt.Println("Deployment Build Time: ", project.LatestDeployment.BuildTime)
		fmt.Println("Deployment Environment Name: ", project.LatestDeployment.DeploymentEnvironmentName)
		fmt.Println("Deployment Commit ID: ", project.LatestDeployment.CommitID)
		fmt.Println("Deployment Branch: ", project.LatestDeployment.Branch)
		fmt.Println("Configuration Node Version: ", project.LatestDeployment.Configuration.NodeVersion)
		fmt.Println("Configuration Install Command: ", project.LatestDeployment.Configuration.InstallCommand)
		fmt.Println("Configuration Build Command: ", project.LatestDeployment.Configuration.BuildCommand)
		fmt.Println("Configuration Framework: ", project.LatestDeployment.Configuration.Framework)
		fmt.Println("Configuration Workspace: ", project.LatestDeployment.Configuration.Workspace)
		fmt.Println("Configuration Publish Dir: ", project.LatestDeployment.Configuration.PublishDir)
		fmt.Println("------Logs------")
		for _, log := range project.LatestDeployment.Logs {
			fmt.Println(log.Log)
		}
		fmt.Println("-----/Logs------")
		


	},
}

func init(){
	rootCmd.AddCommand(projectCmd)
	
	projectCmd.Flags().StringVarP(&ProjectId, "project", "p", "", "Project ID")
}