package spheron

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var EnvFile string

// env push pushes to spheron from specified `envfile`
var envPushCmd = &cobra.Command{
	Use:   "push",
	Short: "Pushes environment variables to Spheron",
	Long:  `Pushes environment variables to Spheron`,
	Run: func(cmd *cobra.Command, args []string) {

		projectId := viper.GetString("project")

		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		if(EnvFile == "") {
			EnvFile = SanitizeInput("Enter the path to the env file: ")
		}

		if(!FileExists(EnvFile)) {
			panic("File does not exist")
		}

		deploymentEnvironmentVariables, err := spheron.GetDeploymentEnvironmentVariables(projectId)
		if err != nil {
			panic(err)
		}

		var showingDepEnvs []string

		for _, env := range deploymentEnvironmentVariables {
			showingDepEnvs = append(showingDepEnvs, env.Name)
		}

		selectedEnvName, err := SanitizeSelectMultiple(showingDepEnvs, "Select Deployment Environment: ", "Do you want to add another Deployment Environment?")

		envsToPush := []spheron.EnvironmentVariables{}

		content, err := ioutil.ReadFile(EnvFile)
		if err != nil {
			panic(err)
		}

		keyValuePairs := strings.Split(string(content), "\n")

		for _, pair := range keyValuePairs {
			keyValue := strings.Split(pair, "=")
			if(len(keyValue) == 2) {
				// fmt.Println(keyValue[0], keyValue[1])
				envsToPush = append(envsToPush, spheron.EnvironmentVariables{
					Name:  keyValue[0],
					Value: keyValue[1],
					DeploymentEnvironments: selectedEnvName,
				})
			}
		}

		pushedEnvs, err := spheron.PostEnvironmentVariables(projectId, envsToPush)
		if err != nil {
			panic(err)
		}

		fmt.Println("Pushing the following environment variables to Spheron:")

		for  _, env := range pushedEnvs {
			fmt.Println(env.Name, env.Value, env.ID)
		}


	},

}

func init() {
	envCommand.AddCommand(envPushCmd)

	envPushCmd.Flags().StringVarP(&EnvFile, "envfile", "e", ".env", "The file to read environment variables from")
}