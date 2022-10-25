package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

// SetDeployment is a command to set the deployment

var setDeploymentCmd = &cobra.Command{
	Use:  "deployment",
	Short: "Sets the deployment",
	Long: `Allows you to fix a deployment ID for use throughout the session`,
	Run: func(cmd *cobra.Command, args []string) {
		projectID := viper.GetString("project")

		if(projectID == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		deployments, err := spheron.GetProjectDeployments(projectID)

		if err != nil {
			panic(err)
		}

		deploymentNames := []string{}

		for _, deployment := range deployments {
			deploymentNames = append(deploymentNames, deployment.Branch + " - " + deployment.CommitID + " - (" + deployment.CommitMessage + ")")
		}

		selectedDeploymentName, err := SanitizeFixedSelect(deploymentNames, "Select a deployment to set: ")

		if err != nil {
			panic(err)
		}

		for _, deployment := range deployments {
			if deployment.Branch + " - " + deployment.CommitID + " - (" + deployment.CommitMessage + ")" == selectedDeploymentName {
				viper.Set("deployment", deployment.ID)
				WriteLocalConfig()
			}
		}

		fmt.Println("Set " + selectedDeploymentName + " as the default deployment")

	},
}

func init() {
	if(version == "TEST_BUILD"){
		setCmd.AddCommand(setDeploymentCmd)
	}
}
