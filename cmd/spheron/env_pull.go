package spheron

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var DeploymentEnvironmentName string
// var DeploymentEnvironmentId string
var envFile string

// env pull pulls from spheron and writes to specified `envfile`
var envPullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls environment variables from Spheron",
	Long:  `Pulls environment variables from Spheron`,
	Run: func(cmd *cobra.Command, args []string) {

		projectId := viper.GetString("project")

		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		if(DeploymentEnvironmentName == "") {
			deploymentEnvironmentVariables, err := spheron.GetDeploymentEnvironmentVariables(projectId)
			if err != nil {
				panic(err)
			}
			fmt.Println(deploymentEnvironmentVariables)

			var showingDepEnvs []string

			for _, env := range deploymentEnvironmentVariables {
				showingDepEnvs = append(showingDepEnvs, env.Name)
			}

			selectedDeploymentEnvironmentName, err := SanitizeFixedSelect(showingDepEnvs, "Select a deployment environment to pull: ")

			if err != nil {
				panic(err)
			}

			DeploymentEnvironmentName = selectedDeploymentEnvironmentName

			for _, env := range deploymentEnvironmentVariables {
				if env.Name == selectedDeploymentEnvironmentName {
					DeploymentEnvironmentId = env.ID
				}
			}
		}

		projectDeployments, err := spheron.GetProjectDeployments(projectId)

		if err != nil {
			panic(err)
		}


		var envsToDump []string

		for _, deployment := range projectDeployments {
			for _, env := range deployment.Project.EnvironmentVariables {
				if(contains(env.DeploymentEnvironments,DeploymentEnvironmentId)){
					//fmt.Println("Env Key", env.Name)
					//fmt.Println("Env Value", env.Value)
					envsToDump = append(envsToDump, fmt.Sprintf("%s=%s", env.Name, env.Value))
				}
				
			}
		}

		fmt.Println("Environment Variables: ")
		for _, env := range envsToDump {
			fmt.Println(env)
		}
		fmt.Println("")

		if(envFile == "") {
			envFile = SanitizeInput("Enter the file name to write the environment variables to: ")
		}

		// open the file for writing
		file, err := os.Create(envFile)
		if err != nil {
			panic(err)
		}

		// write the data to the file
		for _, env := range envsToDump {
			_, err = file.WriteString(env + "\n")
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("")

		fmt.Println("Environment variables written to " + envFile)


	},
}

func init() {
	envCommand.AddCommand(envPullCmd)

	envPullCmd.Flags().StringVarP(&envFile, "envfile", "e", "", "The file to write the environment variables to")
}