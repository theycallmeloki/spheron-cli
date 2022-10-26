package spheron

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var verifyDomainId string

var verifyDomainCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a domain",
	Long:  `Verify a domain`,
	Run: func(cmd *cobra.Command, args []string) {
		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		if(verifyDomainId == "") {
			verifyDomainId = SanitizeInput("Enter domain ID to verify: ")
		}

		verifiedDomain, err := spheron.VerifyDomain(projectId, verifyDomainId)

		if err != nil {
			panic(err)
		}

		if(verifiedDomain) {
			println("Domain verified successfully")
		} else {
			println("Domain not verified")
		}
	},
}

func init() {
	domainCmd.AddCommand(verifyDomainCmd)
	verifyDomainCmd.Flags().StringVarP(&verifyDomainId, "domain", "d", "", "Domain ID to verify")
}