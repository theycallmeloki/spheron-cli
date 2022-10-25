package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

// set project - lets you set the project ID
var setProjectCmd = &cobra.Command{
	Use:  "project",
	Short: "Sets the project",
	Long: `Allows you to fix a project ID for use throughout the session`,
	Run: func(cmd *cobra.Command, args []string) {
		organizationID := viper.GetString("organization")	

		projects, err := spheron.GetOrganizationProjects(organizationID)

		if err != nil {
			panic(err)
		}

		projectNames := []string{}

		for _, project := range projects {
			projectNames = append(projectNames, project.Name)
		}

		selectedProjectName, err := SanitizeFixedSelect(projectNames, "Select a project to set: ")

		if err != nil {
			panic(err)
		}

		for _, project := range projects {
			if project.Name == selectedProjectName {
				viper.Set("project", project.ID)
				WriteLocalConfig()
			}
		}

		fmt.Println("Set " + selectedProjectName + " as the default project")
	},
}

func init() {
	setCmd.AddCommand(setProjectCmd)
}

