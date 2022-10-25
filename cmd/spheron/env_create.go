package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

// EnvCreate is a command to create an environment variable
var EnvCreate = &cobra.Command{
	Use:   "create",
	Short: "Creates an environment variable",
	Long:  `Creates an environment variable`,
	Run: func(cmd *cobra.Command, args []string) {

		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		deploymentEnvironmentVariables, err := spheron.GetDeploymentEnvironmentVariables(projectId)
		if err != nil {
			panic(err)
		}

		var showingDepEnvs []string

		for _, env := range deploymentEnvironmentVariables {
			showingDepEnvs = append(showingDepEnvs, env.Name)
		}


		envKey := SanitizeInput("Enter the environment variable key: ")
		envValue := SanitizeInput("Enter the environment variable value: ")

		selectedEnvName, err := SanitizeSelectMultiple(showingDepEnvs, "Select Deployment Environment: ", "Do you want to add another Deployment Environment?")

		if err != nil {
			panic(err)
		}

		constructedEnvVariable := spheron.EnvironmentVariables{
			Name: envKey,
			Value: envValue,
			DeploymentEnvironments: selectedEnvName,
		}
		envsToPush := []spheron.EnvironmentVariables{constructedEnvVariable}

		pushedEnvs, err := spheron.PostEnvironmentVariables(projectId, envsToPush)

		fmt.Println("Pushing the following environment variables to Spheron: ")

		for  _, env := range pushedEnvs {
			fmt.Println(env.Name, env.Value, env.ID)
		}

	},
}

func init() {
	envCommand.AddCommand(EnvCreate)
}