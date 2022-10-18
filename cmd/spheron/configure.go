package spheron

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var Secret string

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure secret key(s) for use with spheronctl CLI",
	Long:  `
Cache your login credentials for the Spheron Platform in your host machine. 
This will allow you to use the CLI without having to specify your secret key every time.

You can create a secret key from the Spheron Console. 
Profile > User Settings > Tokens > Create Token

Example usage: 

spheornctl configure
spheronctl configure --secret=<YOUR_SECRET_API_KEY>
	`,
	Run: func(cmd *cobra.Command, args []string) {

		// c := spheron.ReadLocalConfig()

		if Secret == "" {
			fmt.Println("Enter your secret API key to get started.")
			fmt.Println("You can create an API key from the Spheron Console. Profile > User Settings > Tokens > Create Token")
			Secret = spheron.SanitizeInput("Secret API Key: ")
		}



		viper.Set("secret", Secret)
		// viper.Set("organization", c.organization)

		spheron.WriteLocalConfig()

		initScope, err := spheron.GetScope()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("\n")
		fmt.Println("Using the first organization as default")
		organization, err := spheron.GetOrganization(initScope.Organizations[0].ID)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Organization ID: ", initScope.Organizations[0].ID)
		fmt.Println("Organization Name: ", organization.Profile.Name)
		viper.Set("organization", initScope.Organizations[0].ID)

		spheron.WriteLocalConfig()

		fmt.Println("\n")
		fmt.Println("Configuration complete! You can now use the CLI without specifying your secret key / organization.")
		fmt.Println("Hint: You can use spheronctl organization switch to change your default organization.")
		fmt.Println("\n")

	},
}


func init() {

    spheron.SetSpheronConfigFile()

    // Add a new flag to the root command
    rootCmd.AddCommand(configureCmd)

    // Accepting flag for Secret API Key
    configureCmd.Flags().StringVarP(&Secret, "secret", "s", "", "Secret API Key for Spheron Platform")
}