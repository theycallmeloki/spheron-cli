package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var organizationCmd = &cobra.Command{
	Use:   "organization",
	Short: "Manage organizations",
	Long:  `Manage organizations`,
	Run: func(cmd *cobra.Command, args []string) {
		organizationID := viper.GetString("organization")
		// get all the projects in the organization and print them
		projects, err := spheron.GetOrganizationProjects(organizationID)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		
		fmt.Printf("You have %d project(s) in your organization\n", len(projects))
		

		// print all the project details
		for _, project := range projects {
			fmt.Println("--------PROJECT--------")
			fmt.Println("Project ID: ", project.ID)
			fmt.Println("Project Name: ", project.Name)
			fmt.Println("Project URL: ", project.URL)
			fmt.Println("Provider: ", project.Provider)
			fmt.Println("-------/PROJECT--------")
		}

		// get the overdue status of the organization, and print it
		overdue, err := spheron.GetOrganizationOverdue(organizationID)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if(overdue.Overdue) {
			fmt.Println(overdue.Message)
			// TODO: print the table for items that are available / exceeded capacity
		} else {
			fmt.Println(overdue.Message)
		}

		// get available coupons of the organization, and print them
		coupons, err := spheron.GetOrganizationCoupons(organizationID)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if(len(coupons) > 0) {
			fmt.Println("You have the following coupons available:")
			for _, coupon := range coupons {
				fmt.Println("--------COUPON--------")
				fmt.Println("Coupon ID: ", coupon.ID)
				fmt.Println("Coupon Name: ", coupon.CouponName)
				fmt.Println("Coupon Activated At: ", coupon.ActivatedAt)
				fmt.Println("Coupon Expires At: ", coupon.ExpiresAt)
				fmt.Println("Coupon Total Days: ", coupon.TotalDays)
				fmt.Println("Coupon Remaining Days: ", coupon.DaysRemaning)
				fmt.Println("Coupon Days Until Activation: ", coupon.DaysUntilActivation)
				fmt.Println("-------/COUPON--------")
			}
		} else {
			fmt.Println("You have no coupons available")
		}

		// get the pending invites of the organization, and print them
		invites, err := spheron.GetOrganizationInvites(organizationID)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if(len(invites) > 0) {
			fmt.Println("You have the following invites pending:")
			for _, invite := range invites {
				fmt.Println("--------INVITE--------")
				fmt.Println("User Email: ", invite.UserEmail)
				fmt.Println("-------/INVITE--------")
			}
		} else {
			fmt.Println("You have no invites pending")
		}


		if(version == "TEST_BUILD") {
			fmt.Println("\n-----------DEBUG-----------\n")
			fmt.Printf("%+v\n", projects)
			fmt.Println("\n-----------DEBUG-----------\n")
		}
		fmt.Println("\n")
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(organizationCmd)
}