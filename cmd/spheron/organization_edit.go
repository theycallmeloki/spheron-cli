package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var Name string
var Username string
var Image string

var organizationEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an organization",
	Long:  `Edit an organization`,
	Run: func(cmd *cobra.Command, args []string) {
		organizationID := viper.GetString("organization")
		organization, err := spheron.GetOrganization(organizationID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Current Organization Details:")
		fmt.Println("\n")
		fmt.Println("Organization ID: ", organizationID)
		fmt.Println("Organization Name: ", organization.Profile.Name)
		fmt.Println("Organization Username: ", organization.Profile.Username)

		if(Name == "" || Username == "" || Image == "") {
			fmt.Println("\n")
			fmt.Println("Enter new organization details:")
			fmt.Println("\n")
		}

		if Name == "" {
			Name = spheron.SanitizeInput("Organization Name: ")
		}
		if Username == "" {
			Username = spheron.SanitizeInput("Organization Username: ")
		}
		if Image == "" {
			Image = spheron.SanitizeInput("Organization Image: ")
		}

		fmt.Println("\n")

		changed, err := spheron.PutOrganization(organizationID, Name, Username, Image)

		if err != nil {
			fmt.Println(err)
		}
		if(changed) {
			fmt.Println("Organization Details Updated!")
			fmt.Println("\n")
			fmt.Println("Organization ID: ", organizationID)
			fmt.Println("Organization Name: ", Name)
			fmt.Println("Organization Username: ", Username)
			fmt.Println("Organization Image: ", Image)
		} else {
			fmt.Println("Organization Not Updated")
		}
	},
}

func init() {
	organizationCmd.AddCommand(organizationEditCmd)

	organizationEditCmd.Flags().StringVarP(&Name, "name", "n", "", "Organization Name")
	organizationEditCmd.Flags().StringVarP(&Username, "username", "u", "", "Organization Username")
	organizationEditCmd.Flags().StringVarP(&Image, "image", "i", "", "Organization Image")
}