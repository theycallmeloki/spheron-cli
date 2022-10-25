package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

// OrganizationSwitch is a command to switch to a different organization

var organizationSwitchCmd = &cobra.Command{
	Use:  "switch",
	Short: "Switch to a different organization",
	Long: `Switch to a different organization`,
	Run: func(cmd *cobra.Command, args []string) {

		scope, err := spheron.GetScope()
		if err != nil {
			panic(err)
		}

		organizationNames := []string{}

		for _, organization := range scope.Organizations {
			organizationNames = append(organizationNames, organization.Name)
		}

		selectedOrganizationName, err := SanitizeFixedSelect(organizationNames, "Select an organization to switch to: ")

		if err != nil {
			panic(err)
		}

		for _, organization := range scope.Organizations {
			if organization.Name == selectedOrganizationName {
				viper.Set("organization", organization.ID)
				WriteLocalConfig()
			}
		}

		fmt.Println("Set " + selectedOrganizationName + " as the default organization")

	},

}

func init() {
	organizationCmd.AddCommand(organizationSwitchCmd)
}
