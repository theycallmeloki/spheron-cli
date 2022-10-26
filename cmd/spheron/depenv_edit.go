package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var editDepEnvId string

var editDepEnvCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a deployment environment",
	Long:  `Edit a deployment environment`,
	Run: func(cmd *cobra.Command, args []string) {

		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		if(editDepEnvId == "") {
			editDepEnvId = SanitizeInput("Enter deployment environment ID to edit: ")
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

		updatedDepEnv, err := spheron.PutDeploymentEnvironmentVariable(projectId, editDepEnvId, constructedDepEnv)

		if err != nil {
			panic(err)
		}

		fmt.Println("Edited Deployment Environment: ")
		fmt.Println("----------Deployment Environment--------")
		fmt.Println("ID: ", updatedDepEnv.ID)
		fmt.Println("Name: ", updatedDepEnv.Name)
		fmt.Println("Branches: ")
		for _,branch := range updatedDepEnv.Branches {
			fmt.Println("- - > " , branch)
		}
		fmt.Println("Status: ", updatedDepEnv.Status)
		fmt.Println("Protocol: ", updatedDepEnv.Protocol)
		fmt.Println("---------/Deployment Environment--------")

	},
}

func init() {
	depEnvCmd.AddCommand(editDepEnvCmd)

	editDepEnvCmd.Flags().StringVarP(&editDepEnvId, "deployment", "d", "", "Deployment Environment ID")
}