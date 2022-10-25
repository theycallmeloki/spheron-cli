package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var editEnvironmentId string

// EnvEdit is a command to edit environment variables
var envEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit environment variables",
	Long:  `Edit environment variables`,
	Run: func(cmd *cobra.Command, args []string) {

		projectId := viper.GetString("project")

		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		if(editEnvironmentId == "") {
			editEnvironmentId = SanitizeInput("Enter environment ID to edit: ")
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

		updatedEnv, err := spheron.PutEnvironmentVariable(projectId, editEnvironmentId, constructedEnvVariable)
		if err != nil {
			panic(err)
		}

		fmt.Println(updatedEnv.Name, updatedEnv.Value, updatedEnv.ID)
		
	},
}

func init() {
	envCommand.AddCommand(envEditCmd)

	envEditCmd.PersistentFlags().StringVarP(&editEnvironmentId, "environment", "e", "", "Environment ID")
}