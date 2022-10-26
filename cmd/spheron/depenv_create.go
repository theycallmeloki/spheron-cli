package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var depEnvCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a deployment environment",
	Long:  `Create a deployment environment`,
	Run: func(cmd *cobra.Command, args []string) {

		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		name := SanitizeInput("Enter the name for the deployment environment: ")
		listOfBranches, err := SanitizeInputMultiple("Enter Branch to Deploy: ", "Do you want to add another branch?")
		
		if err != nil {
			panic(err)
		}

		protocol, err := SanitizeFixedSelect(protocolList, "Select the protocol for the deployment environment: ")

		if err != nil {
			panic(err)
		}

		constructedDepEnv := spheron.DeploymentEnvironment{
			Name: name,
			Branches: listOfBranches,
			Protocol: GetProtocolEnum(protocol),
		}

		postedDepEnv, err := spheron.PostDeploymentEnvironmentVariable(projectId, constructedDepEnv)
		if err != nil {
			panic(err)
		}

		fmt.Println("Created Deployment Environment: ")
		fmt.Println("----------Deployment Environment--------")
		fmt.Println("ID: ", postedDepEnv.ID)
		fmt.Println("Name: ", postedDepEnv.Name)
		fmt.Println("Branches: ")
		for _,branch := range postedDepEnv.Branches {
			fmt.Println("- - > " , branch)
		}
		fmt.Println("Status: ", postedDepEnv.Status)
		fmt.Println("Protocol: ", postedDepEnv.Protocol)
		fmt.Println("---------/Deployment Environment--------")

		
	},
}

func init() {
	depEnvCmd.AddCommand(depEnvCreateCmd)
}