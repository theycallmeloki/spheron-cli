package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Manage domains",
	Long:  `Manage domains`,
	Run: func(cmd *cobra.Command, args []string) {
		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		domains, err := spheron.GetDomains(projectId)
		if err != nil {
			panic(err)
		}

		for _, domain := range domains {
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
		}
		
	},
}

func init() {
	rootCmd.AddCommand(domainCmd)
}