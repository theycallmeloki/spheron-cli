package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var DeploymentEnvironmentName string
// var DeploymentEnvironmentId string

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


		for _, deployment := range projectDeployments {
			for _, env := range deployment.Project.EnvironmentVariables {
				if(contains(env.DeploymentEnvironments,DeploymentEnvironmentId)){
					fmt.Println(env)
				}
				
			}
		}

	},
}

func init() {
	envCommand.AddCommand(envPullCmd)
}