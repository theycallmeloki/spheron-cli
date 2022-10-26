package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

// lists deployment environments
var depEnvCmd = &cobra.Command{
	Use:   "depenv",
	Short: "Manage deployment environments",
	Long:  `Manage deployment environments`,
	Run: func(cmd *cobra.Command, args []string) {

		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		deploymentEnvironmentVariables, err := spheron.GetDeploymentEnvironmentVariables(projectId)
		if err != nil {
			panic(err)
		}

		for _, env := range deploymentEnvironmentVariables {
			fmt.Println("----------Deployment Environment--------")
			fmt.Println("Deployment Environment ID: ", env.ID)
			fmt.Println("Name: ", env.Name)
			fmt.Println("Branches: ")
			for _,branch := range env.Branches {
				fmt.Println("- - > " , branch)
			}
			fmt.Println("Status: ", env.Status)
			fmt.Println("Protocol: ", env.Protocol)
			fmt.Println("---------/Deployment Environment--------")
		}
		
	},
}

func init() {
	rootCmd.AddCommand(depEnvCmd)
}