package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var UserEmail string

var organizationInviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite a user to an organization",
	Long:  `Invite a user to an organization`,
	Run: func(cmd *cobra.Command, args []string) {
		organizationID := viper.GetString("organization")
		if UserEmail == "" {
			UserEmail = SanitizeInput("User Email: ")
		}
		invite, err := spheron.InviteOrganizationMember(organizationID, UserEmail)
		if err != nil {
			panic(err)
		}
		fmt.Println("Invite Sent!")
		fmt.Println("\n")
		fmt.Println("Invite Email: ", invite.UserEmail)
		fmt.Println("Invite Status: ", invite.Status)

		if(version == "TEST_BUILD") {
			fmt.Println("\n-----------DEBUG-----------\n")
			fmt.Printf("%+v\n", invite)
			fmt.Println("\n-----------DEBUG-----------\n")
		}
	},
}

func init(){
	organizationCmd.AddCommand(organizationInviteCmd)

	organizationInviteCmd.Flags().StringVarP(&UserEmail, "email", "e", "", "User Email")
}