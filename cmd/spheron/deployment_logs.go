package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var logsDeploymentCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get logs for latest deployment",
	Long:  `Get logs for latest deployment`,
	Run: func(cmd *cobra.Command, args []string) {

		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		project, err := spheron.GetProject(projectId)

		if err != nil {
			panic(err)
		}

		for _, log := range project.LatestDeployment.Logs {
			fmt.Println(log.Time, log.Log)
		}

		
	},
}

func init() {
	deploymentCmd.AddCommand(logsDeploymentCmd)
}