package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var viewDomainId string

var viewDomainCmd = &cobra.Command{
	Use:   "view",
	Short: "View a domain",
	Long:  `View a domain`,
	Run: func(cmd *cobra.Command, args []string) {
		
		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		if(viewDomainId == "") {
			viewDomainId = SanitizeInput("Enter domain ID to view: ")
		}

		domain, err := spheron.GetDomain(projectId, viewDomainId)

		if err != nil {
			panic(err)
		}

		fmt.Println("----------Domain--------")
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
		fmt.Println("---------/Domain--------")
		
	},
}

func init() {
	domainCmd.AddCommand(viewDomainCmd)

	viewDomainCmd.Flags().StringVarP(&viewDomainId, "domain", "d", "", "Domain ID")
}