package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var editDomainId string

var editDomainCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a domain",
	Long:  `Edit a domain`,
	Run: func(cmd *cobra.Command, args []string) {
		
		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		if(editDomainId == "") {
			editDomainId = SanitizeInput("Enter domain ID to edit: ")
		}

		domainName := SanitizeInput("Enter domain name: ")

		domainLink := SanitizeInput("Enter domain link: ")

		deploymentEnvironmentVariables, err := spheron.GetDeploymentEnvironmentVariables(projectId)
		if err != nil {
			panic(err)
		}

		var showingDepEnvs []string

		for _, env := range deploymentEnvironmentVariables {
			showingDepEnvs = append(showingDepEnvs, env.Name)
		}

		fmt.Println(showingDepEnvs)

		selectedEnvName, err := SanitizeSelectMultiple(showingDepEnvs, "Select Deployment Environment: ", "Do you want to add another Deployment Environment?")

		domainToUpdate := spheron.UpdateDomainPayload{
			Link: domainLink,
			Name: domainName,
			DeploymentEnvironments: selectedEnvName,
			IsLatest: true,
		}


		updatedDomain, err := spheron.PatchDomain(projectId, editDomainId, domainToUpdate)

		if err != nil {
			panic(err)
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
		
	},
}

func init() {
	domainCmd.AddCommand(editDomainCmd)
	editDomainCmd.Flags().StringVarP(&editDomainId, "domain", "d", "", "Domain ID to edit")
}