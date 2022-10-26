package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var domainCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a domain",
	Long:  `Create a domain`,
	Run: func(cmd *cobra.Command, args []string) {
		
		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		domainName := SanitizeInput("Enter domain name: ")

		domainLink := SanitizeInput("Enter domain link: ")
		
		domainType, err := SanitizeFixedSelect(domainTypeList, "Enter domain type: ")

		if err != nil {
			panic(err)
		}

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
		

		domain := spheron.CreateDomainPayload{
			Link: domainLink,
			Name: domainName,
			Type: GetDomainTypeEnum(domainType),
			DeploymentEnvironments: selectedEnvName,
			IsLatest: true,
		}

		createdDomain, err := spheron.PostDomain(projectId, domain)

		if err != nil {
			panic(err)
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
		
	},
}

func init() {
	domainCmd.AddCommand(domainCreateCmd)
}